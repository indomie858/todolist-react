package discord_helper

import (
  "os"
  "fmt"
  "bytes"
  "mime/multipart"
  "net/http"
  "io/ioutil"
)

func Init() {
  DH_Webhook := os.Getenv("DH_Webhook")
  CustomInit(DH_Webhook)
}

func CustomInit(webhook string) {
  Webhook = webhook
}

func SendMesage(){
  request(Webhook, "Automated Message")
}

func Request(webhook string, message string){
  request(webhook, message)
}

func request(webhook string, message string) {

  url := "https://discordapp.com/api/webhooks/" + webhook
  method := "POST"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  _ = writer.WriteField("content", message)
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
    return
  }


  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "multipart/form-data")
  req.Header.Add("Cookie", "__cfduid=df78514c97064b63f99d7e9262702b8f71618700505; __dcfduid=1a991acc1d8a48df9c28d747707ad0e7")

  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
