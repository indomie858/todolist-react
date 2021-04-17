package main

import (
	"context"
	"fmt"
	"time"

	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Status setting constants
const (
	BUSY = "busy"
	FREE = "free"
)

// Repeat setting constants
const (
	NEVER    = "never"
	DAILY    = "daily"
	WEEKLY   = "weekly"
	BIWEEKLY = "biweekly"
	MONTHLY  = "monthly"
	ANNUALLY = "annually"
)

// Reminder setting constants
const (
	NONE = "none"
	ATOE = "at time of event"
	MBE  = "minutes before event"
	DBE  = "days before event"
	WBE  = "weeks before event"
)

// Structure for user data
type User struct {
	Id     string   `firestore:"id,omitempty"`
	Name   string   `firestore:"name,omitempty"`
	Email  string   `firestore:"email,omitempty"`
	Status string   `firestore:"status,omitempty"`
	Lists  []string `firestore:"lists,omitempty"` // array of list ids
}

// Structure for list data
type List struct {
	Id          string   `firestore:"id,omitempty"`
	Name        string   `firestore:"list_name,omitempty"`
	Lock        bool     `firestore:"userid,omitempty"`
	Shared      bool     `firestore:"userid,omitempty"`
	SharedUsers []string `firestore:"userid,omitempty"`
	Tasks       []string `firestore:"userid,omitempty"` // array of task ids
}

// Structure for task data
type Task struct {
	Id   string `firestore:"id,omitempty"`
	Name string `firestore:"task_name,omitempty"`
	Lock bool   `firestore:"lock,omitempty"`

	DueDate   time.Time `firestore:"duedate,omitempty"`
	StartDate time.Time `firestore:"startdate,omitempty"`

	Repeat    string `firestore:"repeat,omitempty"`
	Reminder  string `firestore:"reminder,omitempty"`
	TimeFrame int    `firestore:"time_frame,omitempty"`

	Location    string `firestore:"location,omitempty"`
	Description string `firestore:"description,omitempty"`
	Url         string `firestore:"url,omitempty"`
}

// func GetUser {{{
//
// Returns a user from the Firestore database
func GetUser(userid string, client *firestore.Client, ctx context.Context) *User {
	var user User

	// Get the Firestore path for the user
	useridpath := fmt.Sprintf("users/%s", userid)

	// Pass that to Firestore
	doc := client.Doc(useridpath)

	// Get a snapshot of the user data
	docsnap, err := doc.Get(ctx)
	if err != nil {
		log.Fatalf("Cannot get user snapshot: %v", err) // %v is to format error values
	}

	// Add the data to our structure
	err = docsnap.DataTo(&user)
	if err != nil {
		log.Fatalf("Cannot put data to struct: %v", err) // %v is to format error values
	}

	// Get & set the user ID
	id := docsnap.Ref.ID
	user.Id = id

	return &user
} // }}}

// func GetList {{{
//
// Returns a list using the list id
func (user *User) GetList(listid string, client *firestore.Client, ctx context.Context) *List {
	var list List

	// Get the string version of the path
	listpath := fmt.Sprintf("lists/%s", listid)

	// Get the list document
	doc := client.Doc(listpath)

	// Get a snapshot of the list data
	docsnap, err := doc.Get(ctx)
	if err != nil {
		log.Fatalf("Cannot get list snapshot: %v", err) // %v is to format error values
	}

	// Put that data into our list structure
	err = docsnap.DataTo(&list)
	if err != nil {
		log.Fatalf("Cannot put data to struct: %v", err) // %v is to format error values
	}

	// Get & set the list ID
	id := docsnap.Ref.ID
	list.Id = id

	return &list
} // }}}

// func GetList {{{
//
// Returns a list using the list name
// Ensures we get the correct list by specifying the list owner
func (user *User) GetListByName(listname string, client *firestore.Client, ctx context.Context) *List {
	var list List
	iter := client.Collection("lists").Where("list_name", "==", listname).Where("list_owner", "==", user.Id).Documents(ctx)
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

// func GetClient {{{
//
// Returns a firestore client so we can communicate to the database.
func GetClient(ctx context.Context) *firestore.Client {
	opt := option.WithCredentialsFile("friday-584-firebase-adminsdk.json")
	client, err := firestore.NewClient(ctx, "friday-584", opt)
	if err != nil {
		log.Fatalf("Cannot create client: %v", err) // %v is to format error values
	}

	return client
} // }}}

// func main {{{
func main() {
	ctx := context.Background()

	// Setup our Firestore client ..
	client := GetClient(ctx)
	defer client.Close()

	// Let's try getting user data
	fmt.Printf("Getting User Data -- \n")

	userid := "a3a1hWUx5geKB8qeR6fbk5LZZGI2"
	user := GetUser(userid, client, ctx)

	fmt.Printf("\nData pulled from structure:\n")
	fmt.Printf("Name: %s, Lists: [%s %s]\n", user.Name, user.Lists[0], user.Lists[1])

	// Now for the list data, using list ID
	fmt.Printf("\nGetting List Data by List ID -- \n")

	listid := user.Lists[0]
	list := user.GetList(listid, client, ctx)
	fmt.Printf("\nData pulled from structure:\n")
	fmt.Printf("list_name: %s list_id: %s\n", list.Name, list.Id)

	// Now for the list data, using the lists name
	fmt.Printf("\nGetting List Data by List Name -- \n")
	listname := "list1"
	list = user.GetListByName(listname, client, ctx)
	fmt.Printf("\nData pulled from structure:\n")
	fmt.Printf("list_name: %s list_id: %s\n", list.Name, list.Id)
} // }}}
