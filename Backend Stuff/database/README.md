# Getting Started with Database Server

I believe you need to make a .env file in `database/bin` ... I don't know how to make it so it just auto works like the server *r i p*

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

In terminal navigate to `database/bin` and run `go run api_server.go`

# Database API Requests

http://localhost:10000/

## Adding data to a Collection
**Add User**: http://localhost:10000/create/user/{name}

**Add List**: http://localhost:10000/create/list/{uid}/{name}

**Add Task**: http://localhost:10000/create/task/{uid}/{name}

**Add Subtask**: http://localhost:10000/create/subtask/{uid}/{name}

## Removing data from a Collection
**Destroy User**: http://localhost:10000/destroy/{uid}

**Destory List**: http://localhost:10000/destroy/list/{lists}

**Destroy Task**: http://localhost:10000/destroy/task/{tasks}

## Reading data from a Collection
**Read User**: http://localhost:10000/read/{uid}

**Read Task**: http://localhost:10000/read/list/{uid}/{name}

**Read List**: http://localhost:10000/read/task/{uid}/{name}

## Editing values in the database
**Edit User**: http://localhost:10000/update/{uid}?<params>

**Edit List**: http://localhost:10000/update/list/{uid}/{list}?<params>

*Example*

`http://localhost:10000/update/list/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list1?list_name=list1updated&lock=false`

> REMINDER: once you go to that URL once, it won't work again unless u edit the list name :)

**Edit Task**: http://localhost:10000/update/task/{uid}/{task}?<params>
