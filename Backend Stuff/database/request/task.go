package request

import (
    "time"
    "fmt"
    "net/url"
    "errors"
    "strings"
    "strconv"

    "google.golang.org/api/iterator"
    "cloud.google.com/go/firestore"
)

// Repeat setting constants
const (
    NEVER    = "never"
    DAILY    = "daily"
    WEEKLY   = "weekly"
    BIWEEKLY = "biweekly"
    MONTHLY  = "monthly"
    ANNUALLY = "annually"
)

// Reminder setting constants
const (
    NONE = "none"
    ATOE = "at time of event"
    MBE  = "minutes before event"
    DBE  = "days before event"
    WBE  = "weeks before event"
)

// Structure of the documents in the tasks collection
// encoding (`firestore:"..."`) is firestore so we can
// easily dump requested data into this structure for
// easy access later.
type Task struct {
    // Firestore generated task ID
    Id          string      `firestore:"id,omitempty"`

    // Name of the task
    Name        string      `firestore:"task_name,omitempty"`

    // User ID of the user who owns this task
    Owner       string      `firestore:"task_owner,omitempty"`

    // ID of the parent list
    Parent      string      `firestore:"parent_id,omitempty"`

    // Whether or not someone can edit this task
    Lock        bool        `firestore:"lock,omitempty"`

    // Date this task is due
    DueDate     time.Time   `firestore:"due_date,omitempty"`

    // When the user would idelly like to start this task
    IdealStart  time.Time   `firestore:"ideal_start_date"`

    // Date task was started
    StartDate   time.Time   `firestore:"start_date,omitempty"`

    // Whether or not we should repeat this task
    Repeating   bool        `firestore:"repeating,omitempty"`

    // The frequency of the repat, if we are repeating
    Repeat      string      `firestore:"repeat,omitempty"`

    // Whether or not we should remind the user
    Remind      bool        `firestore:"remind,omitempty"`

    // Time frame before task to remind the user -- string
    Reminder    string      `firestore:"reminder,omitempty"`

    // Time frame before task to remind the user -- int
    TimeFrame   int         `firestore:"time_frame,omitempty"`

    // Location of the task
    Location    string      `firestore:"location,omitempty"`

    // Description of the task
    Description string      `firestore:"description,omitempty"`

    // Url associated with the task -- could be an array if desired
    Url         string      `firestore:"url,omitempty"`

    // Whether or not this list is shared
    Shared      bool     `firestore:"shared,omitempty"`

    // Array of user IDs of the users this list has been shared with
    SharedUsers []string `firestore:"shared_users,omitempty"`

    // IDs of assoociated Subtasks
    Subtasks    []string    `firestore:"sub_tasks,omitempty"`

    // Whether or not this is a subtask
    Subtask     bool        `firestore:"sub_task,omitempty"`
}

// Structure of the documents in the tasks collection
// encoding (`json:"..."`) is json so we can pass the
// structure to the server in the correct json format.
// Task data will be transferred over from Task struct.
type TaskJSON struct {
    // Firestore generated task ID
    Id          string      `json:"id,omitempty"`

    // Name of the task
    Name        string      `json:"task_name,omitempty"`

    // User ID of the user who owns this task
    Owner       string      `json:"task_owner,omitempty"`

    // ID of the parent list
    Parent      string      `json:"parent_id,omitempty"`

    // In case we wanted to allow tasks not to be list bound
    //InList      bool        `json:in_list`

    // Whether or not someone can edit this task
    Lock        bool        `json:"lock,omitempty"`

    // Date this task is due
    DueDate     time.Time   `json:"due_date,omitempty"`

    // When the user would idelly like to start this task
    IdealStart  time.Time   `json:"ideal_start_date"`

    // Date task was started
    StartDate   time.Time   `json:"start_date,omitempty"`

    // Whether or not we should repeat this task
    Repeating   bool        `json:"repeating,omitempty"`

    // The frequency of the repat, if we are repeating
    Repeat      string      `json:"repeat,omitempty"`

    // Whether or not we should remind the user
    Remind      bool        `json:"remind,omitempty"`

    // Time frame before task to remind the user -- string
    Reminder    string      `json:"reminder,omitempty"`

    // Time frame before task to remind the user -- int
    TimeFrame   int         `json:"time_frame,omitempty"`

    // Location of the task
    Location    string      `json:"location,omitempty"`

    // Description of the task
    Description string      `json:"description,omitempty"`

    // Url associated with the task -- could be an array if desired
    Url         string      `json:"url,omitempty"`

    // Whether or not this list is shared
    Shared      bool     `json:"shared,omitempty"`

    // Array of user IDs of the users this list has been shared with
    SharedUsers []string `json:"shared_users,omitempty"`

    // IDs of assoociated Subtasks
    Subtasks    []string    `json:"sub_tasks,omitempty"`

    // Whether or not this is a subtask
    Subtask     bool        `json:"sub_task,omitempty"`
}

