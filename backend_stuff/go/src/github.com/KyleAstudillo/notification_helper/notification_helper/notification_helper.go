package notification_helper

//export GOPATH=/home/kyle/todolist-react/backend_stuff/go:/usr/lib/go-1.13:/home/kyle/go

import (
  _"io/ioutil"
  "encoding/json"
  "fmt"
  "time"
  _"net/http"
  _"github.com/KyleAstudillo/notification_helper/email_helper"
  "github.com/KyleAstudillo/notification_helper/discord_helper"
)

func NotificationHelper() {

  //email_helper.Init()
  discord_helper.Init()

  for { //infinite loop
    go poll_notifications()
    time.Sleep(time.Second * 60)
  }
}

func poll_notifications() {
  fmt.Println("Test ", time.Now())
  //get_list()
  //email_helper.SendEmail()
  discord_helper.SendMesage()
}

func get_list() {
  //resp, err := http.Get("http://localhost:3003/api/userData/a3a1hWUx5geKB8qeR6fbk5LZZGI2/lists")
  task := Task{"ID", "Name", "Owner", true, time.Now(), time.Now(), time.Now(), false, "true", true, "0", 0, "", "", ""}

  jd, err := json.MarshalIndent(task, "", "  ")

  if err != nil {
    fmt.Println("get_list: http_get failed")
    fmt.Println(err)
  }
  fmt.Println(string(jd))

  var t Task
  err = json.Unmarshal(jd, &t)
  fmt.Println(t.Owner)

  /*
  body, err := ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()
  if err != nil {
    fmt.Println("get_list: read_body failed")
    fmt.Println(err)
  }*/
  //fmt.Println(string(body))
  //var r map[string]interface{}
  //err = json.Unmarshal(body, &r)
  //fmt.Println(r["list_name"])
  //a := l["id"]
  //fmt.Println(string(a))
	//fmt.Println(string(data))
}
