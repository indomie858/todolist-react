package user
import (
   "log"
   "context"
   "fmt"

   "cloud.google.com/go/firestore"
)

// func GetUser {{{
//
// Returns a user from the Firestore database
func GetUser(userid string, client *firestore.Client, ctx context.Context) *User {
   var user User

   // Get the Firestore path for the user
   useridpath := fmt.Sprintf("users/%s", userid)

   // Pass that to Firestore
   doc := client.Doc(useridpath)

   // Get a snapshot of the user data
   docsnap, err := doc.Get(ctx)
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

   return &user
} // }}}
