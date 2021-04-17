package main

import (
	"context"
	"fmt"
	"time"

	"log"

	"cloud.google.com/go/firestore"
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

type User struct {
	//UserId string `firestore:"userid,omitempty"`
	Name string `firestore:"name,omitempty"`
	//Email string  `firestore:"email,omitempty"`
	//Status string `firestore:"status,omitempty"`
	Lists []string `firestore:"lists,omitempty"` // points to type List
}

type List struct {
	//ListId string `firestore:"listid,omitempty"`
	Name string `firestore:"list_name,omitempty"`
	//Lock bool
	//Shared bool
	//SharedUsers []string
	//Tasks []string
}

type Task struct {
	TaskId string
	Name   string

	Lock bool

	DueDate   time.Time `firestore:"duedate,omitempty"`
	StartDate time.Time `firestore:"startdate,omitempty"`

	Repeat            string
	Reminder          string
	ReminderTimeFrame int

	Location    string
	Description string
	Url         string
}

func main() {
	// Setup our Firestore client ..
	ctx := context.Background()

	opt := option.WithCredentialsFile("friday-584-firebase-adminsdk.json")
	client, err := firestore.NewClient(ctx, "friday-584", opt)
	if err != nil {
		log.Fatalf("Cannot create client: %v", err) // %v is to format error values
	}
	defer client.Close()

	// Let's try getting user data
	fmt.Printf("Getting User Data -- \n")

	userid := "a3a1hWUx5geKB8qeR6fbk5LZZGI2"

	// This formats a string for us, so we can be eventually passed user id?
	useridpath := fmt.Sprintf("users/%s", userid)

	// Pass that to Firestore
	user := client.Doc(useridpath)

	// Get a snapshot of the user data
	usersnap, err := user.Get(ctx)
	if err != nil {
		log.Fatalf("Cannot get user snapshot: %v", err) // %v is to format error values
	}

	// Get the data in map format - default format for Go's Firestore package
	dataMap := usersnap.Data()
	fmt.Println("Data in map format: ")
	fmt.Println(dataMap)

	// Now let's try getting it into a user struct
	var userStruct User

	// Add user data to our structure
	err = usersnap.DataTo(&userStruct)
	if err != nil {
		log.Fatalf("Cannot put data to struct: %v", err) // %v is to format error values
	}

	fmt.Printf("\nData pulled from structure:\n")
	fmt.Printf("Name: %s, Lists: [%s %s]\n", userStruct.Name, userStruct.Lists[0], userStruct.Lists[1])

	// Now for the list data
	fmt.Printf("\nGetting List Data -- \n")

	// Get the string version of the path
	listpath := fmt.Sprintf("lists/%s", userStruct.Lists[0])

	// Pass that to Firestore
	list := client.Doc(listpath)

	// Get a snapshot of the list data
	listsnap, err := list.Get(ctx)
	if err != nil {
		log.Fatalf("Cannot get list snapshot: %v", err) // %v is to format error values
	}

	// Get the data in map format - default format for Go's Firestore package
	fmt.Println("Data in map format: ")
	listMap := listsnap.Data()
	fmt.Println(listMap)

	// Now let's try getting it into a list struct
	var listStruct List
	err = listsnap.DataTo(&listStruct)
	if err != nil {
		log.Fatalf("Cannot put data to struct: %v", err) // %v is to format error values
	}

	fmt.Printf("\nData pulled from structure:\n")
	fmt.Printf("list_name: %s\n", listStruct.Name)
}
