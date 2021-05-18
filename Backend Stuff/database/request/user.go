package request

import (
    "fmt"
    "errors"
    "strings"
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

    // Client settings
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

    // Client settings
    Settings string `json:"settings,omitempty"`
}

// func AddUser {{{
//
func (r *Request) AddUser(name string, fields url.Values) (*UserJSON, error) {
    var ujson *UserJSON
    // Create a new doc & set our UserId to that doc's ID
    ref := r.Client.Collection("users").NewDoc()
    r.UserId = ref.ID

    // Parse the url fields into a map for Firestore
    data := ParseUserFields(fields)

    // If this wasn't passed in the payload, then let's create a default list array
    if data["lists"] == nil {
        // f is a url.Values variable, which is required for r.AddList
        f := url.Values{}
        list, _ := r.AddList("first_list", f)

        var lists []string
        lists = append(lists, list.Id)
        data["lists"] = lists
    }

    if data["name"] == nil {
        data["name"] = name
    }

    //fmt.Printf("%v\n", data)

    // Pass the field data to Firestore
    _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err setting new user data: %v", err)
        return ujson, errors.New(e)
    }

    ujson, _ = r.GetUser()
    return ujson, nil
} // }}}

// func GetUser {{{
//
// Returns a user from the Firestore database
func (r *Request) GetUser() (*UserJSON, error) {
    var ujson *UserJSON

    var user User

    // Get the Firestore path for the user
    useridpath := fmt.Sprintf("users/%s", r.UserId)

    // Pass that to Firestore
    doc := r.Client.Doc(useridpath)

    // Get a snapshot of the user data
    docsnap, err := doc.Get(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("err getting user snapshot: %v", err)
        return ujson, errors.New(e)
    }

    // Add the data to our structure
    err = docsnap.DataTo(&user)
    if err != nil {
        e := fmt.Sprintf("err putting user data to struct: %v", err)
        return ujson, errors.New(e)
    }

    // Get & set the user ID
    id := docsnap.Ref.ID
    user.Id = id

    // Set our request user to be this user
    r.User = &user
    ujson = r.UserToJSON()
    return ujson, nil
} // }}}

// func UpdateUser {{{
//
func (r *Request) UpdateUser(fields url.Values) (*UserJSON, error) {
    var ujson *UserJSON
    // Uncomment to see how the fields are formatted
    //fmt.Printf("%v", fields)

    // Parse the url fields into a map for Firestore
    data := ParseUserFields(fields)

    //fmt.Printf("%v", data)

    // Get a reference to our user document
    ref := r.Client.Collection("users").Doc(r.User.Id)

    // Send update to Firestore
    _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err setting new user data: %v", err)
        return ujson, errors.New(e)
    }

    // Update the user data in the request struct
    ujson, err = r.GetUser()
    return ujson, err
}  // }}}

// func DestroyUser {{{
//
func (r *Request) DestroyUser() error {
    // Get the Firestore path for the user
    useridpath := fmt.Sprintf("users/%s", r.UserId)

    user, _ := r.GetUser()
    if len(user.Lists) > 0 {
        for _, list := range user.Lists {
            r.DestroyListById(list)
        }
    }

    // Delete that list
    _, err := r.Client.Doc(useridpath).Delete(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("err deleting user: %v", err)
        return errors.New(e)
    }
    return nil
} // }}}

// func ParseUserFields {{{
//
func ParseUserFields(fields url.Values) map[string]interface{} {
    // log.Printf(fields)
    var data = make(map[string]interface{})

    for k, v := range fields {
        // Ensure the key is lower case
        k = strings.ToLower(k)

        // Our value is currently an array of strings; let's fix that
        val := strings.Join(v, "")

        // We want to check the key to ensure we don't just add a bunch of new fields
        if k == "lists" {
            data[k] = v
        } else if k == "tasks" {
            data[k] = v
        } else if k == "name" {
            data[k] = val
        }
    }

    return data
} // }}}

// func UserToJSON {{{
//
func (r *Request) UserToJSON() *UserJSON {
    var userjson UserJSON

    userjson.Id       = r.User.Id
    userjson.Name     = r.User.Name
    userjson.Email    = r.User.Email
    userjson.Status   = r.User.Status
    userjson.Lists    = r.User.Lists
    userjson.Settings = r.User.Settings

    return &userjson
} // }}}

func (r *Request) UpdateUserList(id string) {
    var user User

    // Get the Firestore path for the user
    useridpath := fmt.Sprintf("users/%s", r.UserId)

    // Pass that to Firestore
    doc := r.Client.Doc(useridpath)

    // Get a snapshot of the user data
    docsnap, err := doc.Get(r.Ctx)
    if err != nil {
        fmt.Printf("err getting user snapshot: %v\n", err)
        return
    }

    // Add the data to our structure
    err = docsnap.DataTo(&user)
    if err != nil {
        fmt.Printf("err putting user data to struct: %v\n", err)
        return
    }

    user.Lists = append(user.Lists, id)
    d := make(map[string]interface{})
    d["lists"] = user.Lists
    // Send update to Firestore
    _, err = doc.Set(r.Ctx, d, firestore.MergeAll)
    if err != nil {
        fmt.Printf("err setting new user data: %v\n", err)
    }
}
