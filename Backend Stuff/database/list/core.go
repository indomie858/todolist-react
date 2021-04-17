package list

import (
   "log"
   "context"

   "google.golang.org/api/iterator"
   "cloud.google.com/go/firestore"
)

// func GetListByName {{{
//
// Returns a list using the list name
// Ensures we get the correct list by specifying the list owner
func GetListByName(listname, userId string, client *firestore.Client, ctx context.Context) *List {
   var list List
   iter := client.Collection("lists").Where("list_name", "==", listname).Where("list_owner", "==", userId).Documents(ctx)
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

   return &list
} // }}}

// func Edit() {{{
func (l *List) Edit(fieldname string, newval interface{}, client *firestore.Client, ctx context.Context) error {
   ref := client.Collection("lists").Doc(l.Id)
   err := client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
      /*doc, err := tx.Get(ref) // tx.Get, NOT ref.Get!
      if err != nil {
         return err
      }
      field, err := doc.DataAt(fieldname)
      if err != nil {
         return err
      }*/
      var input map[string]interface{}
      input = make(map[string]interface{})
      input[fieldname] = newval
      return tx.Set(ref, input, firestore.MergeAll)
   })

   return err
} // }}}
