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

*Example*

URL

`http://localhost:10000/create/user/sabra`

Return

```
   New user's name: sabra

   PAYLOAD PARAMATERS

   New user ID: eQq07UmFE0PAhks3jwTt
```

**Add List**: http://localhost:10000/create/{uid}/list/{name}?<params>

*Example*

URL

`http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/testaddlist`

`http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/testaddlist?lock=false`

Return

```
   list_name: testaddlist

   PAYLOAD PARAMATERS

   lock => [false]

   New list ID: ELW6il13VZxWthDKy7lu
```

**Add Task**: http://localhost:10000/create/{uid}/task/{name}

**Add Subtask**: http://localhost:10000/create/{uid}/subtask/{name}

## Removing data from a Collection
**Destroy User**: http://localhost:10000/destroy/{uid}

**Destory List**: http://localhost:10000/destroy/list/{lists}

**Destroy Task**: http://localhost:10000/destroy/task/{tasks}

## Reading data from a Collection
**Read User**: http://localhost:10000/read/{uid}

**Read Task**: http://localhost:10000/read/{uid}/list/{name}

**Read List**: http://localhost:10000/read/{uid}/task/{name}

## Editing values in the database
**Edit User**: http://localhost:10000/update/{uid}?<params>

**Edit List**: http://localhost:10000/update/{uid}/list/{list}?<params>

*Example*

URL

`http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/list1?list_name=list1updated&lock=false`

Return

```
   listname: list1
   PAYLOAD PARAMATERS
   list_name => [list1updated]
   lock => [false]
```

> REMINDER: once you go to that URL once, it won't work again unless u edit the list name :)

**Edit Task**: http://localhost:10000/update/{uid}/task/{task}?<params>
