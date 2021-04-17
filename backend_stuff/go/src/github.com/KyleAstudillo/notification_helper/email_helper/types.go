package email_helper
import "github.com/aws/aws-sdk-go/service/ses"
var (
    // Replace sender@example.com with your "From" address.
    // This address must be verified with Amazon SES.
    //Sender = "kyle.astudillo@yahoo.com"
    Sender String

    // Replace recipient@example.com with a "To" address. If your account
    // is still in the sandbox, this address must be verified.
    //Recipient = "kyle.astudillo@yahoo.com"
    Recipient String

    // Specify a configuration set. To use a configuration
    // set, comment the next line and line 92.
    //ConfigurationSet = "ConfigSet"

    // The subject line for the email.
    //Subject = "Amazon SES Test (AWS SDK for Go)"
    Subject String

    // The HTML body for the email.
    /*HtmlBody =  "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
                "<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
                "<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"*/
    HtmlBody String

    //The email body for recipients with non-HTML email clients.
    //TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
    TextBody String

    // The character encoding for the email.
    CharSet = "UTF-8"
)

var svc *ses.SES
