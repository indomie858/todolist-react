package request

import (
    "fmt"
    "errors"
    "strings"
    "net/url"

    "cloud.google.com/go/firestore"
)

// Structure for user data
type User struct {
    // Firestore generated user ID
    Id              string   `firestore:"id,omitempty"`

    // First name of the user
    FirstName       string   `firestore:"first_name,omitempty"`

    // Last name of the user
    LastName        string   `firestore:"last_name,omitempty"`

    // The users major
    Major           string   `firestore:"major,omitempty"`
}

type UserJSON struct {
    // Firestore generated user ID
    Id              string   `json:"id,omitempty"`

    // First name of the user
    FirstName       string   `json:"first_name,omitempty"`

    // Last name of the user
    LastName        string   `json:"last_name,omitempty"`

    // The users major
    Major           string   `json:"major,omitempty"`
}

// func AddUser {{{
//
func (r *Request) AddUser(firstname, lastname string, fields url.Values) (*UserJSON, error) {
    var ujson *UserJSON
    // Create a new doc & set our UserId to that doc's ID
    ref := r.Client.Collection("hello-world-users").NewDoc()
    r.UserId = ref.ID

    // Parse the url fields into a map for Firestore
    // data := ParseUserFields(fields)

    var data = make(map[string]interface{})
    data["first_name"] = firstname
    data["last_name"] = lastname

    //fmt.Printf("%v\n", data)

    // Pass the field data to Firestore
    _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("AddUser ERR - err setting new user data: %v", err)
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

    // Create a string for the Firestore path for the user
    useridpath := fmt.Sprintf("hello-world-users/%s", r.UserId)

    // Pass that to Firestore to retrieve a document reference
    doc := r.Client.Doc(useridpath)

    // Get a document snapshot of the user data - what it was at the time you requested it
    docsnap, err := doc.Get(r.Ctx)

    // Handle any errors that may have occurred
    if err != nil {
        e := fmt.Sprintf("GetUser ERR - err getting user snapshot: %v", err)
        return ujson, errors.New(e)
    }

    // Add the data from the snapshot to our structure
    err = docsnap.DataTo(&user)

    // Handle any errors that may have occurred
    if err != nil {
        e := fmt.Sprintf("GetUser ERR - err putting user data to struct: %v", err)
        return ujson, errors.New(e)
    }

    // Get & set the user ID
    id := docsnap.Ref.ID

    // Set our user objects Id
    user.Id = id

    // Set our request user to be this user
    r.User = &user

    // Get the JSON version of this user
    ujson = r.UserToJSON()

    return ujson, nil
} // }}}

// func UpdateUser {{{
//
func (r *Request) UpdateUser(major string) (*UserJSON, error) {
    var ujson *UserJSON
    // Uncomment to see how the fields are formatted
    //fmt.Printf("%v", fields)

    // Parse the url fields into a map for Firestore
    //data := ParseUserFields(fields)

    var data = make(map[string]interface{})
    data["major"] = major

    //fmt.Printf("%v", data)

    // Get a reference to our user document
    ref := r.Client.Collection("hello-world-users").Doc(r.User.Id)

    // Send the update to Firestore
    _, err := ref.Set(r.Ctx, data, firestore.MergeAll)

    // Handle any errors that may have occurred
    if err != nil {
        e := fmt.Sprintf("err setting new user data: %v", err)
        return ujson, errors.New(e)
    }

    // Get a JSON version of this users data
    ujson, err = r.GetUser()
    return ujson, err
}  // }}}

// func DestroyUser {{{
//
func (r *Request) DestroyUser() error {
    // Get the Firestore path for the user
    useridpath := fmt.Sprintf("hello-world-users/%s", r.UserId)

    // Delete that user
    _, err := r.Client.Doc(useridpath).Delete(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("DestroyUser ERR - err deleting user: %v", err)
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

        // We want to check the key to ensure we don't just add a bunch of new fields,
        // since we update data by doing using a 'set' function that merges all the data
        switch k {
        case "first_name":
            data[k] = val
            break
        case "last_name":
            data[k] = val
            break
        case "major":
            data[k] = val
            break
        }
    }

    return data
} // }}}

// func UserToJSON {{{
//
// Returns a JSON encoded user structure
func (r *Request) UserToJSON() *UserJSON {
    var userjson UserJSON

    userjson.Id         = r.User.Id
    userjson.FirstName  = r.User.FirstName
    userjson.LastName   = r.User.LastName
    userjson.Major      = r.User.Major

    return &userjson
} // }}}
