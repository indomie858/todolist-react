package notification_helper

//export GOPATH=/home/kyle/todolist-react/backend_stuff/go:/usr/lib/go-1.13:/home/kyle/go

import (
  "os"
  _"io/ioutil"
  _"encoding/json"
  "fmt"
  "time"
  _"net/http"
  "github.com/KyleAstudillo/notification_helper/email_helper"
  "github.com/KyleAstudillo/notification_helper/discord_helper"
)

func NotificationHelper() {

  email_helper.Init()
  discord_helper.Init()

  for { //infinite loop
    go poll_notifications()
    time.Sleep(time.Second * 300)
  }
}

func poll_notifications() {
  fmt.Println("Test ", time.Now())
  tasks := get_list()
  fmt.Println(tasks[0].Url)
  for _, task := range tasks {
    if task.Remind {
      if task.Email {
        email_helper.SendEmail()
      }
      if task.Discord {
        //discord_helper.SendMesage()
        discord_helper.Request(task.Url, "Reminder: "+task.Name)
        time.Sleep(time.Second * 5)
      }
    }
  }
}

func get_list_test() ([]Task) {
  var tasks []Task
  var test_task Task
  test_task.Name = "Hello World"
  test_task.Remind = true
  test_task.Email = true
  test_task.Discord = true
  test_task.Url = os.Getenv("DH_Webhook")
  tasks = append(tasks, test_task)
  return tasks
}

func get_list() ([]Task) {
  resp, err := http.Get("http://localhost:10000/readtaskreminders")
  //task := Task{"ID", "Name", "Owner", true, time.Now(), time.Now(), time.Now(), false, "true", true, "0", 0, "", "", ""}
  //jd, err := json.MarshalIndent(myTask, "", "  ")

  if err != nil {
    fmt.Println("get_list: http_get failed")
    fmt.Println(err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()
  if err != nil {
    fmt.Println("get_list: read_body failed")
    fmt.Println(err)
  }
  //fmt.Println(string(body))
  var r []map[string]interface{}
  err = json.Unmarshal(body, &r)
  var tasks []Task
  for i, _ := range r {
    var task Task
    if v, ok := r[i]["id"].(string); ok {
      task.Id = v
    } else {
      task.Id = "nil"
    }
    if v, ok := r[i]["text"].(string); ok {
      task.Name = v
    } else {
      task.Name = "nil"
    }
    //task.Owner = r[i]["task_owner"].(string)
    //task.Parent = r[i]["parent_id"].(string)
    //task.List = r[i][]
    //task.Lock = r[i]["lock"].(bool)
    //task.DateDue = r[i]["date"].(time.Time)
    //task.Done = r[i]["isComplete"].(bool)
    //task.Repeating = r[i]["repeatFrequency"].(bool)
    //task.Repeat = r[i]["willRepeat"].(string)
    //task.EndRepeat = r[i]["end_repeat"].(time.Time)
    if v, ok := r[i]["remind"].(bool); ok {
      task.Remind = v
    } else {
      task.Remind = false
    }
    if v, ok := r[i]["emailSelected"].(bool); ok {
      task.Email = v
    } else {
      task.Email = false
    }
    if v, ok := r[i]["discordSelected"].(bool); ok {
      task.Discord = v
    } else {
      task.Discord = false
    }
    if v, ok := r[i]["reminder"].(string); ok {
      task.Reminder = v
    } else {
      task.Reminder = "nil"
    }
    if v, ok := r[i]["reminder_time"].(time.Time); ok {
      task.RemindTime = v
    } else {
      task.RemindTime = time.Now()
    }
    //task.Priority = r[i]["priority"].(string)
    //task.Location = r[i][]
    //task.Description = r[i][]
    if v, ok := r[i]["url"].(string); ok {
      task.Url = v
    } else {
      task.Url = os.Getenv("DH_Webhook")
    }
    //task.Shared = r[i]["shared"].(bool)
    //task.SharedUsers = r[i][]
    //task.Subtasks = r[i][]
    tasks = append(tasks, task)
  }

  /*
  var tasks []Task
  var test_task Task
  test_task.Name = "Hello World"
  test_task.Remind = true
  test_task.Email = true
  test_task.Discord = true
  test_task.Url = os.Getenv("DH_Webhook")
  tasks = append(tasks, test_task)
  */
  return tasks
}
