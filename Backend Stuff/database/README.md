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
"lists": [
    {
        "id": "364DgExvwpE4lNC7JV59",
        "list_name": "first_list",
        "list_owner": "gNMA6TlIOCdB52LPSuL5",
        "shared_users": [
            ""
        ],
        "tasks": [
            "MKEUu0LxHZtMOd6KfmsB"
        ]
    }
],
"tasks": [
    [
        {
            "id": "MKEUu0LxHZtMOd6KfmsB",
            "text": "first_task",
            "task_owner": "gNMA6TlIOCdB52LPSuL5",
            "parent_id": "364DgExvwpE4lNC7JV59",
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
            "sub_task": false
        }
    ]
],
"user": {
    "id": "gNMA6TlIOCdB52LPSuL5",
    "name": "testing_user_1",
    "name": "testing_user_1",
    "email": "",
    "status": "",
    "default_list": "",
    "default_reminder": "",
    "lists": [
        "364DgExvwpE4lNC7JV59"
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
"list": {
    "id": "fQksVGJzgTUc6FervXa4",
    "list_name": "test_list_1",
    "list_owner": "gNMA6TlIOCdB52LPSuL5",
    "shared_users": [
        ""
    ],
    "tasks": [
        "UeSCAXmJoGXkiXIy99AB"
    ]
},
"tasks": [
    {
        "id": "UeSCAXmJoGXkiXIy99AB",
        "text": "first_task",
        "task_owner": "gNMA6TlIOCdB52LPSuL5",
        "parent_id": "fQksVGJzgTUc6FervXa4",
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
        "sub_task": false
    }
]
```

URL

`http://localhost:10000/create/gNMA6TlIOCdB52LPSuL5/list/test_list_2?lock=true&shared=false`

Return (new list)

```json
"list": {
    "id": "S02pAtwV1Fsc1o5DHbPn",
    "list_name": "test_list_2",
    "list_owner": "gNMA6TlIOCdB52LPSuL5",
    "lock": true,
    "shared_users": [
        ""
    ],
    "tasks": [
        "SfePwO3UuUyYA1cW49KV"
    ]
},
"tasks": [
    {
        "id": "SfePwO3UuUyYA1cW49KV",
        "text": "first_task",
        "task_owner": "gNMA6TlIOCdB52LPSuL5",
        "parent_id": "S02pAtwV1Fsc1o5DHbPn",
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
        "sub_task": false
    }
]
```

### Add Task
`http://localhost:10000/create/{uid}/task/{name}/parents/{pid}?<params>`

**Example**

URL

`http://localhost:10000//create/gNMA6TlIOCdB52LPSuL5/task/test_task_1/parent/fQksVGJzgTUc6FervXa4?sub_task=false&lock=false&date_due=01/02/2006 3:04:05 PM`


Return (new task)

```json
"task": {
    "id": "mOCohcha1i6sInCJPeEp",
    "text": "test_task_1",
    "task_owner": "gNMA6TlIOCdB52LPSuL5",
    "parent_id": "fQksVGJzgTUc6FervXa4",
    "date": "2006-01-02T15:04:05Z",
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
    "sub_task": false
}
```

### Add Subtask
`http://localhost:10000/create/{uid}/subtask/{name}/parent/{pid}`

**Example**

URL

`http://localhost:10000/create/gNMA6TlIOCdB52LPSuL5/subtask/sub_task_1/parent/mOCohcha1i6sInCJPeEp`

Return (task)

```json
"task": {
    "id": "mOCohcha1i6sInCJPeEp",
    "text": "test_task_1",
    "task_owner": "gNMA6TlIOCdB52LPSuL5",
    "parent_id": "fQksVGJzgTUc6FervXa4",
    "date": "2006-01-02T15:04:05Z",
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
    "sub_task": false,
    "subTasks": [
        "sub_task_1"
    ]
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
`http://localhost:10000/destroy/{uid}/list/{id}`

**Example**

URL

`http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W/list/364DgExvwpE4lNC7JV59`

Return

```json
"result": "list successfully deleted"
```

### Destroy Task
`http://localhost:10000/destroy/{uid}/task/{id}`

**Example**

URL

`http://localhost:10000/destroy/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/MKEUu0LxHZtMOd6KfmsB`

Return

```json
"result": "task successfully deleted"
```

## Reading Data From a Collection
### Read User
`http://localhost:10000/read/{uid}`

**Example**

URL

`http://localhost:10000/read/gNMA6TlIOCdB52LPSuL5`

Return (user)

```json
"lists": [
    {
        "id": "364DgExvwpE4lNC7JV59",
        "list_name": "first_list",
        "list_owner": "gNMA6TlIOCdB52LPSuL5",
        "shared_users": [
            ""
        ],
        "tasks": [
            "MKEUu0LxHZtMOd6KfmsB"
        ]
    },
    {
        "id": "S02pAtwV1Fsc1o5DHbPn",
        "list_name": "test_list_2",
        "list_owner": "gNMA6TlIOCdB52LPSuL5",
        "lock": true,
        "shared_users": [
            ""
        ],
        "tasks": [
            "SfePwO3UuUyYA1cW49KV"
        ]
    },
    {
        "id": "fQksVGJzgTUc6FervXa4",
        "list_name": "test_list_1",
        "list_owner": "gNMA6TlIOCdB52LPSuL5",
        "shared_users": [
            ""
        ],
        "tasks": [
            "UeSCAXmJoGXkiXIy99AB",
            "mOCohcha1i6sInCJPeEp"
        ]
    }
],
"tasks": [
    [
        {
            "id": "MKEUu0LxHZtMOd6KfmsB",
            "text": "first_task",
            "task_owner": "gNMA6TlIOCdB52LPSuL5",
            "parent_id": "364DgExvwpE4lNC7JV59",
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
            "sub_task": false
        }
    ],
    [
        {
            "id": "SfePwO3UuUyYA1cW49KV",
            "text": "first_task",
            "task_owner": "gNMA6TlIOCdB52LPSuL5",
            "parent_id": "S02pAtwV1Fsc1o5DHbPn",
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
            "sub_task": false
        }
    ],
    [
        {
            "id": "UeSCAXmJoGXkiXIy99AB",
            "text": "first_task",
            "task_owner": "gNMA6TlIOCdB52LPSuL5",
            "parent_id": "fQksVGJzgTUc6FervXa4",
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
            "sub_task": false
        },
        {
            "id": "mOCohcha1i6sInCJPeEp",
            "text": "test_task_1",
            "task_owner": "gNMA6TlIOCdB52LPSuL5",
            "parent_id": "fQksVGJzgTUc6FervXa4",
            "date": "2006-01-02T15:04:05Z",
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
            "sub_task": false,
            "subTasks": [
                "sub_task_1"
            ]
        }
    ]
],
"user": {
    "id": "gNMA6TlIOCdB52LPSuL5",
    "name": "testing_user_1",
    "lists": [
        "364DgExvwpE4lNC7JV59",
        "fQksVGJzgTUc6FervXa4",
        "S02pAtwV1Fsc1o5DHbPn"
    ]
}
```

### Read List
`http://localhost:10000/read/{uid}/list/{id}`

**Example**

URL

`http://localhost:10000/read/gNMA6TlIOCdB52LPSuL5/list/fQksVGJzgTUc6FervXa4`

Return (list)

```json
"list": {
    "id": "fQksVGJzgTUc6FervXa4",
    "list_name": "test_list_1",
    "list_owner": "gNMA6TlIOCdB52LPSuL5",
    "shared_users": [
        ""
    ],
    "tasks": [
        "UeSCAXmJoGXkiXIy99AB",
        "mOCohcha1i6sInCJPeEp"
    ]
},
"tasks": [
    {
        "id": "UeSCAXmJoGXkiXIy99AB",
        "text": "first_task",
        "task_owner": "gNMA6TlIOCdB52LPSuL5",
        "parent_id": "fQksVGJzgTUc6FervXa4",
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
        "sub_task": false
    },
    {
        "id": "mOCohcha1i6sInCJPeEp",
        "text": "test_task_1",
        "task_owner": "gNMA6TlIOCdB52LPSuL5",
        "parent_id": "fQksVGJzgTUc6FervXa4",
        "date": "2006-01-02T15:04:05Z",
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
        "sub_task": false,
        "subTasks": [
            "sub_task_1"
        ]
    }
]
```

### Read List*s*
( emphasis on the s )

`http://localhost:10000/read/{uid}/lists`

**Example**

URL

`http://localhost:10000/read/gNMA6TlIOCdB52LPSuL5/lists`

Return (all users lists)
```json
"lists": [
    {
        "id": "364DgExvwpE4lNC7JV59",
        "list_name": "first_list",
        "list_owner": "gNMA6TlIOCdB52LPSuL5",
        "shared_users": [
            ""
        ],
        "tasks": [
            "MKEUu0LxHZtMOd6KfmsB"
        ]
    },
    {
        "id": "S02pAtwV1Fsc1o5DHbPn",
        "list_name": "test_list_2",
        "list_owner": "gNMA6TlIOCdB52LPSuL5",
        "lock": true,
        "shared_users": [
            ""
        ],
        "tasks": [
            "SfePwO3UuUyYA1cW49KV"
        ]
    },
    {
        "id": "fQksVGJzgTUc6FervXa4",
        "list_name": "test_list_1",
        "list_owner": "gNMA6TlIOCdB52LPSuL5",
        "shared_users": [
            ""
        ],
        "tasks": [
            "UeSCAXmJoGXkiXIy99AB",
            "mOCohcha1i6sInCJPeEp"
        ]
    }
],
"tasks": [
    [
        {
            "id": "MKEUu0LxHZtMOd6KfmsB",
            "text": "first_task",
            "task_owner": "gNMA6TlIOCdB52LPSuL5",
            "parent_id": "364DgExvwpE4lNC7JV59",
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
            "sub_task": false
        }
    ],
    [
        {
            "id": "SfePwO3UuUyYA1cW49KV",
            "text": "first_task",
            "task_owner": "gNMA6TlIOCdB52LPSuL5",
            "parent_id": "S02pAtwV1Fsc1o5DHbPn",
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
            "sub_task": false
        }
    ],
    [
        {
            "id": "UeSCAXmJoGXkiXIy99AB",
            "text": "first_task",
            "task_owner": "gNMA6TlIOCdB52LPSuL5",
            "parent_id": "fQksVGJzgTUc6FervXa4",
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
            "sub_task": false
        },
        {
            "id": "mOCohcha1i6sInCJPeEp",
            "text": "test_task_1",
            "task_owner": "gNMA6TlIOCdB52LPSuL5",
            "parent_id": "fQksVGJzgTUc6FervXa4",
            "date": "2006-01-02T15:04:05Z",
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
            "sub_task": false,
            "subTasks": [
                "sub_task_1"
            ]
        }
    ]   
]
```

### Read Task
`http://localhost:10000/read/{uid}/task/{id}`

**Example**

URL

`http://localhost:10000/read/gNMA6TlIOCdB52LPSuL5/task/mOCohcha1i6sInCJPeEp`

Return (task)

```json
"task": {
    "id": "mOCohcha1i6sInCJPeEp",
    "text": "test_task_1",
    "task_owner": "gNMA6TlIOCdB52LPSuL5",
    "parent_id": "fQksVGJzgTUc6FervXa4",
    "date": "2006-01-02T15:04:05Z",
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
    "sub_task": false,
    "subTasks": [
        "sub_task_1"
    ]
}
```

### Read Task*s*
( emphasis on the s )

`http://localhost:10000/read/{uid}/tasks/{pid}`

**Example**

URL

`http://localhost:10000/read/pSdsua0LrFB4IiGIsI3j/tasks/Tju8b4Deg2f5lHvgE4PJ`

Return (all tasks in the list)

```json
"tasks": [
    {
        "id": "D59h0RRGCtlidmgcxBG9",
        "text": "first_task",
        "task_owner": "pSdsua0LrFB4IiGIsI3j",
        "parent_id": "Tju8b4Deg2f5lHvgE4PJ",
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
        "sub_task": false
    },
    {
        "id": "i0GGeohSDPz1YFNRrbzU",
        "text": "test_task_1",
        "task_owner": "pSdsua0LrFB4IiGIsI3j",
        "parent_id": "Tju8b4Deg2f5lHvgE4PJ",
        "date": "2006-01-02T15:04:05Z",
        "isComplete": true,
        "willRepeat": false,
        "repeatFrequency": "never",
        "end_repeat": "0001-01-01T00:00:00Z",
        "remind": false,
        "emailSelected": false,
        "discordSelected": true,
        "reminder": "none",
        "reminder_time": "0001-01-01T00:00:00Z",
        "priority": "none",
        "shared": false,
        "sub_task": false,
        "subTasks": [
            "sub_task_1"
        ]
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
`http://localhost:10000/update/{uid}/list/{id}?<params>`

**Example**

URL

`http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/364DgExvwpE4lNC7JV59?list_name=list1updated&lock=false`

Return (updated list)

```json

```

### Edit Task
`http://localhost:10000/update/{uid}/task/{id}?<params>`

**Example**

URL

`http://localhost:10000/update/SFCCBJMyEA3NyBnCRe4j/task/jHExtupOREWA1CyHcEUX?done=true&discord=true`

Return (updated task)

```json
"task": {
    "id": "jHExtupOREWA1CyHcEUX",
    "text": "test_task_1",
    "task_owner": "SFCCBJMyEA3NyBnCRe4j",
    "parent_id": "uK3YOEJizfAS4WqBrjfa",
    "date": "2006-01-02T15:04:05Z",
    "isComplete": true,
    "willRepeat": false,
    "repeatFrequency": "never",
    "end_repeat": "0001-01-01T00:00:00Z",
    "remind": false,
    "emailSelected": false,
    "discordSelected": true,
    "reminder": "none",
    "reminder_time": "0001-01-01T00:00:00Z",
    "priority": "none",
    "shared": false,
    "sub_task": false,
    "subTasks": [
        "sub_task_1"
    ]
}
```
