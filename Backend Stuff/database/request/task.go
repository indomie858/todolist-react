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
    DAILY    = "every day"
    WEEKLY   = "every week"
    BIWEEKLY = "every 2 weeks"
    MONTHLY  = "every month"
    ANNUALLY = "every year"
)

// Reminder setting constants
const (
    NONE = "none"
    ATOE = "at time of event"
    MBE  = "minutes before event"
    DBE  = "days before event"
    WBE  = "weeks before event"
)

// Priority levels
const (
    LOW  = "low"
    MED  = "medium"
    HIGH = "high"
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

    // ID of the parent list or parent task, if a subtask
    Parent      string      `firestore:"parent_id,omitempty"`

    // Name of the list the task is in
    List        string      `firestore:"list,omitempty"`

    // Whether or not someone can edit this task
    Lock        bool        `firestore:"lock,omitempty"`

    // Date this task (includes the time it is due)
    DateDue     time.Time   `firestore:"date_due,omitempty"`

    // Whether or not the task is complete / finished
    Done        bool        `firestore:"done"`

    // Whether or not we should repeat this task, used for queries
    Repeating   bool        `firestore:"repeating"`

    // The frequency of the repeat, if we are repeating
    Repeat      string      `firestore:"repeat,omitempty"`

    // The date we should stop repeating this task
    EndRepeat   time.Time   `firestore:"end_repeat,omitempty"`

    // Whether or not we should remind the user, used for queries
    Remind      bool        `firestore:"remind"`

    // What type of reminder they want, discord or email
    RemindType  string      `firestore:"remind_type,omitempty"`

    // Time frame before task to remind the user -- string
    // Similar to 'Alert' in Google Calendar
    Reminder    string      `firestore:"reminder,omitempty"`

    // Time frame before task to remind the user
    RemindTime  time.Time    `firestore:"reminder_time,omitempty"`

    // Priority level of the task
    Priority    string      `firestore:"priority,omitempty"`

    // Location of the task
    Location    string      `firestore:"location,omitempty"`

    // Description of the task (similar to notes on Apple Reminders)
    Description string      `firestore:"description,omitempty"`

    // Url associated with the task -- could be an array if desired
    Url         string      `firestore:"url,omitempty"`

    // Whether or not this list is shared
    Shared      bool        `firestore:"shared"`

    // Array of user IDs of the users this task has been shared with
    SharedUsers []string    `firestore:"shared_users,omitempty"`

    // Whether or not this is a subtask
    Subtask     bool        `firestore:"sub_task"`

    // IDs of assoociated Subtasks
    Subtasks    []string    `firestore:"sub_tasks,omitempty"`
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

    // ID of the parent list or parent task, if a subtask
    Parent      string      `json:"parent_id,omitempty"`

    // Whether or not someone can edit this task
    Lock        bool        `json:"lock,omitempty"`

    // Name of the list the task is in
    List        string      `json:"list,omitempty"`

    // Date this task is due (includes the time it is due)
    DateDue     time.Time   `json:"date_due,omitempty"`

    // Whether or not the task is complete / finished
    Done        bool        `json:"done"`

    // Whether or not we should repeat this task, used for queries
    Repeating   bool        `json:"repeating"`

    // The frequency of the repeat, if we are repeating
    Repeat      string      `json:"repeat,omitempty"`

    // The date we should stop repeating this task
    EndRepeat   time.Time   `json:"end_repeat,omitempty"`

    // Whether or not we should remind the user, used for queries
    Remind      bool        `json:"remind"`

    // What type of reminder they want, discord or email
    RemindType  string      `json:"remind_type,omitempty"`

    // Time frame before task to remind the user -- string
    Reminder    string      `json:"reminder,omitempty"`

    // Time frame before task to remind the user
    RemindTime  time.Time   `json:"reminder_time,omitempty"`

    // Priority level of the task
    Priority    string      `json:"priority,omitempty"`

    // Location of the task
    Location    string      `json:"location,omitempty"`

    // Description of the task
    Description string      `json:"description,omitempty"`

    // Url associated with the task -- could be an array if desired
    Url         string      `json:"url,omitempty"`

    // Whether or not this list is shared
    Shared      bool        `json:"shared"`

    // Array of user IDs of the users this list has been shared with
    SharedUsers []string    `json:"shared_users,omitempty"`

    // Whether or not this is a subtask
    Subtask     bool        `json:"sub_task"`

    // IDs of assoociated Subtasks
    Subtasks    []string    `json:"sub_tasks,omitempty"`
}

