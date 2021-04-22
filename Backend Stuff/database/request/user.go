package request

import (
   "log"
   "fmt"
   "strings"
   //"strconv"
   "net/url"
   "cloud.google.com/go/firestore"
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
      log.Printf("ERR: Cannot get user snapshot: %v", err)    // %v is to format error values
   }

   // Add the data to our structure
   err = docsnap.DataTo(&user)
   if err != nil {
      log.Printf("ERR: Cannot put data to struct: %v", err)   // %v is to format error values
   }

   // Get & set the user ID
   id := docsnap.Ref.ID
   user.Id = id

   r.User = &user
} // }}}

func (r *Request) AddUser(name string, fields url.Values) error {
   var data = make(map[string]interface{})

   for k, v := range fields {
      k = strings.ToLower(k)
      val := strings.Join(v, "")
      // We want to check the key to ensure we don't just add a bunch of new fields
      if k == "lists" {
         var lists []string
         lists = append(lists, val)
         data[k] = v
      }
      if k == "tasks" {
         var tasks []string
         tasks = append(tasks, val)
         data[k] = tasks
      }
   }

   if data["lists"] == nil {
      var lists []string
      lists = append(lists, "")
      data["lists"] = lists
   }

   if data["tasks"] == nil {
      var tasks []string
      tasks = append(tasks, "")
      data["tasks"] = tasks
   }

   if data["name"] == nil {
      data["name"] = name
   }

   //fmt.Printf("%v\n", data)

   ref := r.Client.Collection("users").NewDoc()
   r.UserId = ref.ID

   _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
   if err != nil {
      // Handle any errors in an appropriate way, such as returning them.
      log.Printf("ERR updating user data err : %s", err)
   }

   r.GetUser()
   return err
}

// func Update {{{
func (r *Request) UpdateUser(fields url.Values) error {
   var data = make(map[string]interface{})

   // Uncomment to see how the fields are formatted
   //fmt.Printf("%v", fields)

   for k, v := range fields {
      k = strings.ToLower(k)

      val := strings.Join(v,"")
      if k == "lists" {
         var lists []string
         lists = append(lists, val)
         data[k] = lists
      }
      if k == "tasks" {
         var tasks []string
         tasks = append(tasks, val)
         data[k] = tasks
      }
   }

   // Uncomment to seee how the data to firestore is formatted
   //fmt.Printf("%v", data)

   ref := r.Client.Collection("users").Doc(r.User.Id)
   _, err := ref.Set(r.Ctx, data, firestore.MergeAll)

   r.GetUser()
   return err
} // }}}

func (r *Request) DestroyUser() error {
   // Get the Firestore path for the user
   useridpath := fmt.Sprintf("users/%s", r.UserId)

   // Delete that list
   _, err := r.Client.Doc(useridpath).Delete(r.Ctx)

   return err
} // }}}
