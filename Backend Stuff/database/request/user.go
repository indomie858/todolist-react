package request

import (
   "log"
   "fmt"
)

// Status setting constants
const (
   BUSY = "busy"
   FREE = "free"
)

// Structure for user data
type User struct {
   // Firestore generated user ID
   Id     string   `firestore:"id,omitempty"`

   // Name of the user
   Name   string   `firestore:"name,omitempty"`

   // Email of the user -- could possibly be an array if desired
   Email  string   `firestore:"email,omitempty"`

   // Status to show other users
   Status string   `firestore:"status,omitempty"`

   // Array of list ids
   Lists  []string `firestore:"lists,omitempty"`

   Settings string `firestore:"settings,omitempty"`
}

type UserJSON struct {
   // Firestore generated user ID
   Id     string   `json:"id,omitempty"`

   // Name of the user
   Name   string   `json:"name,omitempty"`

   // Email of the user -- could possibly be an array if desired
   Email  string   `json:"email,omitempty"`

   // Status to show other users
   Status string   `json:"status,omitempty"`

   // Array of list ids
   Lists  []string `json:"lists,omitempty"`

   Settings string `json:"settings,omitempty"`
}

/*func AddUser() {

}*/

// func GetUser {{{
//
// Returns a user from the Firestore database
func (r *Request) GetUser() {
   var user User

   // Get the Firestore path for the user
   useridpath := fmt.Sprintf("users/%s", r.UserId)

   // Pass that to Firestore
   doc := r.Client.Doc(useridpath)

   // Get a snapshot of the user data
   docsnap, err := doc.Get(r.Ctx)
   if err != nil {
      log.Fatalf("Cannot get user snapshot: %v", err)    // %v is to format error values
   }

   // Add the data to our structure
   err = docsnap.DataTo(&user)
   if err != nil {
      log.Fatalf("Cannot put data to struct: %v", err)   // %v is to format error values
   }

   // Get & set the user ID
   id := docsnap.Ref.ID
   user.Id = id

   r.User = &user
} // }}}
