package request

import (
   "fmt"
   "log"
   "net/url"
   "strings"
   "strconv"

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

func (r *Request) GetLists() []*List {
   var lists []*List

   iter := r.Client.Collection("lists").Where("list_owner", "==", r.UserId).Documents(r.Ctx)
   for {
      docsnap, err := iter.Next()
      if err == iterator.Done {
         break
      }
      if err != nil {
         log.Printf("err: %v", err)
      }

      // create a new list
      var list List

      // Put data into our list structure
      docsnap.DataTo(&list)

      // Get & set the list ID
      id := docsnap.Ref.ID
      list.Id = id

      lists = append(lists, &list)
   }

   return lists
} // }}}

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
         log.Printf("err: %v", err)
      }

      // Put data into our list structure
      docsnap.DataTo(&list)

      // Get & set the list ID
      id := docsnap.Ref.ID
      list.Id = id
   }

   r.List = &list
} // }}}

func (r *Request) GetListByID() {
   var list List

   // Get the Firestore path for the user
   listidpath := fmt.Sprintf("lists/%s", r.List.Id)

   // Pass that to Firestore
   doc := r.Client.Doc(listidpath)

   // Get a snapshot of the user data
   docsnap, err := doc.Get(r.Ctx)
   if err != nil {
      log.Printf("ERR: Cannot get list by id snapshot: %v", err)    // %v is to format error values
   }

   // Add the data to our structure
   err = docsnap.DataTo(&list)
   if err != nil {
      log.Printf("ERR: Cannot put data to struct: %v", err)   // %v is to format error values
   }

   // Get & set the user ID
   id := docsnap.Ref.ID
   list.Id = id

   r.List = &list
} // }}}

func (r *Request) AddList(name string, fields url.Values) error {
   var list List
   var data = make(map[string]interface{})

   for k, v := range fields {
      k = strings.ToLower(k)

      val := strings.Join(v,"")
      if k == "list_name" {
         if val != name {
            data[k] = val
         }
      }
      if k == "lock" {
         data[k], _ = strconv.ParseBool(val)
      }
   }

   data["list_name"] = name
   data["list_owner"] = r.UserId

   if data["tasks"] == nil {
      var tasks []string
      tasks = append(tasks, "")
      data["tasks"] = tasks
   }

   //fmt.Printf("%v\n", data)

   ref := r.Client.Collection("lists").NewDoc()
   list.Id = ref.ID
   r.List = &list

   _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
   if err != nil {
      // Handle any errors in an appropriate way, such as returning them.
      log.Printf("ERR adding new list: %v", err)
   }
   return err
}

// func Update {{{
func (r *Request) UpdateList(fields url.Values) error {
   var data = make(map[string]interface{})
   //log.Printf("%v", fields)

   for k, v := range fields {
      k = strings.ToLower(k)

      val := strings.Join(v,"")
      if k == "list_name" {
         data[k] = val
      }
      if k == "lock" {
         data[k], _ = strconv.ParseBool(val)
      }
   }

   //log.Printf("%v", data)

   ref := r.Client.Collection("lists").Doc(r.List.Id)
   _,err := ref.Set(r.Ctx, data, firestore.MergeAll)

   return err
} // }}}

func (r *Request) DestroyList() error {
   // Get the Firestore path for the user
   listidpath := fmt.Sprintf("lists/%s", r.List.Id)

   // Delete that list
   _, err := r.Client.Doc(listidpath).Delete(r.Ctx)

   return err
} // }}}