// AddTask {{{
//
// Adds a task to the task collection, setting any fields provided
// Returns the newly added task in JSON format and nil if no errors
// occurs, returns the error and null TaskJSON if an error occurss
func (r *Request) AddTask(name, parentid string, fields url.Values) (*TaskJSON, error) {
    //fmt.Printf("ParentId: %s\n", parentid)
    var tjson *TaskJSON

    // Create new task document in Firestore
    ref := r.Client.Collection("tasks").NewDoc()

    // Create a new map for the task data
    var data = make(map[string]interface{})

    // Let's set some default values real quick -

    data["task_owner"] = r.UserId
    data["task_name"] = name
    data["parent_id"] = parentid
    data["lock"] = false
    data["done"] = false
    data["repeating"] = false
    data["repeat"] = NEVER
    data["remind"] = false
    data["reminder"] = NONE
    data["priority"] = NONE
    data["location"] = ""
    data["description"] = ""
    data["url"] = ""
    data["shared"] = false
    data["sub_task"] = false

    // Now let's update our map to reflect the values we were given
    data = r.ParseTaskFields(fields, data)

    //fmt.Printf("%v\n", data)

    // Send the parsed task values to Firestore
    _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err setting new task data: %v", err)
        return tjson, errors.New(e)
    }

    if data["sub_task"].(bool) {
        r.UpdateTaskSubtasks(data["parent_id"].(string), ref.ID)
    } else if name != "first_task" {
        r.UpdateListTasks(data["parent_id"].(string), ref.ID)
    }

    return r.GetTaskByName(name, parentid)
} // }}}

// func GetTaskByName {{{
//
// Returns a task using the tasks name
// Ensures we get the correct task by specifying the parent list
func (r *Request) GetTaskByName(name, parentid string) (*TaskJSON, error) {
    var tjson *TaskJSON
    var task Task

    // Get all tasks from Firestore where the task_name is the same as the one provided
    iter := r.Client.Collection("tasks").Where("task_name", "==", name).Where("parent_id", "==", parentid).Documents(r.Ctx)

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

        // Make sure we're getting a task the user actually owns
        if task.Owner != r.UserId {
            // If they don't own it, check if its shared with them
            if task.SharedUsers == nil || !r.CheckIfShared(task.SharedUsers) {
                return tjson, errors.New("err getting task: requestor does not have permission")
            }
        }

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

    // Make sure we're getting a task the user actually owns
    if task.Owner != r.UserId {
        // If they don't own it, check if its shared with them
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
    iter := r.Client.Collection("tasks").Where("parent_id", "==", parentid).Documents(r.Ctx)

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

        /*/ Only get tasks the user actually owns or have been shared with them
        if task.Owner != r.UserId {
            if task.SharedUsers == nil || !r.CheckIfShared(task.SharedUsers) {
                continue
            }
            }*/

        // Get & set the task ID
        id := docsnap.Ref.ID
        task.Id = id
        r.GetTaskByName(task.Name, task.Parent)

        // Add task to the tasks array
        if r.Task != nil {
            tasks = append(tasks, r.TaskToJSON())
        }
    }

    return tasks, nil
} // }}}

// func UpdateTask {{{
//
func (r *Request) UpdateTask(name, parent string, fields url.Values) (*TaskJSON, error) {
    // Get the task using it's name & the id of the parent list / task
    tjson, err := r.GetTaskByName(name, parent)
    if err != nil {
        e := fmt.Sprintf("err getting task for update: %v", err)
        return tjson, errors.New(e)
    }

    // Parse the url fields into a map for Firestore
    var data = make(map[string]interface{})
    data = r.ParseTaskFields(fields, data)

    // Get a reference to our task
    ref := r.Client.Collection("tasks").Doc(tjson.Id)

    // Send update to Firestore
    _, err = ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err updating task data: %v", err)
        return tjson, errors.New(e)
    }

    if data["task_name"] != nil {
        tjson, err = r.GetTaskByName(data["task_name"].(string), parent)
        return tjson, err
    }

    return r.GetTaskByName(name, parent)
} // }}}

