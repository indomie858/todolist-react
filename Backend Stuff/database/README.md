# TOC
- [Getting Started with Database Server](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#getting-started-with-database-server)
- [Datbase API Requests](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#database-api-requests)
   - [Adding Data to a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#adding-data-to-a-collection)
      - [Add User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-user)
      - [Add List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-list)
      - [Add Task w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-task)
      - [Add Subtask w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-subtask)
   - [Removing Data From a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#removing-data-from-a-collection)
      - [Destroy User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-user)
      - [Destroy List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-list)
      - [Destroy Task w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-task)
   - [Reading Data From a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#reading-data-from-a-collection)
      - [Read User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-user)
      - [Read List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-list)
      - [Read Lists w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-lists)
      - [Read Task w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-task)
      - [Read Tasks w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-tasks)
   - [Editing Data](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#editing-values-in-the-database)
      - [Edit User](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-user)
      - [Edit List](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-list)
      - [Edit Task w/ Ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-task)


# Getting Started with Database Server

## The GOPATH environment variable
[snippet from here](https://golang.org/doc/gopath_code#GOPATH)

The GOPATH environment variable specifies the location of your workspace. It defaults to a directory named go inside your home directory, so $HOME/go on Unix, $home/go on Plan 9, and %USERPROFILE%\go (usually C:\Users\YourName\go) on Windows.

_If you would like to work in a different location, you will need to set GOPATH to the path to that directory._ (Another common setup is to set GOPATH=$HOME). Note that GOPATH must not be the same path as your Go installation.

The command `go env GOPATH` prints the effective current GOPATH; it prints the default location if the environment variable is unset.

For convenience, add the workspace's bin subdirectory to your PATH:

`$ export PATH=$PATH:$(go env GOPATH)/bin`

To learn more about the GOPATH environment variable, see 'go help gopath'.

## .env File
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

`http://localhost:10000/`

## Adding Data to a Collection

### Add User
`http://localhost:10000/create/user/{name}`

**Example**

URL

`http://localhost:10000/create/user/testing_user_1`

Return (new user)

```json
"result": {
    "User": {
        "id": "XIatYAGJBbZgJwBVEYJ3",
        "name": "testing_user_1",
        "lists": [
            "wYb9jQqrLQERkDEjdtHI"
        ],
        "discord_reminder": false,
        "email_reminder": false
    },
    "List": null,
    "Lists": [
        {
            "id": "wYb9jQqrLQERkDEjdtHI",
            "list_name": "first_list",
            "list_owner": "XIatYAGJBbZgJwBVEYJ3",
            "shared_users": [
                ""
            ],
            "tasks": [
                "3C8lL89d4cGi3aZ3OokE"
            ]
        }
    ],
    "Task": null,
    "Tasks": null,
    "AllTasks": [
        [
            {
                "id": "3C8lL89d4cGi3aZ3OokE",
                "text": "first_task",
                "task_owner": "XIatYAGJBbZgJwBVEYJ3",
                "parent_id": "wYb9jQqrLQERkDEjdtHI",
                "date": "0001-01-01T00:00:00Z",
                "isComplete": false,
                "willRepeat": false,
                "repeatFrequency": "never",
                "end_repeat": "0001-01-01T00:00:00Z",
                "remind": false,
                "emailSelected": false,
                "discordSelected": false,
                "reminder": "none",
                "reminder_time": "0001-01-01T00:00:00Z",
                "priority": "none",
                "shared": false,
                "subTasks": [
                    ""
                ]
            }
        ]
    ]
}
```

### Add List
`http://localhost:10000/create/{uid}/list/{name}?<params>`

**Example**

URL

`http://localhost:10000/create/8MFkaIrLbLjkxpzGMCwH/list/test_list_1`

Return (new list)

```json
"result": {
    "User": null,
    "List": {
        "id": "GqU6uVezdZlBYtK4MW72",
        "list_name": "test_list_1",
        "list_owner": "XIatYAGJBbZgJwBVEYJ3",
        "shared_users": [
            ""
        ],
        "tasks": [
            "pzs73PE2dsGYHHjrIRKq"
        ]
    },
    "Lists": null,
    "Task": null,
    "Tasks": [
        {
            "id": "pzs73PE2dsGYHHjrIRKq",
            "text": "first_task",
            "task_owner": "XIatYAGJBbZgJwBVEYJ3",
            "parent_id": "GqU6uVezdZlBYtK4MW72",
            "date": "0001-01-01T00:00:00Z",
            "isComplete": false,
            "willRepeat": false,
            "repeatFrequency": "never",
            "end_repeat": "0001-01-01T00:00:00Z",
            "remind": false,
            "emailSelected": false,
            "discordSelected": false,
            "reminder": "none",
            "reminder_time": "0001-01-01T00:00:00Z",
            "priority": "none",
            "shared": false,
            "subTasks": [
                ""
            ]
        }
    ],
    "AllTasks": null
}
```

URL

`http://localhost:10000/create/gNMA6TlIOCdB52LPSuL5/list/test_list_2?lock=true&shared=false`

Return (new list)

```json
"result": {
    "User": null,
    "List": {
        "id": "RpFOrCSLmsPrQ3qcWL5a",
        "list_name": "test_list_2",
        "list_owner": "XIatYAGJBbZgJwBVEYJ3",
        "lock": true,
        "shared_users": [
            ""
        ],
        "tasks": [
            "gvBuQhMgNZF12M9P7n7C"
        ]
    },
    "Lists": null,
    "Task": null,
    "Tasks": [
        {
            "id": "gvBuQhMgNZF12M9P7n7C",
            "text": "first_task",
            "task_owner": "XIatYAGJBbZgJwBVEYJ3",
            "parent_id": "RpFOrCSLmsPrQ3qcWL5a",
            "date": "0001-01-01T00:00:00Z",
            "isComplete": false,
            "willRepeat": false,
            "repeatFrequency": "never",
            "end_repeat": "0001-01-01T00:00:00Z",
            "remind": false,
            "emailSelected": false,
            "discordSelected": false,
            "reminder": "none",
            "reminder_time": "0001-01-01T00:00:00Z",
            "priority": "none",
            "shared": false,
            "subTasks": [
                ""
            ]
        }
