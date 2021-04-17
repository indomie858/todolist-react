package user

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
}