// func DestroyTasks {{{
//
//
// TODO: Delete all subtasks as well
func (r *Request) DestroyTask(name, parentid string) error {
    // Get the task using it's name
    task, err := r.GetTaskByName(name, parentid)
    if err != nil {
        e := fmt.Sprintf("err getting task for delete: %v", err)
        return errors.New(e)
    }

    // Check if we have any subtasks to delete
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

			batch.Delete(doc.Ref)
			numDeleted++
        }

        // Let's add the task we originally wanted to delete to the batch -

        // Get the Firestore path for the task
        taskidpath := fmt.Sprintf("tasks/%s", task.Id)

        taskDoc := r.Client.Doc(taskidpath)
        tdoc, err := taskDoc.Get(r.Ctx)
        if err != nil {
            e := fmt.Sprintf("DestroyTask: err getting task snapshot: %v", err)
            return errors.New(e)
        }

        batch.Delete(tdoc.Ref)
        numDeleted++

        // If there are no documents to delete,
        // the process is over.
        if numDeleted == 0 {
            return nil
    	}

    	_, err = batch.Commit(r.Ctx)
    	if err != nil {
    		return err
    	}
    } else {
        batch := r.Client.Batch()

        // Get the Firestore path for the task
        taskidpath := fmt.Sprintf("tasks/%s", task.Id)

        taskDoc := r.Client.Doc(taskidpath)
        tdoc, err := taskDoc.Get(r.Ctx)
        if err != nil {
            e := fmt.Sprintf("DestroyTaskById: err getting task snapshot: %v", err)
            return errors.New(e)
        }
        batch.Delete(tdoc.Ref)

    	_, err = batch.Commit(r.Ctx)
    	if err != nil {
    		return err
    	}
    }

    return nil
} // }}}

// func DestroyTaskById {{{
//
//
// TODO: Delete all subtasks as well
func (r *Request) DestroyTaskById(id string) error {
    // Get the task using it's name
    task, err := r.GetTaskByID(id)
    if err != nil {
        e := fmt.Sprintf("err getting task for delete: %v", err)
        return errors.New(e)
    }

    // Check if we have any subtasks to delete
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

			batch.Delete(doc.Ref)
			numDeleted++
        }

        // Let's add the task we originally wanted to delete to the batch -

        // Get the Firestore path for the task
        taskidpath := fmt.Sprintf("tasks/%s", task.Id)

        taskDoc := r.Client.Doc(taskidpath)
        tdoc, err := taskDoc.Get(r.Ctx)
        if err != nil {
            e := fmt.Sprintf("DestroyTaskById: err getting task snapshot: %v", err)
            return errors.New(e)
        }
        batch.Delete(tdoc.Ref)
        numDeleted++

        // If there are no documents to delete,
        // the process is over.
        if numDeleted == 0 {
            return nil
    	}

    	_, err = batch.Commit(r.Ctx)
    	if err != nil {
    		return err
    	}
    } else {
        batch := r.Client.Batch()

        // Get the Firestore path for the task
        taskidpath := fmt.Sprintf("tasks/%s", task.Id)

        taskDoc := r.Client.Doc(taskidpath)
        tdoc, err := taskDoc.Get(r.Ctx)
        if err != nil {
            e := fmt.Sprintf("DestroyTaskById: err getting task snapshot: %v", err)
            return errors.New(e)
        }
        batch.Delete(tdoc.Ref)

    	_, err = batch.Commit(r.Ctx)
    	if err != nil {
    		return err
    	}
    }

    return nil
} // }}}

