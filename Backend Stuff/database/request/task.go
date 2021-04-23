package request

import (
   "time"
   "fmt"
   "log"
   "net/url"
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

func (r *Request) GetTasks(parentid string) []*Task {
   var tasks []*Task

   iter := r.Client.Collection("tasks").Where("task_owner", "==", r.UserId).Where("parent_id", "==", parentid).Documents(r.Ctx)
   for {
      docsnap, err := iter.Next()
      if err == iterator.Done {
         break
      }
      if err != nil {
         log.Printf("err: %v", err)
      }

      // create a new task struct
      var task Task

      // Put data into our task structure
      docsnap.DataTo(&task)

      // Get & set the task ID
      id := docsnap.Ref.ID
      task.Id = id

      tasks = append(tasks, &task)
   }

   return tasks
} // }}}

// func GetTaskByName {{{
//
// Returns a task using the tasks name
// Ensures we get the correct task by specifying the task owner
func (r *Request) GetTaskByName(taskname string) {
   var task Task
   iter := r.Client.Collection("tasks").Where("task_name", "==", taskname).Where("task_owner", "==", r.UserId).Documents(r.Ctx)
   for {
      docsnap, err := iter.Next()
      if err == iterator.Done {
         break
      }
      if err != nil {
         log.Printf("err: %v", err)
      }

      // Put data into our task structure
      docsnap.DataTo(&task)

      // Get & set the task ID
      id := docsnap.Ref.ID
      task.Id = id
   }

   r.Task = &task
} // }}}

func (r *Request) GeTaskByID() {
   var task Task

   // Get the Firestore path for the task
   taskidpath := fmt.Sprintf("tasks/%s", r.Task.Id)

   // Pass that to Firestore
   doc := r.Client.Doc(taskidpath)

   // Get a snapshot of the user data
   docsnap, err := doc.Get(r.Ctx)
   if err != nil {
      log.Printf("ERR Cannot get task by id snapshot: %v", err)    // %v is to format error values
   }

   // Add the data to our structure
   err = docsnap.DataTo(&task)
   if err != nil {
      log.Printf("ERR Cannot put task data to struct: %v", err)   // %v is to format error values
   }

   // Get & set the user ID
   id := docsnap.Ref.ID
   task.Id = id

   r.Task = &task
} // }}}

func (r *Request) AddTask(name string, fields url.Values) error {
   var task Task
   var data = make(map[string]interface{})

   for k, v := range fields {
      k = strings.ToLower(k)

      val := strings.Join(v,"")
      if k == "task_name" {
         if val != name {
            data[k] = val
         }
      }
      if k == "lock" {
         data[k], _ = strconv.ParseBool(val)
      }
   }

   if data["task_name"] != name {
      data["task_name"] = name
   }

   data["task_owner"] = r.UserId

   if data["sub_tasks"] == nil {
      var tasks []string
      tasks = append(tasks, "")
      data["sub_tasks"] = tasks
   }

   //fmt.Printf("%v\n", data)

   ref := r.Client.Collection("tasks").NewDoc()
   task.Id = ref.ID
   r.Task = &task

   _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
   if err != nil {
      // Handle any errors in an appropriate way, such as returning them.
      log.Printf("ERR adding new task: %v", err)
   }
   return err
}

// func Update {{{
func (r *Request) UpdateTask(fields url.Values) error {
   var data = make(map[string]interface{})
   //log.Printf("%v", fields)

   for k, v := range fields {
      k = strings.ToLower(k)

      val := strings.Join(v,"")
      if k == "task_name" {
         data[k] = val
      }
      if k == "lock" {
         data[k], _ = strconv.ParseBool(val)
      }
   }

   //log.Printf("%v", data)

   ref := r.Client.Collection("tasks").Doc(r.Task.Id)
   _,err := ref.Set(r.Ctx, data, firestore.MergeAll)

   return err
} // }}}

func (r *Request) DestroyTask() error {
   // Get the Firestore path for the user
   taskidpath := fmt.Sprintf("tasks/%s", r.Task.Id)

   // Delete that list
   _, err := r.Client.Doc(taskidpath).Delete(r.Ctx)

   return err
} // }}}
