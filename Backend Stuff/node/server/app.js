//imports
const createError = require('http-errors');
const express = require('express');
const path = require('path');
const fs = require('fs')
const http = require('http')
// importing the dependencies
const cookieParser = require('cookie-parser');

// const bodyParser = require('body-parser');
const cors = require('cors');
const helmet = require('helmet');
const logger = require('morgan');

const indexRouter = require('./routes/index');
const usersRouter = require('./routes/users');

const app = express(); //create express app for routing api calls

let jsonObject = require('./test.json');
const {
  getDefaultNormalizer
} = require('@testing-library/dom');

// view engine setup for express
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

app.use(logger('dev'));
// using bodyParser to parse JSON bodies into JS objects
app.use(express.json());
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter); //default path
app.use('/users', usersRouter); //unused /users path

// adding Helmet to enhance your API's security
app.use(helmet());

// enabling CORS for all requests
app.use(cors());





// starting the server
app.listen(3003, () => {
  console.log('listening on port 3003');
});
module.exports = app;





//user info api calls
app.get('/api/userData/:id', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //necessary callback to send results back to client

    console.log(result)
    res.send(result)
  })


})

app.get('/api/userData/:id/name', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //necessary callback to send results back to client

    console.log(result)
    res.send(result.user.name)
  })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/userData/:id/email', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //necessary callback to send results back to client

    console.log(result)
    res.send(result.user.email)
  })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/userData/:id/status', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //necessary callback to send results back to client

    console.log(result)
    res.send(result.user.status)
  })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/userData/:id/lists', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //necessary callback to send results back to client

    console.log(result)
    res.send(result.lists)
  })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/userData/:id/list/:listID', (req, res) => {
  console.log(req.params.id)
  readAPI(req.params.id, `/list/${req.params.listID}`, (result) => { //necessary callback to send results back to client
    console.log(result)
    res.send(result)
  })
  // res.send(JSON.stringify(name.Name))

})

//begin CRUD api calls section


//CREATE
app.post('/api/create/:uid', (req, res) => {
  let body = req.body

  switch (body.create) {
    case 'user': {
      delete body.create //remove the update parameter to simplify object
      console.log(body) //there should be nothing left in the body here
      createAPI("user/" + req.params.uid, (result) => {
        console.log(result)
        res.send(result)
      })

      break;
    }
    case 'list': {
      delete body.create //remove the update parameter to simplify object
      const list_name = body.list_name;
      delete body.list_name;
      createAPIJSON(req.params.uid + `/list/${list_name}`, body, (result) => { //add /List/:ID to url
        console.log(result)
        res.send(result)
      })

      break;
    }

    case 'task': {
      delete body.create //remove the update parameter to simplify object
      const task_name = body.task_name;
      const parentId = body.parentId;
      console.log(parentId + " " + task_name)
      delete body.task_name;
      delete body.parentId;
      createAPIJSON(req.params.uid + `/task/${task_name}/parent/${parentId}`, body, (result) => { //add /Task/:ID to url
        console.log(result)
        res.send(result)
      })



      break;
    }

    case 'subtask': {
      delete body.create //remove the update parameter to simplify object
      const subtask_name = body.subtask_name;
      const parentId = body.parentId;
      delete body.subtask_name;
      delete body.parentId;
      createAPIJSON(req.params.uid + `/task/${subtask_name}/parent/${parentId}`, body, (result) => { //add /Task/:ID to url
        console.log(result)
        res.send(result)
      })


      break;
    }


    default:
      console.log('default case hit')
  }


})




