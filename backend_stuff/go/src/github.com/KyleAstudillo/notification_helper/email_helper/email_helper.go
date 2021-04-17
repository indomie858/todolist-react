package email_helper

import (
  "os"
  "fmt"
  "time"
  _ "reflect"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ses"
  "github.com/aws/aws-sdk-go/aws/awserr"
)

func Init() {
  EH_Sender = os.Getenv("EH_Sender")
  EH_Recipient = os.Getenv("EH_Recipient")
  EH_Subject = os.Getenv("EH_Subject")
  EH_HtmlBody =  os.Getenv("EH_HtmlBody")
  EH_TextBody = os.Getenv("EH_TextBody")
  CustomInit(EH_Sender, EH_Recipient, EH_Subject, EH_HtmlBody)
}

func CustomInit(sender String, recipient String, subject String, html_body String, text_body String) {
  EH_Sender = sender
  EH_Recipient = recipient
  EH_Subject = subject
  EH_HtmlBody =  html_body
  EH_TextBody = text_body


  sess, err := session.NewSession(&aws.Config{
      Region:aws.String("us-west-1"),
      //Endpoint:aws.String("us-west-1.amazonaws.com"),
      CredentialsChainVerboseErrors: &[]bool{true}[0]},
  )
  //TODO(kyle): Handle Errors later
  if (true == false) {
    fmt.Println(err)
  }

  // Create an SES session.
  svc = ses.New(sess)
}

func check_notifications() {
  fmt.Println("Test ", time.Now())
}

func SendEmail(){
  fmt.Println("Sent Email")
  send_email(svc)
}

func send_email(svc *ses.SES) {
  // Assemble the email.
  input := &ses.SendEmailInput{
      Destination: &ses.Destination{
          CcAddresses: []*string{
          },
          ToAddresses: []*string{
              aws.String(Recipient),
          },
      },
      Message: &ses.Message{
          Body: &ses.Body{
              Html: &ses.Content{
                  Charset: aws.String(CharSet),
                  Data:    aws.String(HtmlBody),
              },
              Text: &ses.Content{
                  Charset: aws.String(CharSet),
                  Data:    aws.String(TextBody),
              },
          },
          Subject: &ses.Content{
              Charset: aws.String(CharSet),
              Data:    aws.String(Subject),
          },
      },
      Source: aws.String(Sender),
          // Uncomment to use a configuration set
          //ConfigurationSetName: aws.String(ConfigurationSet),
  }

  // Attempt to send the email.
  result, err := svc.SendEmail(input)

  // Display error messages if they occur.
  if err != nil {
      if aerr, ok := err.(awserr.Error); ok {
          switch aerr.Code() {
          case ses.ErrCodeMessageRejected:
              fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
          case ses.ErrCodeMailFromDomainNotVerifiedException:
              fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
          case ses.ErrCodeConfigurationSetDoesNotExistException:
              fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
          default:
              fmt.Println(aerr.Error())
          }
      } else {
          // Print the error, cast err to awserr.Error to get the Code and
          // Message from an error.
          fmt.Println(err.Error())
      }
  }

  fmt.Println("Email Sent to address: " + Recipient)
  fmt.Println(result)
}
