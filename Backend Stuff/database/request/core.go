package request

import (
   "log"
   "context"

   "database/list"
   "database/user"

   "google.golang.org/api/option"
   "cloud.google.com/go/firestore"
)

// func DataBaseRequest {{{
func DataBaseRequest(uid, action, item string, desc map[string]interface{}) error {
   var r Request

   r.Action = action
   r.Item = item
   r.ItemName = desc["name"]
   r.ItemField = desc["field"]
   r.NewValue = desc["newval"]

   r.Ctx = context.Background()
   r.GetClient()
   defer r.Client.Close()

   r.User = user.GetUser(uid, r.Client, r.Ctx)

   return r.HandleRequest()
} // }}}


// func GetClient {{{
//
// Returns a firestore client so we can communicate to the database.
func (r *Request) GetClient() {
   opt := option.WithCredentialsFile("/Users/sabra/Library/Mobile Documents/com~apple~CloudDocs/School/COMP 482/Project1/todolist-react/Backend Stuff/database/conf/friday-584-firebase-adminsdk-es3qw-60b3f32bf1.json")
   client, err := firestore.NewClient(r.Ctx, "friday-584", opt)
	if err != nil {
		log.Fatalf("Cannot create client: %v", err)    // %v is to format error values
	}

   r.Client = client
} // }}}

// func HandleRequest {{{
func (r *Request) HandleRequest() error {
   switch r.Action {
   case "edit":
      if r.Item == "list" {
         list := list.GetListByName(r.ItemName.(string), r.User.Id, r.Client, r.Ctx)
         return list.Edit(r.ItemField.(string), r.NewValue, r.Client, r.Ctx)
      }
   }
   return nil
} // }}}
