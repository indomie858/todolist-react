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

// Structure of the documents in the lists collection
// encoding (`firestore:"..."`) is firestore so we can
// easily dump requested data into this structure for
// easy access later.
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

// Structure of the documents in the list collection
// encoding (`json:"..."`) is json so we can pass the
// structure to the server in the correct json format.
// List data will be transferred over from List struct.
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
// Adds a list to the list collection, setting any fields provided
// Returns the newly added list in JSON format and nil if no errors
// occurs, returns the error and null ListJSON if an error occurss
func (r *Request) AddList(name string, fields url.Values) (*ListJSON, error) {
    var ljson *ListJSON
    var list List

    // Create a new document
    ref := r.Client.Collection("lists").NewDoc()
    list.Id = ref.ID
    //fmt.Printf("listId: %s\n", list.Id)

    // Create a new map for the list data
    var data = make(map[string]interface{})

    // Let's set some default values real quick -
    var sharedUsers []string
    sharedUsers = append(sharedUsers, "")

    data["list_name"] = name
    data["list_owner"] = r.UserId
    data["lock"] = false
    data["shared"] = false
    data["shared_users"] = sharedUsers

    // Now let's update our map to reflect the values we were given
    data = ParseListFields(fields, data)

    // If our tasks array is empty, lets create a default one
    if data["tasks"] == nil {
        f := url.Values{}
        task, _ := r.AddTask("first_task", list.Id, f)

        var tasks []string
        tasks = append(tasks, task.Id)
        data["tasks"] = tasks
    }

    //fmt.Printf("%v\n", data)

    // Set the data in our new document to the provided data
    _, err := ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err setting new list data: %v", err)
        return ljson, errors.New(e)
    }

    r.GetListByName(name)
    ljson = r.ListToJSON()
    if name != "first_list" {
        r.UpdateUserList(ref.ID)
    }

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
            e := fmt.Sprintf("GetListByName: err getting list snapshot: %v", err)
            return ljson, errors.New(e)
        }

        // Put data into our list structure
        docsnap.DataTo(&list)

        // Get & set the list ID
        id := docsnap.Ref.ID
        list.Id = id
        // Set our request list to be this list
        ljson, _ = r.GetListByID(id)
    }

    return ljson, nil
} // }}}