// func ParseTaskFields {{{
func (r *Request) ParseTaskFields(fields url.Values, data map[string]interface{}) map[string]interface{} {
    //fmt.Printf("task fields: %v\n", fields)

    // Parse url fields
    for k, v := range fields {
        // Ensure the key is lower case
        k = strings.ToLower(k)

        // Our value is currently an array of strings; let's fix that
        val := strings.Join(v,"")

        // We want to check that the each key matches a field in
        // in the task to ensure we don't just add a bunch of new ones
        switch k {
        case "task_name":
            // I *probably* don't need to be checking this, cause it should
            // be passed to AddTask along with fields, not *in* fields
            data[k] = val
            break
        case "parent_id":
            // parent_id can be either a list_id OR a task_id
            data[k] = val
            break
        case "lock":
            // Unsure if we are even going to use this ..
            data[k], _ = strconv.ParseBool(val)
            break
        case "date_due":
            data[k], _ = time.Parse("01/02/2006 3:04:05 PM", val)
            break
        case "done":
            data[k], _ = strconv.ParseBool(val)
            break
        case "repeat":
            data[k] = val
            if val != NEVER {
                data["repeating"] = true
            }
            // Need function to update date_due at the repeat interval
            break
        case "end_repeat":
            data[k], _ = time.Parse("01/02/2006", val)
            break
        case "reminder":
            // I am going to set the time we need to remind them at right here
            // so we *MUST* be passed date_due BEFORE we are passed this.
            if val == NEVER {
                break
            }
            // Lets make an array of the words in our reminder
            reminder := strings.Split(val, " ")
            if len(reminder) == 4 {
                // Only way this could be the case is if it's "At time of event"
                data["reminder"] = ATOE
                data["reminder_time"] = data["date_due"]
                data["remind"] = true
                break
            }

            // So reminder must be some time before the event
            timeBefore, _ := strconv.Atoi(reminder[0])

            // Let's determine if it's minutes, days, or weeks before
            // which is indicated by the second word in the reminder
            interval := reminder[1]

            // We're only going to look at the first letter of the word
            i := interval[0]
            if i == 'd' {
                data["reminder"] = reminder[0] + DBE
                var remindTime time.Time
                due := data["date_due"].(time.Time)

                remindTime = due.AddDate(0, 0, -timeBefore)
                data["reminder_time"] = remindTime
            }

            if i == 'm' {
                data["reminder"] = reminder[0] + MBE
                var remindTime time.Time
                due := data["date_due"].(time.Time)
                var before time.Duration
                before = time.Duration(timeBefore)
                remindTime = due.Add(-before * time.Minute)
                data["reminder_time"] = remindTime
            }

            if i == 'w' {
                data["reminder"] = reminder[0] + WBE
                var remindTime time.Time
                due := data["date_due"].(time.Time)
                remindTime = due.AddDate(0, 0, -7 * timeBefore)
                data["reminder_time"] = remindTime
            }
            data["remind"] = true
            break
        case "remind_type":
            data[k] = val
            break
        case "priority":
            data[k] = val
            break
        case "location":
            data[k] = val
            break
        case "description":
            data[k] = val
            break
        case "url":
            data[k] = val
            break
        case "shared":
            data[k], _ = strconv.ParseBool(val)
            break
        case "shared_users":
            data[k] = val
            break
        case "sub_task":
            data[k], _ = strconv.ParseBool(val)
            break
        case "sub_tasks":
            data[k] = val
            break
        }
    }
    return data
} // }}}

// func TaskToJSON {{{
//
func (r *Request) TaskToJSON() *TaskJSON {
    var taskjson TaskJSON

    taskjson.Id          = r.Task.Id
    taskjson.Name        = r.Task.Name
    taskjson.Owner       = r.Task.Owner
    taskjson.Parent      = r.Task.Parent
    taskjson.Lock        = r.Task.Lock
    taskjson.DateDue     = r.Task.DateDue
    taskjson.Done        = r.Task.Done
    taskjson.Repeating   = r.Task.Repeating
    taskjson.Repeat      = r.Task.Repeat
    taskjson.EndRepeat   = r.Task.EndRepeat
    taskjson.Remind      = r.Task.Remind
    taskjson.RemindType  = r.Task.RemindType
    taskjson.Reminder    = r.Task.Reminder
    taskjson.RemindTime  = r.Task.RemindTime
    taskjson.Priority    = r.Task.Priority
    taskjson.Location    = r.Task.Location
    taskjson.Description = r.Task.Description
    taskjson.Url         = r.Task.Url
    taskjson.Shared      = r.Task.Shared
    taskjson.SharedUsers = r.Task.SharedUsers
    taskjson.Subtasks    = r.Task.Subtasks
    taskjson.Subtask     = r.Task.Subtask

    return &taskjson
} // }}}

func (r *Request) UpdateTaskSubtasks(taskid, id string) {
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

    // Add the new id to our subtask array
    task.Subtasks = append(task.Subtasks, id)

    // Make a map of the new subtasks to send to Firestore
    d := make(map[string]interface{})
    d["sub_tasks"] = task.Subtasks

    // Send update to Firestore
    _, err = doc.Set(r.Ctx, d, firestore.MergeAll)
    if err != nil {
        fmt.Printf("err setting new task data: %v\n", err)
    }
}
