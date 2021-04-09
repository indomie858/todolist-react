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
   NEVER = "never"
   DAILY = "daily"
   WEEKLY = "weekly"
   BIWEEKLY = "biweekly"
   MONTHLY = "monthly"
   ANNUALLY = "annually"
)

// Reminder setting constants
const (
   NONE = "none"
   ATOE = "at time of event"
   MBE = "minutes before event"
   DBE = "days before event"
   WBE = "weeks before event"
)

type User struct {
   //UserId string `firestore:"userid,omitempty"`
   Name string `firestore:"name,omitempty"`
   //Email string  `firestore:"email,omitempty"`
   //Status string `firestore:"status,omitempty"`
   Lists []string `firestore:"lists,omitempty"`    // points to type List
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
   Name string

   Lock bool

   DueDate time.Time `firestore:"duedate,omitempty"`
   StartDate time.Time `firestore:"startdate,omitempty"`

   Repeat string
   Reminder string
   ReminderTimeFrame int

   Location string
   Description string
   Url string
}

func main() {
   userid := "a3a1hWUx5geKB8qeR6fbk5LZZGI2"

   ctx := context.Background()

   opt := option.WithCredentialsFile("/Users/sabra/Library/Mobile Documents/com~apple~CloudDocs/School/COMP 482/Project1/todolist-react/Backend Stuff/go/friday-584-firebase-adminsdk-es3qw-60b3f32bf1.json")
   client, err := firestore.NewClient(ctx, "friday-584", opt)
	if err != nil {
		log.Fatalf("Cannot create client: %v", err)    // %v is to format error values
	}
	defer client.Close()

   fmt.Printf("Getting User Data -- \n")
   useridpath := fmt.Sprintf("users/%s", userid)
   user := client.Doc(useridpath)
   usersnap, err := user.Get(ctx)
   if err != nil {
      log.Fatalf("Cannot get user snapshot: %v", err)    // %v is to format error values
   }

   // Get the data in map format
   dataMap := usersnap.Data()
   fmt.Println("Data in map format: ")
   fmt.Println(dataMap)

   // Now let's try getting it into a user struct
   var userStruct User

   // Add user data to our structure
   err = usersnap.DataTo(&userStruct)
   if err != nil {
      log.Fatalf("Cannot put data to struct: %v", err)   // %v is to format error values
   }
   fmt.Printf("\nData pulled from structure:\n")
   fmt.Printf("Name: %s, Lists: [%s %s]\n",userStruct.Name, userStruct.Lists[0], userStruct.Lists[1])


   // Now for the list data
   fmt.Printf("\nGetting List Data -- \n")
   listpath := fmt.Sprintf("lists/%s", userStruct.Lists[0])
   list := client.Doc(listpath)
   listsnap, err := list.Get(ctx)
   if err != nil {
      log.Fatalf("Cannot get list snapshot: %v", err)    // %v is to format error values
   }
   fmt.Println("Data in map format: ")
   listMap := listsnap.Data()

   fmt.Println(listMap)

   var listStruct List
   err = listsnap.DataTo(&listStruct)
   if err != nil {
      log.Fatalf("Cannot put data to struct: %v", err)   // %v is to format error values
   }

   fmt.Printf("\nData pulled from structure:\n")
   fmt.Printf("list_name: %s\n", listStruct.Name)
}
