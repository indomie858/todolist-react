package request

import (
   "time"
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
}