//UPDATE
app.post('/api/update/:uid', (req, res) => {

  let body = req.body

  // getName(req.params.id,(name)=>{ //necessary callback to send results back to client
  switch (body.update) {
    case 'userSettings':
      delete body.update //remove the update parameter to simplify object
      console.log(body)
      updateAPIJSON(req.params.uid, body, (result) => {
        console.log(result)
        res.send(result)
      })

      break;
    case 'listSettings': {
      delete body.update //remove the update parameter to simplify object
      const listId = body.listId;
      delete body.listId;
      updateAPIJSON(req.params.uid + `/list/${listId}`, body, (result) => { //add /List/:ID to url
        console.log(result)
        res.send(result)
      })

      break;
    }
    case 'taskSettings': {
      delete body.update //remove the update parameter to simplify object
      const taskId = body.taskId;
      delete body.taskId;
      updateAPIJSON(req.params.uid + `/task/${taskId}`, body, (result) => { //add /Task/:ID to url
        console.log(result)
        res.send(result)
      })

      // http://localhost:10000/update/{uid}/task/{id}?

      break;
    }

    case 'subtasks': {
      delete body.update;
      const taskId = body.taskId;
      delete body.taskId;
      updateAPISubtasks(req.params.uid, `/task/${taskId}`, body.sub_tasks, (result) => { //add /Task/:ID to url
        console.log(result)
        res.send(result)
      })
      break;
    }

    default:
      console.log('default case hit')
  }


})



app.get('/api/update/:uid/list/:list_id', (req, res) => {
  updateAPI(req.params.uid, `/list/${req.params.list_id}`, req.query, (output) => {
    console.log(output)
    res.send(output)
  })

})



//DELETE/DESTROY
app.delete('/api/delete/:uid', (req, res) => {
  let body = req.body

  console.log(body.delete)

  switch (body.delete) {
    case 'user':
      delete body.delete //remove the update parameter to simplify object
      console.log(body)
      destroyAPI(req.params.uid, "", (result) => {
        console.log(result)
        res.send(result)
      })

      break;
    case 'list':
      delete body.delete //remove the update parameter to simplify object
      const listId = body.listId;
      delete body.listId;
      destroyAPI(req.params.uid, `/list/${listId}`, (result) => { //add /List/:ID to url
        console.log(result)
        res.send(result)
      })

      break;
    case 'task':
      delete body.delete //remove the update parameter to simplify object
      const taskId = body.taskId;
      console.log("tdoctor" + taskId)
      delete body.taskId;
      destroyAPI(req.params.uid, `/task/${taskId}`, (result) => { //add /Task/:ID to url
        console.log(result)
        res.send(result)
      })


      break;
    case 'subtask':
      delete body.delete //remove the update parameter to simplify object



      break;
    default:
      console.log('default case hit')
  }

})

//Begin Helper Functions

function createAPI(uid, callback) { //general create function (mostly for users)
  var url = 'http://localhost:10000/create/' + uid

  apiCall(url, (output) => {
    callback(output) //return output to the passed in callback function
  })
}

function createAPIJSON(uid, json, callback) { //create function with payload parameters
  var url = 'http://localhost:10000/create/' + uid + "?"

  for (const query in json) {
    url += query + "=" + json[query] + "&"
  }

  url = url.substr(0, url.length - 1)
  console.log(url)
  apiCall(url, (output) => {
    callback(output) //return output to the passed in callback function
  })
}
//

function readAPI(uid, parameters, callback) { //general read function with parameters baked into function call
  apiCall('http://localhost:10000/read/' + uid + parameters, (output) => {
    callback(output)
  })
}

function updateAPI(uid, parameters, queries, callback) { //update function with payload parameters
  var url = 'http://localhost:10000/update/' + uid + parameters + "?"

  for (const query in queries) {
    url += query + "=" + queries[query] + "&"
  }

  url = url.substr(0, url.length - 1)
  console.log(url)
  apiCall(url, (output) => {
    callback(output)
  })
}

function updateAPISubtasks(uid, parameters, subtasks, callback) { //update function just for subtasks since they are special
  var url = 'http://localhost:10000/update/' + uid + parameters




  console.log(url)
  apiCallWithPayload(url, subtasks, (output) => {
    callback(output)
  })
}

function updateAPIJSON(uid, json, callback) { //update function with payload parameters received in json format
  var url = 'http://localhost:10000/update/' + uid + "?"

  for (const query in json) {
    url += query + "=" + json[query] + "&"
  }

  url = url.substr(0, url.length - 1)

  console.log(url)

  apiCall(url, (output) => {
    callback(output) //return output to the passed in callback function
  })
}