// func GetListByID {{{
//
func (r *Request) GetListByID(lid string) (*ListJSON, error) {
    var ljson *ListJSON
    var list List

    // Get the Firestore path for the user
    listidpath := fmt.Sprintf("lists/%s", lid)

    // Pass that to Firestore
    doc := r.Client.Doc(listidpath)

    // Get a snapshot of the user data
    docsnap, err := doc.Get(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("GetListByID: err getting list snapshot: %v", err)
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

    if list.Owner != r.UserId {
        if list.SharedUsers == nil || !r.CheckIfShared(list.SharedUsers){
            return ljson, errors.New("err getting list: requestor does not have permission")
        }
    }

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
            e := fmt.Sprintf("GetLists: err geting snapshot of list: %v", err)
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

// func GetSharedLists {{{
//
// Returns all lists shared with the current user
func (r *Request) GetSharedLists() ([]*ListJSON, error) {
    var lists []*ListJSON

    // Get all lists from Firestore where the shared is true and owner != requestor
    iter := r.Client.Collection("lists").Where("shared", "==", true).Where("list_owner", "!=", r.UserId).Documents(r.Ctx)

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
            e := fmt.Sprintf("GetSharedLists: err geting snapshot of list: %v", err)
            return lists, errors.New(e)
        }

        // Create a new list
        var list List

        // Put data into our list structure
        docsnap.DataTo(&list)
        if list.SharedUsers == nil || !r.CheckIfShared(list.SharedUsers) {
            continue
        }
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
func (r *Request) UpdateList(id string, fields url.Values) (*ListJSON, error) {
    ljson, err := r.GetListByID(id)
    fmt.Printf("ljsn: %s\n", ljson)
    if err != nil {
        e := fmt.Sprintf("err getting list for update: %v", err)
        return ljson, errors.New(e)
    }

    if ljson.Owner != r.UserId {
        if ljson.SharedUsers == nil || !r.CheckIfShared(ljson.SharedUsers)  {
            return ljson, errors.New("err updating list: requestor does not have permission")
        }
    }

    //log.Printf("%v", fields)
    // Parse the url fields into a map for Firestore
    var data = make(map[string]interface{})
    data = ParseListFields(fields, data)

    //log.Printf("%v", data)

    // Get a reference to our
    ref := r.Client.Collection("lists").Doc(ljson.Id)

    // Send update to Firestore
    _,err = ref.Set(r.Ctx, data, firestore.MergeAll)
    if err != nil {
        e := fmt.Sprintf("err updating list data: %v", err)
        return ljson, errors.New(e)
    }
    if data["list_name"] != nil {
        ljson, err = r.GetListByName(data["list_name"].(string))
        return ljson, err
    }
    ljson, err = r.GetListByID(id)
    return ljson, err
} // }}}

// func DestroyList {{{
//
func (r *Request) DestroyList(name string) error {
    // Get the list by name
    list, err := r.GetListByName(name)
    if err != nil {
        e := fmt.Sprintf("err getting list for delete: %v", err)
        return errors.New(e)
    }
    if list.Owner != r.UserId {
        if list.SharedUsers == nil || !r.CheckIfShared(list.SharedUsers) {
            return errors.New("err deleting list: requestor does not have permission")
        }
    }

    // Check if the list has any tasks
    if len(list.Tasks) > 0 {
        for _, task := range list.Tasks {
            //fmt.Printf("Task to Delete: %s\n", task)
            r.DestroyTaskById(task)
        }
    }

    batch := r.Client.Batch()

    // Get the Firestore path for the user
    listidpath := fmt.Sprintf("lists/%s", list.Id)

    listDoc := r.Client.Doc(listidpath)
    ldoc, err := listDoc.Get(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("DestroyList: err getting list snapshot: %v", err)
        return errors.New(e)
    }

    batch.Delete(ldoc.Ref)
    _, err = batch.Commit(r.Ctx)
    if err != nil {
        return err
    }

    return nil
} // }}}

// func DestroyListById {{{
//
func (r *Request) DestroyListById(id string) error {
    list, err := r.GetListByID(id)
    //fmt.Printf("ListId: %s\n", list.Id)
    if err != nil {
        e := fmt.Sprintf("DestroyListById: err getting list for delete: %v", err)
        return errors.New(e)
    }

    if list.Owner != r.UserId {
        if list.SharedUsers == nil || !r.CheckIfShared(list.SharedUsers) {
            return errors.New("DestroyListById: err deleting list: requestor does not have permission")
        }
    }

    // Check if the list has any tasks
    if len(list.Tasks) > 0 {
        for _, task := range list.Tasks {
            //fmt.Printf("Task to Delete: %s\n", task)
            r.DestroyTaskById(task)
        }
    }

    batch := r.Client.Batch()

    // Get the Firestore path for the user
    listidpath := fmt.Sprintf("lists/%s", list.Id)

    listDoc := r.Client.Doc(listidpath)
    ldoc, err := listDoc.Get(r.Ctx)
    if err != nil {
        e := fmt.Sprintf("DestroyListById: err getting list snapshot: %v", err)
        return errors.New(e)
    }
    batch.Delete(ldoc.Ref)
    _, err = batch.Commit(r.Ctx)
    if err != nil {
        return err
    }

    return nil
} // }}}

// func ParseListFields {{{
//
func ParseListFields(fields url.Values, data map[string]interface{}) map[string]interface{} {
    // Parse the url fields into a map for Firestore
    for k, v := range fields {
        // Ensure the key is lower case
        k = strings.ToLower(k)

        // Our value is currently an array of strings; let's fix that
        val := strings.Join(v,"")

        // We want to check the key to ensure we don't just add a bunch of new fields
        switch k {
        case "list_name":
            data[k] = val
            break
        case "lock":
            data[k], _ = strconv.ParseBool(val)
            break
        case "tasks":
            data[k] = v
            break
        case "list_owner":
            data[k] = val
            break
        case "shared":
            data[k], _ = strconv.ParseBool(val)
            break
        case "shared_users":
            data[k] = v
            break
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


func (r *Request) UpdateListTasks(listid, id string) {
    var list List

    // Get the Firestore path for the user
    listidpath := fmt.Sprintf("lists/%s", listid)

    // Pass that to Firestore
    doc := r.Client.Doc(listidpath)

    // Get a snapshot of the user data
    docsnap, err := doc.Get(r.Ctx)
    if err != nil {
        fmt.Printf("UpdateListTasks: err getting list snapshot: %v\n", err)
        return
    }

    // Add the data to our structure
    err = docsnap.DataTo(&list)
    if err != nil {
        fmt.Printf("err putting list data to struct: %v\n", err)
        return
    }

    list.Tasks = append(list.Tasks, id)
    d := make(map[string]interface{})
    d["tasks"] = list.Tasks
    // Send update to Firestore
    _, err = doc.Set(r.Ctx, d, firestore.MergeAll)
    if err != nil {
        fmt.Printf("err setting new list data: %v\n", err)
    }
}
