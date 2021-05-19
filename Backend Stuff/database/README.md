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
      - [Edit Task](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-task)


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
"lists": [
    {
        "id": "Cz9C4EtgrVU48A1Y7jyF",
        "list_name": "test_list_2",
        "list_owner": "07mp0ArPHpyaUsguqdgQ",
        "lock": true,
        "shared_users": [
            ""
        ],
        "tasks": [
            "LkQSwXomSGibYslCkQOT"
        ]
    },
    {
        "id": "LPDDhsz0TrybCXCgpCFA",
        "list_name": "test_list_1",
        "list_owner": "07mp0ArPHpyaUsguqdgQ",
        "shared_users": [
            ""
        ],
        "tasks": [
            "FLzwaUZFIgBQOkctQmO8",
            "DIwCu71w4KwgQzvDGbjl"
        ]
    },
    {
        "id": "iQGkZETGh9UaunfHiRHy",
        "list_name": "first_list",
        "list_owner": "07mp0ArPHpyaUsguqdgQ",
        "shared_users": [
            ""
        ],
        "tasks": [
            "OHH8IDDrToHrA43Fe4ip"
        ]
    }
 ],
 "tasks": [
     [{
         "id": "LkQSwXomSGibYslCkQOT",
         "task_name": "first_task",
         "task_owner": "07mp0ArPHpyaUsguqdgQ",
         "parent_id": "Cz9C4EtgrVU48A1Y7jyF",
         "date_due": "0001-01-01T00:00:00Z",
         "done": false,
         "repeating": false,
         "repeat": "never",
         "end_repeat": "0001-01-01T00:00:00Z",
         "remind": false,
         "reminder": "none",
         "reminder_time": "0001-01-01T00:00:00Z",
         "priority": "none",
         "shared": false,
         "sub_task": false
     }],
     [{
         "id": "DIwCu71w4KwgQzvDGbjl",
         "task_name": "test_task_1",
         "task_owner": "07mp0ArPHpyaUsguqdgQ",
         "parent_id": "LPDDhsz0TrybCXCgpCFA",
         "date_due": "2006-01-02T15:04:05Z",
         "done": false,
         "repeating": false,
         "repeat": "never",
         "end_repeat": "0001-01-01T00:00:00Z",
         "remind": false,
         "reminder": "none",
         "reminder_time": "0001-01-01T00:00:00Z",
         "priority": "none",
         "shared": false,
         "sub_task": false,
         "sub_tasks": [
             "KFDfex3MmfDGjIdkWZ5y"
         ]
     },
     {
         "id": "FLzwaUZFIgBQOkctQmO8",
         "task_name": "first_task",
         "task_owner": "07mp0ArPHpyaUsguqdgQ",
         "parent_id": "LPDDhsz0TrybCXCgpCFA",
         "date_due": "0001-01-01T00:00:00Z",
         "done": false,
         "repeating": false,
         "repeat": "never",
         "end_repeat": "0001-01-01T00:00:00Z",
         "remind": false,
         "reminder": "none",
         "reminder_time": "0001-01-01T00:00:00Z",
         "priority": "none",
         "shared": false,
         "sub_task": false
     }],
     [{
         "id": "OHH8IDDrToHrA43Fe4ip",
         "task_name": "first_task",
         "task_owner": "07mp0ArPHpyaUsguqdgQ",
         "parent_id": "iQGkZETGh9UaunfHiRHy",
         "date_due": "0001-01-01T00:00:00Z",
         "done": false,
         "repeating": false,
         "repeat": "never",
         "end_repeat": "0001-01-01T00:00:00Z",
         "remind": false,
         "reminder": "none",
         "reminder_time": "0001-01-01T00:00:00Z",
         "priority": "none",
         "shared": false,
         "sub_task": false
     }]
   ],
   "user": {
       "id": "07mp0ArPHpyaUsguqdgQ",
       "name": "testing_user_1",
       "lists": [
           "iQGkZETGh9UaunfHiRHy",
           "LPDDhsz0TrybCXCgpCFA",
           "Cz9C4EtgrVU48A1Y7jyF"
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
    "id": "EneLiFY9NSqUZXMrz6n4",
    "list_name": "test_list_1",
    "list_owner": "8MFkaIrLbLjkxpzGMCwH",
    "shared_users": [
        ""
    ],
    "tasks": [
        "hsyJEeUcLxfNEMReHig8"
    ]
}
```

URL

`http://localhost:10000/create/8MFkaIrLbLjkxpzGMCwH/list/test_list_2?lock=true&shared=false`

Return (new list)

```json
"result": {
    "id": "pT3MWftv7KMYRYdOAsGl",
    "list_name": "test_list_2",
    "list_owner": "8MFkaIrLbLjkxpzGMCwH",
    "lock": true,
    "shared_users": [
        ""
    ],
    "tasks": [
        "XPPbqt4j9DlUziKG0opK"
    ]
}
```

### Add Task
`http://localhost:10000/create/{uid}/task/{name}/parents/{pid}?<params>`

**Example**

URL

`http://localhost:10000/create/8MFkaIrLbLjkxpzGMCwH/task/test_task_1/parent/EneLiFY9NSqUZXMrz6n4?sub_task=false&lock=false&date_due=01/02/2006 3:04:05 PM`


Return (new task)

```json
"result": {
    "id": "WALBFFAGJBbDYycqF1g7",
    "task_name": "test_task_1",
    "task_owner": "8MFkaIrLbLjkxpzGMCwH",
    "parent_id": "EneLiFY9NSqUZXMrz6n4",
    "date_due": "2006-01-02T15:04:05Z",
    "repeat": "never",
    "end_repeat": "0001-01-01T00:00:00Z",
    "reminder": "none",
    "reminder_time": "0001-01-01T00:00:00Z",
    "priority": "none"
}
```

### Add Subtask
`http://localhost:10000/create/{uid}/subtask/{name}/parent/{pid}`

**Example**

URL

`http://localhost:10000/create/8MFkaIrLbLjkxpzGMCwH/subtask/sub_task_1/parent/WALBFFAGJBbDYycqF1g7`

Return (new subtask)

```json
"result": {
    "id": "8PNXOJBOBBAvcCexKt8Q",
    "task_name": "sub_task_1",
    "task_owner": "8MFkaIrLbLjkxpzGMCwH",
    "parent_id": "WALBFFAGJBbDYycqF1g7",
    "date_due": "0001-01-01T00:00:00Z",
    "repeat": "never",
    "end_repeat": "0001-01-01T00:00:00Z",
    "reminder": "none",
    "reminder_time": "0001-01-01T00:00:00Z",
    "priority": "none",
    "sub_task": true    
}
```

## Removing Data From a Collection
### Destroy User
`http://localhost:10000/destroy/{uid}`

**Example**

URL

`http://localhost:10000/destroy/8MFkaIrLbLjkxpzGMCwH`

Return

```json
"result": "user successfully deleted"
```

### Destroy List
`http://localhost:10000/destroy/{uid}/list/{name}`

**Example**

URL

`http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W/list/first_list`

Return

```json
"result": "list successfully deleted"
```

### Destroy Task
`http://localhost:10000/destroy/{uid}/task/{name}`

**Example**

URL

`http://localhost:10000/destroy/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1`

Return

```json
"result": "task successfully deleted"
```

## Reading Data From a Collection
### Read User
`http://localhost:10000/read/{uid}`

**Example**

URL

`http://localhost:10000/read/8MFkaIrLbLjkxpzGMCwH`

Return (user)

```json
"result": {
    "id": "8MFkaIrLbLjkxpzGMCwH",
    "name": "testing_user_1",
    "lists": [
        "tcJQcK8bnLCfzE12p6BJ",
        "EneLiFY9NSqUZXMrz6n4",
        "pT3MWftv7KMYRYdOAsGl"
    ]
 }
```

### Read List
`http://localhost:10000/read/{uid}/list/{name}`

**Example**

URL

`http://localhost:10000/read/8MFkaIrLbLjkxpzGMCwH/list/test_list_1`

Return (list)

```json
"result": {
    "id": "EneLiFY9NSqUZXMrz6n4",
    "list_name": "test_list_1",
    "list_owner": "8MFkaIrLbLjkxpzGMCwH",
    "shared_users": [
        ""
    ],
    "tasks": [
        "hsyJEeUcLxfNEMReHig8",
        "WALBFFAGJBbDYycqF1g7"
    ]
}
```

### Read List*s*
( emphasis on the s )

`http://localhost:10000/read/{uid}/lists`

**Example**

URL

`http://localhost:10000/read/8MFkaIrLbLjkxpzGMCwH/lists`

Return (all users lists)
```json
"result": [
    {
        "id": "EneLiFY9NSqUZXMrz6n4",
        "list_name": "test_list_1",
        "list_owner": "8MFkaIrLbLjkxpzGMCwH",
        "shared_users": [
            ""
        ],
        "tasks": [
            "hsyJEeUcLxfNEMReHig8",
            "WALBFFAGJBbDYycqF1g7"
        ]
    },
    {
        "id": "pT3MWftv7KMYRYdOAsGl",
        "list_name": "test_list_2",
        "list_owner": "8MFkaIrLbLjkxpzGMCwH",
        "lock": true,
        "shared_users": [
            ""
        ],
        "tasks": [
            "XPPbqt4j9DlUziKG0opK"
        ]
    },
    {
        "id": "tcJQcK8bnLCfzE12p6BJ",
        "list_name": "first_list",
        "list_owner": "8MFkaIrLbLjkxpzGMCwH",
        "shared_users": [
            ""
        ],
        "tasks": [
            "kCqFAdMNeSmi9zRcTmWb"
        ]
    }
]
```

### Read Task
`http://localhost:10000/read/{uid}/task/{name}/parent/{pid}`

**Example**

URL

`http://localhost:10000//read/8MFkaIrLbLjkxpzGMCwH/task/test_task_1/parent/EneLiFY9NSqUZXMrz6n4s`

Return (task)

```json
"result": {
    "id": "WALBFFAGJBbDYycqF1g7",
    "task_name": "test_task_1",
    "task_owner": "8MFkaIrLbLjkxpzGMCwH",
    "parent_id": "EneLiFY9NSqUZXMrz6n4",
    "date_due": "2006-01-02T15:04:05Z",
    "repeat": "never",
    "end_repeat": "0001-01-01T00:00:00Z",
    "reminder": "none",
    "reminder_time": "0001-01-01T00:00:00Z",
    "priority": "none",
    "sub_tasks": [
        "8PNXOJBOBBAvcCexKt8Q"
    ]
}

```

### Read Task*s*
( emphasis on the s )

`http://localhost:10000/read/{uid}/tasks/{parent_id}`

**Example**

URL

`http://localhost:10000/read/8MFkaIrLbLjkxpzGMCwH/tasks/EneLiFY9NSqUZXMrz6n4`

Return (all tasks in the list)

```json
"result": [
    {
        "id": "WALBFFAGJBbDYycqF1g7",
        "task_name": "test_task_1",
        "task_owner": "8MFkaIrLbLjkxpzGMCwH",
        "parent_id": "EneLiFY9NSqUZXMrz6n4",
        "date_due": "2006-01-02T15:04:05Z",
        "repeat": "never",
        "end_repeat": "0001-01-01T00:00:00Z",
        "reminder": "none",
        "reminder_time": "0001-01-01T00:00:00Z",
        "priority": "none",
        "sub_tasks": [
            "8PNXOJBOBBAvcCexKt8Q"
        ]
    },
    {
        "id": "hsyJEeUcLxfNEMReHig8",
        "task_name": "first_task",
        "task_owner": "8MFkaIrLbLjkxpzGMCwH",
        "parent_id": "EneLiFY9NSqUZXMrz6n4",
        "date_due": "0001-01-01T00:00:00Z",
        "repeat": "never",
        "end_repeat": "0001-01-01T00:00:00Z",
        "reminder": "none",
        "reminder_time": "0001-01-01T00:00:00Z",
        "priority": "none"
    }
]
```

## Editing values in the database
### Edit User
`http://localhost:10000/update/{uid}?<params>`

**Example**

Add list to list array

URL

`http://localhost:10000/update/MIUVfleqSkxAtzwNeW0W?lists=qqEkD06oFudIRrCVPAc5`

Return (updated user)

```json

```

### Edit List
`http://localhost:10000/update/{uid}/list/{list}?<params>`

**Example**

URL

`http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/list1?list_name=list1updated&lock=false`

Return (updated list)

```json

```


> REMINDER: once you go to that URL once, it won't work again unless u edit the list name :)

URL

`http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/list1updated?list_name=list1`

Return (updated list)

```json

```

### Edit Task
`http://localhost:10000/update/{uid}/task/{task}?<params>`
