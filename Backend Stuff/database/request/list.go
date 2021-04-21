package request

import (
   "log"
   "context"

   "google.golang.org/api/iterator"
   "cloud.google.com/go/firestore"
)

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

// Structure for list data
type ListJSON struct {
   // Firestore generated list ID
   Id          string   `json:"id,omitempty"`

   // Name of the list
   Name        string   `json:"list_name,omitempty"`

   // User ID of the user who owns this list
   Owner       string   `json:"list_owner,omitempty"`

   // Whether or not someone can edit this list
   Lock        bool     `json:"lock,omitempty"`

   // Whether or not this list is shared
   Shared      bool     `json:"shared,omitempty"`

   // Array of user IDs of the users this list has been shared with
   SharedUsers []string `json:"shared_users,omitempty"`

   // Array of task ids
   Tasks       []string `json:"tasks,omitempty"`
}

// func GetListByName {{{
//
// Returns a list using the list name
// Ensures we get the correct list by specifying the list owner
func (r *Request) GetListByName(listname string) {
   var list List
   iter := r.Client.Collection("lists").Where("list_name", "==", listname).Where("list_owner", "==", r.UserId).Documents(r.Ctx)
   for {
      docsnap, err := iter.Next()
      if err == iterator.Done {
         break
      }
      if err != nil {
         log.Fatalf("err: %v", err)
      }

      // Put data into our list structure
      docsnap.DataTo(&list)

      // Get & set the list ID
      id := docsnap.Ref.ID
      list.Id = id
   }

   r.List = &list
} // }}}

// func Update {{{
func (r *Request) Update(fields interface{}) error {

   ref := r.Client.Collection("lists").Doc(r.List.Id)
   err := r.Client.RunTransaction(r.Ctx, func(ctx context.Context, tx *firestore.Transaction) error {
      /*doc, err := tx.Get(ref) // tx.Get, NOT ref.Get!
      if err != nil {
         return err
      }

      field, err := doc.DataAt(fieldname)
      if err != nil {
         return err
      }*/
      return tx.Set(ref, fields, firestore.MergeAll)
   })

   return err
} // }}}
