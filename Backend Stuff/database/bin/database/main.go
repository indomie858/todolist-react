package main

import (
   "fmt"
   "log"
   "time"
   "database/request"
)

var uid = "a3a1hWUx5geKB8qeR6fbk5LZZGI2"

// Sample call :
// send_to_goserver(uid, "edit", "listname", {description of what the command needs to do})
func main() {
   log.Printf("Starting database request testing\n\n")
   start := time.Now()

   fmt.Println("TEST : testEditList: ")
   if err := testEditList(); err != nil {
      fmt.Println("TEST FAILED\nerr := %v", err)
   } else {
      fmt.Println("TEST PASSED\n")
   }

   elapsed := time.Since(start)
   log.Printf("Database testing took: %s", elapsed)
}

func testEditList() error {
   action := "edit"
   item := "list"
   description := map[string]interface{}{
      "name": "list1",
      "field": "lock",
      "newval": false,
   }
   err := request.DataBaseRequest(uid, action, item, description)

   return err
}
