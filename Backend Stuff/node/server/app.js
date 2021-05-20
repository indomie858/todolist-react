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

const app = express();

let jsonObject = require('./test.json');
const {
  getDefaultNormalizer
} = require('@testing-library/dom');

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

app.use(logger('dev'));
// using bodyParser to parse JSON bodies into JS objects
app.use(express.json());
// app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/users', usersRouter);

// adding Helmet to enhance your API's security
app.use(helmet());

// enabling CORS for all requests
app.use(cors());

// var server = http.createServer(app);
// server.listen(port);
// server.on('error', onError);
// server.on('listening', onListening);


// to test login page from front end. will delete later - gaven
app.use('/login', (req, res) => {
  res.send({
    token: 'test123'
  });
});


// defining an endpoint to return all ads
app.get('/test1', (req, res) => {
  res.send("you rang?");
  res.end("hello?")
});


app.get('/api/:input', (req, res) => {
  if (req.params.input in jsonObject.users) {
    res.send(jsonObject.users[req.params.input])
  } else {
    setCustomError(req, res, "Invalid API Reference. Please check the path again.", "Invalid API Reference", 404)
    res.render('error')
  }
})


app.get('/api/userData/:id', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //this is terrible and I hate it

    console.log(result)
    res.send(result)
  })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/userData/:id/name', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //this is terrible and I hate it

    console.log(result)
    res.send(result.user.name)
  })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/userData/:id/email', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //this is terrible and I hate it

    console.log(result)
    res.send(result.user.email)
  })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/userData/:id/status', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //this is terrible and I hate it

    console.log(result)
    res.send(result.user.status)
  })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/userData/:id/lists', (req, res) => {
  console.log(req.params.id)
  getName(req.params.id, (result) => { //this is terrible and I hate it

    console.log(result)
    res.send(result.lists)
  })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/userData/:id/list/:listID', (req, res) => {
  console.log(req.params.id)
  readAPI(req.params.id, `/list/${req.params.listID}`, (result) => { //this is terrible and I hate it
    console.log(result)
    res.send(result)
  })
  // res.send(JSON.stringify(name.Name))

})



app.post('/api/update/:uid', (req, res) => {

  let body = req.body

  // getName(req.params.id,(name)=>{ //this is terrible and I hate it
  switch (body.update) {
    case 'userSettings':
      delete body.update //remove the update parameter to simplify object
      console.log(body)
      updateAPIJSON(req.params.uid, body, (result) => {
        console.log(result)
        res.send(result)
      })

      break;
    case 'listSettings':
      delete body.update //remove the update parameter to simplify object
      const listId = body.listId;
      delete body.listId;
      updateAPIJSON(req.params.uid + `/list/${body.listId}`, body, (result) => { //add /List/:ID to url
        console.log(result)
        res.send(result)
      })

      break;
    case 'taskSettings':
      delete body.update //remove the update parameter to simplify object
      const taskId = body.taskId;
      delete body.taskId;
      updateAPIJSON(req.params.uid + `/task/${taskId}`, body, (result) => { //add /Task/:ID to url
        console.log(result)
        res.send(result)
      })

      // http://localhost:10000/update/{uid}/task/{id}?

      break;
    default:
      console.log('default case hit')
  }
  // console.log(name)
  // res.send(JSON.stringify(name))
  // })
  // res.send(JSON.stringify(name.Name))

})

app.get('/api/update/:uid/list/:list_id', (req, res) => {
  updateAPI(req.params.uid, `/list/${req.params.list_id}`, req.query, (output) => {
    console.log(output)
    res.send(output)
  })

})



function readAPI(uid, parameters, callback) {
  apiCall('http://localhost:10000/read/' + uid + parameters, (output) => {
    callback(output)
  })
}

function updateAPI(uid, parameters, queries, callback) {
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

function updateAPIJSON(uid, json, callback) {
  var url = 'http://localhost:10000/update/' + uid + "?"

  for (const query in json) {
    url += query + "=" + json[query] + "&"
  }

  url = url.substr(0, url.length - 1)
  apiCall(url, (output) => {
    callback(output) //return output to the passed in callback function
  })
}

function apiCall(url, callback) {
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


function getName(uid, callback) {
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

function setCustomError(req, res, message, name, status) {
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

// starting the server
app.listen(3003, () => {
  console.log('listening on port 3003');
});
module.exports = app;