// AddTask {{{
//
// Adds a task to the task collection, setting any fields provided
// Returns the newly added task in JSON format and nil if no errors
// occurs, returns the error and null TaskJSON if an error occurss
func (r *Request) AddTask(name string, fields url.Values) (*TaskJSON, error) {
    var tjson *TaskJSON
    var task Task

    // Create new task document in Firestore
    ref := r.Client.Collection("tasks").NewDoc()
    task.Id = ref.ID

    data := r.ParseTaskFields(fields)

    if data["task_name"] != name {
        data["task_name"] = name
    }

    data["task_owner"] = r.UserId

    // I'm going to not make sub tasks automatically.. can re-add later.
    /*if data["sub_tasks"] == nil {
        var tasks []string
        tasks = append(tasks, "")
        data["sub_tasks"] = tasks
    } */

    //fmt.Printf("%v\n", data)

    // Send the parsed task values to Firestore
    _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err setting new task data: %v", err)
        return tjson, errors.New(e)
    }
    if name != "first_task" && data["parent_id"] != nil{
        if data["sub_task"] != nil && data["sub_task"].(bool) {
            r.UpdateTaskSubtask(data["parent_id"].(string), ref.ID)
        } else {
            r.UpdateListTask(data["parent_id"].(string),ref.ID)
        }
    }

    tjson, err = r.GetTaskByName(name)
    return tjson, err
} // }}}

// func GetTaskByName {{{
//
// Returns a task using the tasks name
// Ensures we get the correct task by specifying the task owner
func (r *Request) GetTaskByName(name string) (*TaskJSON, error) {
    var tjson *TaskJSON
    var task Task

    // Get all tasks from Firestore where the task_name is the same as the one provided
    iter := r.Client.Collection("tasks").Where("task_name", "==", name).Where("task_owner", "==", r.UserId).Documents(r.Ctx)

    // For each document
    for {
        // Get a snapshot of the data
        docsnap, err := iter.Next()

        // Check if we're done with our loop
        if err == iterator.Done {
            break
        }

        // Check if we have some other error
        if err != nil {
            e := fmt.Sprintf("err getting snapshot of task data: %v", err)
            return tjson, errors.New(e)
        }

        // Put data into our task structure
        docsnap.DataTo(&task)
        // Get & set the task ID
        id := docsnap.Ref.ID
        task.Id = id
        r.Task = &task
    }

    // Set our request task to be this task
    if r.Task != nil {
        tjson = r.TaskToJSON()
    }

    return tjson, nil
} // }}}

// func GetTaskByID {{{
//
// Return a task with a given id; check to see if requesting user
// has permission to view that task before returning it.
func (r *Request) GetTaskByID(tid string) (*TaskJSON, error) {
    var tjson *TaskJSON
    var task Task

    // Get the Firestore path for the task
    taskidpath := fmt.Sprintf("tasks/%s", tid)

    // Pass that to Firestore
    doc := r.Client.Doc(taskidpath)

    // Get a snapshot of the task data
    docsnap, err := doc.Get(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("err getting snapshot of task data: %v", err)
        return tjson, errors.New(e)
    }

    // Add the data to our structure
    err = docsnap.DataTo(&task)
    if err != nil {
        e := fmt.Sprintf("err putting task data to structure: %v", err)
        return tjson, errors.New(e)
    }

    // Get & set the user ID
    id := docsnap.Ref.ID
    task.Id = id

    if task.Owner != r.UserId {
        if task.SharedUsers == nil || !r.CheckIfShared(task.SharedUsers){
            return tjson, errors.New("err getting task: requestor does not have permission")
        }
    }

    // Set our request task to be this task
    r.Task = &task
    if r.Task != nil {
        tjson = r.TaskToJSON()
    }
    return tjson, nil
} // }}}

