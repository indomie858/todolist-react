package list

// Structure for list data
type List struct {
   // Firestore generated list ID
   Id          string   `firestore:"id,omitempty"`

   // Name of the list
   Name        string   `firestore:"list_name,omitempty"`

   // User ID of the user who owns this list
   Owner       string   `firestore:"list_owner,omitempty"`

   // Whether or not someone can edit this list
   Lock        bool     `firestore:"lock,omitempty"`

   // Whether or not this list is shared
   Shared      bool     `firestore:"shared,omitempty"`

   // Array of user IDs of the users this list has been shared with
   SharedUsers []string `firestore:"shared_users,omitempty"`

   // Array of task ids
   Tasks       []string `firestore:"tasks,omitempty"`
}
