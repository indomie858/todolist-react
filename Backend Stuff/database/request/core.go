package request

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"database/list"
	"database/user"

	"github.com/joho/godotenv"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// func GetCredentials {{{
//
// Reads the credential variables from the env file, formatting them into the
// proper JSON format
// Returns that as an array of bytes to passed to Google's option.WithCredentialsJSON
func GetCredentials() []byte {
	t := os.Getenv("TYPE")
	pid := os.Getenv("PROJECT_ID")
	pkid := os.Getenv("PRIVATE_KEY_ID")
	pk := os.Getenv("PRIVATE_KEY")
	ce := os.Getenv("CLIENT_EMAIL")
	cid := os.Getenv("CLIENT_ID")
	au := os.Getenv("AUTH_URI")
	tu := os.Getenv("TOKEN_URI")
	ap := os.Getenv("AUTH_PROVIDER")
	cert := os.Getenv("CLIENT_X509_CERT_URL")

	json := fmt.Sprintf(`{
      "type": "%s",
      "project_id": "%s",
      "private_key_id": "%s",
      "private_key": "%s",
      "client_email": "%s",
      "client_id": "%s",
      "auth_uri": "%s",
      "token_uri": "%s",
      "auth_provider_x509_cert_url": "%s",
      "client_x509_cert_url": "%s"
   }`, t, pid, pkid, pk, ce, cid, au, tu, ap, cert)

	credentials := []byte(json)
	return credentials
} // }}}

// func GetClient {{{
//
// Returns a firestore client so we can communicate to the database.
func (r *Request) GetClient() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load env file: %v", err)
	}

	// Set our credential variables
	pid := os.Getenv("PROJECT_ID")
	credentials := GetCredentials()
	opt := option.WithCredentialsJSON(credentials)

	// Retrieve the firestore client
	client, err := firestore.NewClient(r.Ctx, pid, opt)
	if err != nil {
		log.Fatalf("Cannot create client: %v", err) // %v is to format error values
	}

	r.Client = client
} // }}}

// func HandleRequest {{{
//
// Handle the database request
func (r *Request) HandleRequest() error {
	switch r.Action {
	case ADD:
		return r.Add()
	case EDIT:
		return r.Edit()
	case GET:
		return r.Get()
	default:
		err := errors.New("ERR : Unknown action specified. Accepted values: add, delete, edit, get.")
		return err
	}
} // }}}

// func Add {{{
//
// Requests to add an item to collection are handled here
func (r *Request) Add() error {
	switch r.Item {
	case CLIENT:
		return nil
	case LIST:
		return nil
		//l := list.GetListByName(r.ItemName.(string), r.User.Id, r.Client, r.Ctx)
		//return l.Add(r.ItemField.(string), r.NewValue, r.Client, r.Ctx)
	case TASK:
		return nil
		//t := task.GetTaskByName(r.ItemName.(string), r.User.Id, r.Client, r.Ctx)
		//return t.Add(r.ItemField.(string), r.NewValue)
	case USER:
		return nil
		// return users.Add()
	default:
		err := errors.New("ERR : Unknown item specified. Accepted values: client, list, task, user.")
		return err
	}
}

// }}}

// func Get {{{
//
// Requests to get an item to collection are handled here
func (r *Request) Get() error {
	switch r.Item {
	case CLIENT:
		return nil
	case LIST:
		return nil
		//l := list.GetListByName(r.ItemName.(string), r.User.Id, r.Client, r.Ctx)
		//return l.GetField(r.ItemField.(string), r.NewValue, r.Client, r.Ctx)
	case TASK:
		return nil
		//t := task.GetTaskByName(r.ItemName.(string), r.User.Id, r.Client, r.Ctx)
		//return t.GetField(r.ItemField.(string), r.Client, r.Ctx)
	case USER:
		return nil
		// return r.User.Get(r.ItemField.(string), r.Client, r.Ctx)
	default:
		err := errors.New("ERR : Unknown item specified. Accepted values: client, list, task, user.")
		return err
	}
} // }}}

// func Delete {{{
//
// Requests to edit an item in a collection are handled here
func (r *Request) Delete() error {
	switch r.Item {
	case CLIENT:
		return nil
	case LIST:
		return nil
		//l := list.GetListByName(r.ItemName.(string), r.User.Id, r.Client, r.Ctx)
		//return l.Delete(r.ItemField.(string), r.NewValue, r.Client, r.Ctx)
	case TASK:
		return nil
		//t := task.GetTaskByName(r.ItemName.(string), r.User.Id, r.Client, r.Ctx)
		//return t.Edit(r.ItemField.(string), r.NewValue, r.Client, r.Ctx)
	case USER:
		return nil
		//return r.User.Edit(r.ItemField.(string), r.NewValue, r.Client, r.Ctx)
	default:
		err := errors.New("ERR : Unknown item specified. Accepted values: client, list, task, user.")
		return err
	}
}

// }}}

// func Edit {{{
//
// Requests to edit an item in a collection are handled here
func (r *Request) Edit() error {
	switch r.Item {
	case CLIENT:
		return nil
	case LIST:
		list := list.GetListByName(r.ItemName.(string), r.User.Id, r.Client, r.Ctx)
		return list.Edit(r.ItemField.(string), r.NewValue, r.Client, r.Ctx)
	case TASK:
		return nil
		//task := task.GetTaskByName(r.ItemName.(string), r.User.Id, r.Client, r.Ctx)
		//return task.Edit(r.ItemField.(string), r.NewValue, r.Client, r.Ctx)
	case USER:
		return nil
		//return r.User.Edit(r.ItemField.(string), r.NewValue, r.Client, r.Ctx)
	default:
		err := errors.New("ERR : Unknown item specified. Accepted values: client, list, task, user.")
		return err
	}
}

// }}}

// func DataBaseRequest {{{
//
// Requests to the database are passed here
func DataBaseRequest(uid, action, item string, desc map[string]interface{}) (Request, error) {

	// Create a new request
	var r Request

	// Set the request parameters
	r.Action = action
	r.Item = item
	r.ItemName = desc["name"]
	r.ItemField = desc["field"]
	r.NewValue = desc["newval"]

	// Set the context and retrieve the firestore client
	r.Ctx = context.Background()
	r.GetClient()
	defer r.Client.Close()

	// Get the user the request is targeting
	r.User = user.GetUser(uid, r.Client, r.Ctx)
	var err = r.HandleRequest()
	// Handle the request and return the resulting error
	return r, err
} // }}}