// func GetTasks {{{
//
// Returns all tasks in a given list that the user has permission to view
func (r *Request) GetTasks(parentid string) ([]*TaskJSON, error) {
   var tasks []*TaskJSON

   // Get all tasks from Firestore where the owner is the requesting user and the parent is the same as the one provided
   iter := r.Client.Collection("tasks").Where("parent", "==", parentid).Documents(r.Ctx)

   // For each document
   for {
       // Get a snapshot of the data
       docsnap, err := iter.Next()

       // Check if we're done with our loop
       if err == iterator.Done {
           break
       }

       // Check if we have some other error
       if err != nil {
           e := fmt.Sprintf("err getting snapshot of task data: %v", err)
           return tasks, errors.New(e)
       }

       // create a new task struct
       var task Task

       // Put doc data into our task structure
       docsnap.DataTo(&task)
       if task.Owner != r.UserId {
           if task.SharedUsers == nil || !r.CheckIfShared(task.SharedUsers){
               continue
           }
       }

       // Get & set the task ID
       id := docsnap.Ref.ID
       task.Id = id
       r.Task = &task

       // Add task to the tasks array
       if r.Task != nil {
           tasks = append(tasks, r.TaskToJSON())
       }
    }

    return tasks, nil
} // }}}

// func UpdateTask {{{
//
func (r *Request) UpdateTask(name string, fields url.Values) (*TaskJSON, error) {
    // Get the task using it's name
    tjson, err := r.GetTaskByName(name)
    if err != nil {
        e := fmt.Sprintf("err getting task for update: %v", err)
        return tjson, errors.New(e)
    }

    if tjson.Owner != r.UserId {
        if tjson.SharedUsers == nil || !r.CheckIfShared(tjson.SharedUsers){
            return tjson, errors.New("err updating task: requestor does not have permission")
        }
    }

    // Parse the url fields into a map for Firestore
    data := r.ParseTaskFields(fields)

    //log.Printf("%v", data)

    // Get a reference to our task
    ref := r.Client.Collection("tasks").Doc(tjson.Id)

    // Send update to Firestore
    _,err = ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err updating task data: %v", err)
        return tjson, errors.New(e)
    }
    tjson, err = r.GetTaskByID(ref.ID)
    return tjson, err
} // }}}

// func DestroyTasks {{{
//
//
// TODO: Delete all subtasks as well
func (r *Request) DestroyTask(name string) error {
    // Get the task using it's name
    task, err := r.GetTaskByName(name)
    if err != nil {
        e := fmt.Sprintf("err getting task for delete: %v", err)
        return errors.New(e)
    }
    if task.Owner != r.UserId {
        if task.SharedUsers == nil || !r.CheckIfShared(task.SharedUsers){
            return errors.New("err deleting task: requestor does not have permission")
        }
    }

    // Get the Firestore path for the task
    taskidpath := fmt.Sprintf("tasks/%s", task.Id)

    // Check if the task has any subtasks
    if len(task.Subtasks) > 0 {
        // src: https://github.com/GoogleCloudPlatform/golang-samples/blob/810112812f3699d1cf9ad62ba3abf39f8ea99d7d/firestore/firestore_snippets/save.go#L295-L334
        // Retrieve all documents that have this task as their parent
        iter := r.Client.Collection("tasks").Where("parent", "==", task.Id).Documents(r.Ctx)
        numDeleted := 0

        batch := r.Client.Batch()
        for {
            doc, err := iter.Next()
            if err == iterator.Done {
				break
			}
			if err != nil {
                e := fmt.Sprintf("err getting snapshot of subtask for delete: %v", err)
                return errors.New(e)
			}

            // create a new task struct for the subtask
            var subtask Task

            // Put doc data into our subtask structure
            doc.DataTo(&subtask)

            if subtask.Owner != r.UserId {
                if subtask.SharedUsers == nil || !r.CheckIfShared(subtask.SharedUsers) {
                    continue
                }
            }

            if len(subtask.Subtasks) > 0 {
                r.DestroyTaskById(subtask.Id)
            }

			batch.Delete(doc.Ref)
			numDeleted++
        }

        // If there are no documents to delete,
        // the process is over.
        if numDeleted == 0 {
            return nil
    	}

    	_, err := batch.Commit(r.Ctx)
    	if err != nil {
    		return err
    	}
    }

    // Now we can delete the task
    _, err = r.Client.Doc(taskidpath).Delete(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("err deleting task: %v", err)
        return errors.New(e)
    }
    return nil
} // }}}

