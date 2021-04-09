# GO LANG
Get [golang](https://golang.org/dl/)

# FIRESTORE SDK

In terminal in the project folder run
```bash
   go get firebase.google.com/go
```

[Go here](https://console.firebase.google.com/u/0/project/friday-584/settings/serviceaccounts/adminsdk)

Generate new private key for Go admin account *idk if clicking on Go matters or not tbh*

Then run
```bash
   export GOOGLE_APPLICATION_CREDENTIALS="your-path-to-the-downloaded-service-account.json"
```

Change line 79 in database.go to the path used above

In terminal run
```bash
   go run database.go
```

database.go shows very basic reading of data from firestore 
