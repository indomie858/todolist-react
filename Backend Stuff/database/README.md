# REQUESTS
send_to_goserver(uid, action, item, {description of what the command needs to do})

## Actions
- ADD
- GET
- EDIT
- DELETE

## Items
- Client
- List
- Task
- User

### EXAMPLE
```javascript
   send_to_goserver("a3a1hWUx5geKB8qeR6fbk5LZZGI2", "edit", "list", {
      "name": "list1",
      "field": "lock",
      "newval": false
   });
```

# TESTING

Go to database folder in command line
```bash
   cd Backend Stuff/database/bin/database
```

Then run the program
```bash
   go run main.go
```

## TESTS / CODE TO WRITE
TestAddUser
TestAddList
TestAddTask

TestGetClientSetting
TestGetUser
TestGetList
TestGetTask

TestEditClientSetting
TestEditUser
~~TestEditList~~
TestEditTask

TestDeleteUser
TestDeleteList
TestDeleteTask

*Note*

You all probably need to create your own .env file in /Backend Stuff/database/bin/database ...

I don't know how to make it so it just auto works like the server r i p

File should contain
```env
# .env file
# Configuration for Firestore SDK

TYPE=service_account
PROJECT_ID=friday-584
PRIVATE_KEY_ID=
PRIVATE_KEY=
CLIENT_EMAIL=
CLIENT_ID=
AUTH_URI=
TOKEN_URI=
AUTH_PROVIDER=
CLIENT_X509_CERT_URL=
```
