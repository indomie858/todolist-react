package request

import (
	"fmt"
	"log"
	"os"
   "context"
   "github.com/joho/godotenv"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)


// func GetCredentials {{{
//
// Reads the credential variables from the env file, formatting them into the
// proper JSON format
// Returns that as an array of bytes to passed to Google's option.WithCredentialsJSON
func GetCredentials() []byte {
	t := os.Getenv("TYPE")
	pid := os.Getenv("PROJECT_ID")
	pkid := os.Getenv("PRIVATE_KEY_ID")
	pk := os.Getenv("PRIVATE_KEY")
	ce := os.Getenv("CLIENT_EMAIL")
	cid := os.Getenv("CLIENT_ID")
	au := os.Getenv("AUTH_URI")
	tu := os.Getenv("TOKEN_URI")
	ap := os.Getenv("AUTH_PROVIDER")
	cert := os.Getenv("CLIENT_X509_CERT_URL")

	jsonCreds := fmt.Sprintf(`{
      "type": "%s",
      "project_id": "%s",
      "private_key_id": "%s",
      "private_key": "%s",
      "client_email": "%s",
      "client_id": "%s",
      "auth_uri": "%s",
      "token_uri": "%s",
      "auth_provider_x509_cert_url": "%s",
      "client_x509_cert_url": "%s"
   }`, t, pid, pkid, pk, ce, cid, au, tu, ap, cert)

	credentials := []byte(jsonCreds)
	return credentials
} // }}}

// func GetClient {{{
//
// Returns a firestore client so we can communicate to the database.
func (r *Request) GetClient() {
   ctx := context.Background()
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load env file: %v", err)
	}

	// Set our credential variables
	pid := os.Getenv("PROJECT_ID")
	credentials := GetCredentials()
	opt := option.WithCredentialsJSON(credentials)

	// Retrieve the firestore client
	client, err := firestore.NewClient(ctx, pid, opt)
	if err != nil {
		log.Fatalf("Cannot create client: %v", err) // %v is to format error values
	}
	r.Client = client
} // }}}
