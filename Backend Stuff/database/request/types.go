package request
import (
   "context"
   "cloud.google.com/go/firestore"
)

// Actions that can be performed
const (
   CREATE = "create"
   READ = "read"
   UPDATE = "update"
   DESTROY = "destroy"
)

// Items that can be modified/retrieved
const (
   CLIENT = "client"
   LIST = "list"
   SETTINGS = "settings"
   TASK = "task"
   USER = "user"
)

// Structure for the request
type Request struct {
   // ID of the user requesting database access
   UserId  string

   // Pointers to structure for the various documents we might need
   User *User
   List *List
   Task *Task

   // The action to be performed : add, edit, delete
   Type    string

   // Firestore client for the session
   Client  *firestore.Client

   // Context for the session
   Ctx     context.Context
}
