package request
import (
   "context"
   "database/user"
   "cloud.google.com/go/firestore"
)

// Actions that can be performed
const (
   ADD = "add"
   GET = "get"
   EDIT = "edit"
   DELETE = "delete"
)

// Items that can be modified/retrieved
const (
   CLIENT = "client"
   LIST = "list"
   TASK = "task"
   USER = "user"
)

// Structure for the request
type Request struct {
   // Pointer to a struct representing the user the request is targeting
   User      *user.User

   // The action to be performed : add, edit, delete
   Action    string

   // Item the requesting is targeting
   Item      string

   // Name of the item
   ItemName  interface{}

   // Field to be edited
   ItemField interface{}

   // New value of the field, if applicable -- only used in edit
   NewValue  interface{}

   // Firestore client for the session
   Client    *firestore.Client

   // Context for the session
   Ctx       context.Context
}
