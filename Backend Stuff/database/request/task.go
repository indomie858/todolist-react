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

// Structure for task data
type Task struct {
    // Firestore generated task ID
    Id          string    `firestore:"id,omitempty"`

    // Name of the task
    Name        string    `firestore:"task_name,omitempty"`

    // User ID of the user who owns this task
    Owner       string   `firestore:"task_owner,omitempty"`

    // ID of the parent list
    Parent      string   `firestore:"parent_id"`

    // Whether or not someone can edit this task
    Lock        bool      `firestore:"lock,omitempty"`

    // Date this task is due
    DueDate     time.Time `firestore:"due_date,omitempty"`

    // When the user would idelly like to start this task
    IdealStart  time.Time `firestore:"ideal_start_date"`

    // Date task was started
    StartDate   time.Time `firestore:"start_date,omitempty"`

    // Whether or not we should repeat this task
    Repeating   bool      `firestore:"repeating,omitempty"`

    // The frequency of the repat, if we are repeating
    Repeat      string    `firestore:"repeat,omitempty"`

    // Whether or not we should remind the user
    Remind      bool      `firestore:"remind,omitempty"`

    // Time frame before task to remind the user -- string
    Reminder    string    `firestore:"reminder,omitempty"`

    // Time frame before task to remind the user -- int
    TimeFrame   int       `firestore:"time_frame,omitempty"`

    // Location of the task
    Location    string    `firestore:"location,omitempty"`

    // Description of the task
    Description string    `firestore:"description,omitempty"`

    // Url associated with the task -- could be an array if desired
    Url         string    `firestore:"url,omitempty"`

    // IDs of assoociated Subtasks
    Subtasks    []string `firestore:"sub_tasks,omitempty"`
}

// Structure for json task data
type TaskJSON struct {
    // Firestore generated task ID
    Id          string    `json:"id,omitempty"`

    // Name of the task
    Name        string    `json:"task_name,omitempty"`

    // User ID of the user who owns this task
    Owner       string   `json:"task_owner,omitempty"`

    // ID of the parent list
    Parent      string   `json:"parent_id"`

    // Whether or not someone can edit this task
    Lock        bool      `json:"lock,omitempty"`

    // Date this task is due
    DueDate     time.Time `json:"due_date,omitempty"`

    // When the user would idelly like to start this task
    IdealStart  time.Time `json:"ideal_start_date"`

    // Date task was started
    StartDate   time.Time `json:"start_date,omitempty"`

    // Whether or not we should repeat this task
    Repeating   bool      `json:"repeating,omitempty"`

    // The frequency of the repat, if we are repeating
    Repeat      string    `json:"repeat,omitempty"`

    // Whether or not we should remind the user
    Remind      bool      `json:"remind,omitempty"`

    // Time frame before task to remind the user -- string
    Reminder    string    `json:"reminder,omitempty"`

    // Time frame before task to remind the user -- int
    TimeFrame   int       `json:"time_frame,omitempty"`

    // Location of the task
    Location    string    `json:"location,omitempty"`

    // Description of the task
    Description string    `json:"description,omitempty"`

    // Url associated with the task -- could be an array if desired
    Url         string    `json:"url,omitempty"`

    // IDs of assoociated Subtasks
    Subtasks    []string `json:"sub_tasks,omitempty"`
}

// AddTask {{{
//
func (r *Request) AddTask(name string, fields url.Values) (*TaskJSON, error) {
    var tjson *TaskJSON
    var task Task

    // Create new task document in Firestore
    ref := r.Client.Collection("tasks").NewDoc()
    task.Id = ref.ID

    data := ParseTaskFields(fields)

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

    // Send the parsed task values to Firstore
    _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err setting new task data: %v", err)
        return tjson, errors.New(e)
    }

    tjson, _ = r.GetTaskByName(name)
    return tjson, err
} // }}}

// func GetTaskByName {{{
//
// Returns a task using the tasks name
// Ensures we get the correct task by specifying the task owner
func (r *Request) GetTaskByName(taskname string) (*TaskJSON, error) {
    var tjson *TaskJSON
    var task Task
    // Get all tasks from Firestore where the owner is the requesting user and the task_name is the same as the one provided
    iter := r.Client.Collection("tasks").Where("task_name", "==", taskname).Where("task_owner", "==", r.UserId).Documents(r.Ctx)

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
    }

    // Set our request task to be this task
    r.Task = &task
    tjson = r.TaskToJSON()
    return tjson, nil
} // }}}

// func GetTaskByID {{{
//
func (r *Request) GetTaskByID() (*TaskJSON, error) {
    var tjson *TaskJSON
    var task Task

    // Get the Firestore path for the task
    taskidpath := fmt.Sprintf("tasks/%s", r.Task.Id)

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

    // Set our request task to be this task
    r.Task = &task
    tjson = r.TaskToJSON()
    return tjson, nil
} // }}}

// func GetTasks {{{
//
func (r *Request) GetTasks(parentid string) ([]*TaskJSON, error) {
   var tasks []*TaskJSON

   // Get all tasks from Firestore where the owner is the requesting user and the parent_id is the same as the one provided
   iter := r.Client.Collection("tasks").Where("task_owner", "==", r.UserId).Where("parent_id", "==", parentid).Documents(r.Ctx)

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

       // Get & set the task ID
       id := docsnap.Ref.ID
       task.Id = id
       r.Task = &task
       // Add task to the tasks array
       tasks = append(tasks, r.TaskToJSON())
    }

    return tasks, nil
} // }}}

// func UpdateTask {{{
//
func (r *Request) UpdateTask(fields url.Values) (*TaskJSON, error) {
    var tjson *TaskJSON
    //log.Printf("%v", fields)

    // Parse the url fields into a map for Firestore
    data := ParseTaskFields(fields)

    //log.Printf("%v", data)

    // Get a reference to our task
    ref := r.Client.Collection("tasks").Doc(r.Task.Id)

    // Send update to Firestore
    _,err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err updating task data: %v", err)
        return tjson, errors.New(e)
    }
    tjson, _ = r.GetTaskByID()
    return tjson, err
} // }}}

// func DestroyTasks {{{
//
//
// TODO: Delete all subtasks as well
func (r *Request) DestroyTask() error {
    // Get the Firestore path for the user
    taskidpath := fmt.Sprintf("tasks/%s", r.Task.Id)

    // Delete that list
    _, err := r.Client.Doc(taskidpath).Delete(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("err deleting task: %v", err)
        return errors.New(e)
    }
    return nil
} // }}}

// func ParseTaskFields {{{
func ParseTaskFields(fields url.Values) map[string]interface{} {
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
        } else if k == "lock" {
          data[k], _ = strconv.ParseBool(val)
        } else if k == "parent_id" {
            // parent_id can be either a list_id OR a task_id
            data[k] = val
        } else if k == "sub_task" {
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
    taskjson.Subtasks   = r.Task.Subtasks

    return &taskjson
} // }}}
