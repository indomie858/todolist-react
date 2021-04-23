package request

import (
    "fmt"
    "net/url"
    "strings"
    "strconv"
    "errors"
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

// func AddList {{{
//
func (r *Request) AddList(name string, fields url.Values) (*ListJSON, error) {
    var ljson *ListJSON
    var list List
    // Create a new document
    ref := r.Client.Collection("lists").NewDoc()
    list.Id = ref.ID
    r.List = &list

    data := ParseListFields(fields)

    data["list_name"] = name
    data["list_owner"] = r.UserId

    // If our tasks array is empty, lets create a default one
    if data["tasks"] == nil {
        f := url.Values{}
        f.Add("parent_id", list.Id)
        f.Add("sub_task", "false")

        r.AddTask("first_task", f)
        r.GetTaskByName("first_task")

        var tasks []string
        tasks = append(tasks, r.Task.Id)
        data["tasks"] = tasks
    }

    //fmt.Printf("%v\n", data)

    // Set the data in our new document to the provided data
    _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err setting new list data: %v", err)
        return ljson, errors.New(e)
    }

    r.GetListByID()
    ljson = r.ListToJSON()
    return ljson, nil
} // }}}

// func GetListByName {{{
//
// Returns a list using the list name
// Ensures we get the correct list by specifying the list owner
func (r *Request) GetListByName(listname string) (*ListJSON, error) {
    var ljson *ListJSON
    var list List

    // Get all lists from Firestore where the owner is the requesting user and the list_name is the same as the one provided
    iter := r.Client.Collection("lists").Where("list_name", "==", listname).Where("list_owner", "==", r.UserId).Documents(r.Ctx)

    // For each document
    for {
        // Get a snapshot of the data
        docsnap, err := iter.Next()

        // Check if we're done with our loop
        if err == iterator.Done {
            break
        }

        // Check if we have some other error
        if err != nil {
            e := fmt.Sprintf("err getting list snapshot: %v", err)
            return ljson, errors.New(e)
        }

        // Put data into our list structure
        docsnap.DataTo(&list)

        // Get & set the list ID
        id := docsnap.Ref.ID
        list.Id = id
    }

    // Set our request list to be this list
    r.List = &list

    ljson = r.ListToJSON()
    return ljson, nil
} // }}}

// func GetListByID {{{
//
func (r *Request) GetListByID() (*ListJSON, error) {
    var ljson *ListJSON
    var list List

    // Get the Firestore path for the user
    listidpath := fmt.Sprintf("lists/%s", r.List.Id)

    // Pass that to Firestore
    doc := r.Client.Doc(listidpath)

    // Get a snapshot of the user data
    docsnap, err := doc.Get(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("err gtting list snapshot: %v", err)
        return ljson, errors.New(e)
    }

    // Add the data to our structure
    err = docsnap.DataTo(&list)
    if err != nil {
        e := fmt.Sprintf("err putting data to list  struct: %v", err)
        return ljson, errors.New(e)
    }

    // Get & set the user ID
    id := docsnap.Ref.ID
    list.Id = id

    // Set our request list to be this list
    r.List = &list
    ljson = r.ListToJSON()
    return ljson, nil
} // }}}

// func GetLists {{{
//
func (r *Request) GetLists() ([]*ListJSON, error) {
    var lists []*ListJSON
    // Get all lists from Firestore where the owner is the requesting user
    iter := r.Client.Collection("lists").Where("list_owner", "==", r.UserId).Documents(r.Ctx)

    // For each document
    for {
        // Get a snapshot of the data
        docsnap, err := iter.Next()

        // Check if we're done with our loop
        if err == iterator.Done {
            break
        }

         // Check if we have some other error
        if err != nil {
            e := fmt.Sprintf("err geting snapshot of list: %v", err)
            return lists, errors.New(e)
        }

        // Create a new list
        var list List

        // Put data into our list structure
        docsnap.DataTo(&list)

        // Get & set the list ID
        id := docsnap.Ref.ID
        list.Id = id
        r.List = &list

        // Add list to the lists array
        lists = append(lists, r.ListToJSON())
    }

    return lists, nil
} // }}}

// func UpdateList {{{
//
func (r *Request) UpdateList(fields url.Values) (*ListJSON, error) {
    var ljson *ListJSON
    //log.Printf("%v", fields)
    // Parse the url fields into a map for Firestore
    data := ParseListFields(fields)
    //log.Printf("%v", data)

    // Get a reference to our
    ref := r.Client.Collection("lists").Doc(r.List.Id)

    // Send update to Firestore
    _,err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err updating list data: %v", err)
        return ljson, errors.New(e)
    }
    ljson, _ = r.GetListByID()
    return ljson, nil
} // }}}

// func DestroyList {{{
//
// TODO: Delete all tasks as well
func (r *Request) DestroyList() error {
    // Get the Firestore path for the user
    listidpath := fmt.Sprintf("lists/%s", r.List.Id)

    // Delete that list
    _, err := r.Client.Doc(listidpath).Delete(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("err deleting list: %v", err)
        return errors.New(e)
    }
    return nil
} // }}}

// func ParseListFields {{{
//
func ParseListFields(fields url.Values) map[string]interface{} {
    var data = make(map[string]interface{})

    // Parse the url fields into a map for Firestore
    for k, v := range fields {
        // Ensure the key is lower case
        k = strings.ToLower(k)

        // Our value is currently an array of strings; let's fix that
        val := strings.Join(v,"")

        // We want to check the key to ensure we don't just add a bunch of new fields
        if k == "list_name" {
            data[k] = val
        } else if k == "lock" {
            data[k], _ = strconv.ParseBool(val)
        } else if k == "tasks" {
            data[k] = v
        } else if k == "list_owner" {
            data[k] = val
        } else if k == "shared" {
            data[k], _ = strconv.ParseBool(val)
        } else if k == "shared_users" {
            data[k] = v
        }
    }
    return data
} // }}}

// func ListToJSON {{{
//
func (r *Request) ListToJSON() *ListJSON {
    var listjson ListJSON

    listjson.Id          = r.List.Id
    listjson.Name        = r.List.Name
    listjson.Owner       = r.List.Owner
    listjson.Lock        = r.List.Lock
    listjson.Shared      = r.List.Shared
    listjson.SharedUsers = r.List.SharedUsers
    listjson.Tasks       = r.List.Tasks

    return &listjson
} // }}}
