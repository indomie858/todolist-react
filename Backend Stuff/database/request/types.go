package request
import (
   "context"
   "database/user"
   "cloud.google.com/go/firestore"
)
// Structure for the request
type Request struct {
   User      *user.User
   Action    string
   Item      string
   ItemName  interface{}
   ItemField interface{}
   NewValue  interface{}
   Client    *firestore.Client
   Ctx       context.Context
}
