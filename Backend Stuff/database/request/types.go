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

/*switch (req.Type) {
case "delete":
   // Get each item in the payload and delete it
case "update":
   // Get each item in the payload
   // if any value != database_value && value != null
   //    change value
   //    return true if all successful
case "read":
   // Get each item in the payload and return any field that was not omitted
case "create":
   // Create new object in firestore, return new object & its id
}*/

// Structure for the request
type Request struct {
   // Pointer to a struct representing the user the request is targeting
   UserId  string

   User *User
   List *List

   // The action to be performed : add, edit, delete
   Type    string

   Payload Payload

   // Firestore client for the session
   Client  *firestore.Client

   // Context for the session
   Ctx     context.Context
}

type Payload struct {
   Users []*User
   Lists []*List
   Tasks []*Task
}
