package notification_helper

import ( "time" )

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
    Email       bool        `firestore:"email"`
    Discord     bool        `firestore:"discord"`

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
    Name        string      `json:"text,omitempty"`

    // User ID of the user who owns this task
    Owner       string      `json:"task_owner,omitempty"`

    // ID of the parent list or parent task, if a subtask
    Parent      string      `json:"parent_id,omitempty"`

    // Whether or not someone can edit this task
    Lock        bool        `json:"lock,omitempty"`

    // Name of the list the task is in
    List        string      `json:"list,omitempty"`

    // Date this task is due (includes the time it is due)
    DateDue     time.Time   `json:"date,omitempty"`

    // Whether or not the task is complete / finished
    Done        bool        `json:"isComplete"`

    // Whether or not we should repeat this task, used for queries
    Repeating   bool        `json:"willRepeat"`

    // The frequency of the repeat, if we are repeating
    Repeat      string      `json:"repeatFrequency,omitempty"`

    // The date we should stop repeating this task
    EndRepeat   time.Time   `json:"end_repeat,omitempty"`

    // Whether or not we should remind the user, used for queries
    Remind      bool        `json:"remind"`

    // What type of reminder they want, discord or email
    Email       bool        `json:"emailSelected"`
    Discord     bool        `json:"discordSelected"`

    // Time frame before task to remind the user -- string
    Reminder    string      `json:"reminder,omitempty"`

    // The actual time we are going to remind the user
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

    // IDs of assoociated Subtasks
    Subtasks    []string    `json:"subTasks,omitempty"`
}