function destroyAPI(uid, parameters, callback) { //delete function
  var url = 'http://localhost:10000/destroy/' + uid + parameters

  console.log(url)

  apiCall(url, (output) => {
    callback(output) //return output to the passed in callback function
  })
}




function apiCall(url, callback) { //regular read function
  http.get(url, (resp) => {
    let data = ''

    // A chunk of data has been received.
    resp.on('data', (chunk) => {
      data += chunk
    })

    // The whole response has been received. Print out the result.
    resp.on('end', () => {
      // console.log(JSON.parse(data).explanation)
      callback(data)
      // return output
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message)
  })
}

function apiCallWithPayload(url, payload, callback) { //function for specifically items that must be put in arrays

  url += `?sub_tasks=[`

  for (const item in payload) {
    url += `${item}, `
  }
  url = url.substr(0, url.length - 2) + "]"

  http.get(url, (resp) => {
    let data = ''

    // A chunk of data has been received.
    resp.on('data', (chunk) => {
      data += chunk
    })

    // The whole response has been received. Print out the result.
    resp.on('end', () => {
      // console.log(JSON.parse(data).explanation)
      callback(data)
      // return output
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message)
  })
}
// function apiCallWithPayload(url, payload, callback) { //attempt at making it run by sending an array json

//   const options = {
//     method: 'POST',
//     headers: {
//       Accept: 'application/json',
//       'Content-Type': 'application/json'
//     },
//     body: JSON.stringify(payload)
//   }

//   http.request(url, options, (resp) => {
//     let data = ''

//     // A chunk of data has been received.
//     resp.on('data', (chunk) => {
//       data += chunk
//     })

//     // The whole response has been received. Print out the result.
//     resp.on('end', () => {
//       // console.log(JSON.parse(data).explanation)
//       callback(data)
//       // return output
//     });

//   }).on("error", (err) => {
//     console.log("Error: " + err.message)
//   })
// }

function getTasks(callback) { //get all tasks
  // var output =1
  http.get('http://localhost:10000/readtaskreminders', (resp) => {
    let data = ''

    // A chunk of data has been received.
    resp.on('data', (chunk) => {
      data += chunk
    })

    // The whole response has been received. Print out the result.
    resp.on('end', () => {
      console.log(JSON.parse(data).explanation)
      callback(JSON.parse(data))
      // return output
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message)
  })
  // return output
}

function getName(uid, callback) { //simple api call for reads
  // var output =1
  http.get('http://localhost:10000/read/' + uid, (resp) => {
    let data = ''

    // A chunk of data has been received.
    resp.on('data', (chunk) => {
      data += chunk
    })

    // The whole response has been received. Print out the result.
    resp.on('end', () => {
      console.log(JSON.parse(data).explanation)
      callback(JSON.parse(data))
      // return output
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message)
  })
  // return output
}

function setCustomError(req, res, message, name, status) { //error message for 404
  const err = createError(status, name)
  res.locals.title = name
  res.locals.message = message
  res.locals.error = req.app.get('env') === 'development' ? err : {};
  res.status(status)

}



// catch 404 and forward to error handler
app.use(function (req, res, next) {
  next(createError(404));
});

// error handler
app.use(function (err, req, res, next) {
  // set locals, only providing error in development
  res.locals.title = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.locals.message = res.statusCode + " " + err.message
  res.render('error');
});

function errorHandler(err, req, res, next) {

}



//test function
app.get('/api/:input', (req, res) => {
  if (req.params.input in jsonObject.users) {
    res.send(jsonObject.users[req.params.input])
  } else {
    setCustomError(req, res, "Invalid API Reference. Please check the path again.", "Invalid API Reference", 404)
    res.render('error')
  }
})

//test function
app.get('/getTasks', (req, res) => {
  console.log("HIT /api/getTasks");
  getTasks((result) => { //necessary callback to send results back to client

    //console.log(result)
    var myresults = result["result"]["Tasks"]
    res.send(myresults)
  })
});