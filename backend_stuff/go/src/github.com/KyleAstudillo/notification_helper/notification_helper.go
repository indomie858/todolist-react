package notification_helper

import (
  "fmt"
  "time"
  "github.com/KyleAstudillo/notification_helper/email_helper"
  _ "github.com/KyleAstudillo/notification_helper/discord_helper"
)

func NotificationHelper() {
  email_helper.Init()
  for { //infinite loop
    go poll_notifications()
    time.Sleep(time.Second * 60)
  }
}

func poll_notifications() {
  fmt.Println("Test ", time.Now())
  email_helper.SendEmail()
}