// func DestroyTasks {{{
//
func (r *Request) DestroyTaskById(id string) error {
    // Get the task using it's name
    task, err := r.GetTaskByID(id)
    if err != nil {
        e := fmt.Sprintf("err getting task for delete: %v", err)
        return errors.New(e)
    }
    if task.Owner != r.UserId {
        if task.SharedUsers == nil || !r.CheckIfShared(task.SharedUsers){
            return errors.New("err deleting task: requestor does not have permission")
        }
    }

    // Check if the task has any subtasks
    if len(task.Subtasks) > 0 {
        // src: https://github.com/GoogleCloudPlatform/golang-samples/blob/810112812f3699d1cf9ad62ba3abf39f8ea99d7d/firestore/firestore_snippets/save.go#L295-L334
        // Retrieve all documents that have this task as their parent
        iter := r.Client.Collection("tasks").Where("parent_id", "==", task.Id).Documents(r.Ctx)
        numDeleted := 0

        batch := r.Client.Batch()
        for {
            doc, err := iter.Next()
            if err == iterator.Done {
				break
			}
			if err != nil {
                e := fmt.Sprintf("err getting snapshot of subtask for delete: %v", err)
                return errors.New(e)
			}

            // create a new task struct for the subtask
            var subtask Task

            // Put doc data into our subtask structure
            doc.DataTo(&subtask)

            if subtask.Owner != r.UserId {
                if subtask.SharedUsers == nil || !r.CheckIfShared(task.SharedUsers) {
                    continue
                }
            }

            if len(subtask.Subtasks) > 0 {
                r.DestroyTaskById(subtask.Id)
            }
			batch.Delete(doc.Ref)
			numDeleted++
        }

        // If there are no documents to delete,
        // the process is over.
        if numDeleted == 0 {
            return nil
    	}

    	_, err := batch.Commit(r.Ctx)
    	if err != nil {
    		return err
    	}
    }
    // We aren't going to be destroying the task here.
    // It will be deleted during the batch delete, which
    // is the only place this function is called
    return nil
} // }}}

// func ParseTaskFields {{{
func (r *Request) ParseTaskFields(fields url.Values) map[string]interface{} {
    //log.Printf("%v", fields)
    var data = make(map[string]interface{})

    // Parse url fields
    for k, v := range fields {
        // Ensure the key is lower case
        k = strings.ToLower(k)

        // Our value is currently an array of strings; let's fix that
        val := strings.Join(v,"")

        // We want to check the key to ensure we don't just add a bunch of new fields
        if k == "task_name" {
          data[k] = val
        }
        if k == "lock" {
          data[k], _ = strconv.ParseBool(val)
        }
        if k == "parent_id" {
            // parent_id can be either a list_id OR a task_id
            data[k] = val
        }
        if k == "sub_task" {
            data[k], _ = strconv.ParseBool(val)
        }
    }
    return data
} // }}}

// func TaskToJSON {{{
//
func (r *Request) TaskToJSON() *TaskJSON {
    var taskjson TaskJSON

    taskjson.Id         = r.Task.Id
    taskjson.Name       = r.Task.Name
    taskjson.Owner      = r.Task.Owner
    taskjson.Parent     = r.Task.Parent
    taskjson.Lock       = r.Task.Lock
    taskjson.DueDate    = r.Task.DueDate
    taskjson.IdealStart = r.Task.IdealStart
    taskjson.StartDate  = r.Task.StartDate
    taskjson.Repeating  = r.Task.Repeating
    taskjson.Repeat     = r.Task.Repeat
    taskjson.Remind     = r.Task.Remind
    taskjson.Reminder   = r.Task.Reminder
    taskjson.TimeFrame  = r.Task.TimeFrame
    taskjson.Location   = r.Task.Location
    taskjson.Description = r.Task.Description
    taskjson.Url        = r.Task.Url
    taskjson.Shared      = r.Task.Shared
    taskjson.SharedUsers = r.Task.SharedUsers
    taskjson.Subtasks   = r.Task.Subtasks
    taskjson.Subtask    = r.Task.Subtask

    return &taskjson
} // }}}

func (r *Request) UpdateTaskSubtask(taskid, id string) {
    var task Task

    // Get the Firestore path for the user
    taskidpath := fmt.Sprintf("tasks/%s", taskid)

    // Pass that to Firestore
    doc := r.Client.Doc(taskidpath)

    // Get a snapshot of the user data
    docsnap, err := doc.Get(r.Ctx)
    if err != nil {
        fmt.Printf("err getting task snapshot: %v\n", err)
        return
    }

    // Add the data to our structure
    err = docsnap.DataTo(&task)
    if err != nil {
        fmt.Printf("err putting task data to struct: %v\n", err)
        return
    }

    task.Subtasks = append(task.Subtasks, id)
    d := make(map[string]interface{})
    d["sub_tasks"] = task.Subtasks
    // Send update to Firestore
    _, err = doc.Set(r.Ctx, d, firestore.MergeAll)
    if err != nil {
        fmt.Printf("err setting new task data: %v\n", err)
    }
}
