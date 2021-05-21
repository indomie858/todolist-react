import React from 'react'
import { useState, useEffect } from 'react';
import { Redirect } from 'react-router-dom';
import Header from '../Header';
import Tasks from '../Tasks';
import BottomNavBar from '../BottomNavBar';
import AddTask from '../AddTask.js'
import ListNav from '../ListNav.js'
import Options from '../Options.js'
import Container from '@material-ui/core/Container';
import { LensTwoTone } from '@material-ui/icons';
import moment from 'moment'

const Home = () => {

  // Current user id is stored in the session object
  const userId = JSON.parse(sessionStorage.getItem("token")).uid;

  const getToken = () => {
    //tokens are stored locally so user doesn't have to keep logging in
    const token = sessionStorage.getItem('token');
    return token
  };

  //state for displaying tasks. Starts with empty placeholders while the page loads, but is quickly replaced
  const [tasks, setTasks] = useState(
    [
      {
        id: "FAKE1",
        text: " ",
        date: "0001-01-01T00:00:00Z",
        parent_id: "mJmia9sdFy6yfb134ygs",
        list: "updated_isolated_list",
        willRepeat: false,
        repeatFrequency: "Never",
        emailSelected: false,
        discordSelected: false,
        end_repeat: "0001-01-01T00:00:00Z",
        isComplete: false,
        lock: true,
        priority: "none",
        remind: false,
        reminder: "none",
        reminder_time: "0001-01-01T00:00:00Z",
        shared: false,
        subTasks: [],
        sub_task: true,
        task_owner: "a3a1hWUx5geKB8qeR6fbk5LZZGI2"
      },
      {
        date: "0001-01-01T00:00:00Z",
        discordSelected: false,
        emailSelected: false,
        end_repeat: "0001-01-01T00:00:00Z",
        id: "FAKE2",
        isComplete: false,
        parent_id: "mJmia9sdFy6yfb134ygs",
        list: "updated_isolated_list",
        priority: "none",
        remind: false,
        reminder: "none",
        reminder_time: "0001-01-01T00:00:00Z",
        repeatFrequency: "never",
        shared: false,
        sub_task: false,
        subTasks: [],
        task_owner: "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
        text: " ",
        willRepeat: false
      }
    ]
  )


  // A bunch of state-aware variables
  const [userLists, setUserLists] = useState([]);
  const [discordDefault, setDiscordDefault] = useState(false);
  const [emailDefault, setEmailDefault] = useState(false);
  const [defaultList, setDefaultList] = useState("Main");
  const [selectedList, setSelectedList] = useState("Main");

  // The most important function!  Whenever something happens that updates the database, this is called afterward.
  function refreshTasks() {
    fetch(`http://localhost:3003/api/userData/${userId}`).then(
      data => data.text().then(
        value => {
          const userData = JSON.parse(value).result;
          console.log("user:")
          console.log(userData.User)
          setDefaultList(userData.User.default_list)
          setDiscordDefault(userData.User.discord_reminder)
          setEmailDefault(userData.User.email_reminder)
          const listsFromDb = userData.Lists;
          let listNames = []; 
          listsFromDb.forEach(list => {
            listNames.push([list.list_name, list.id])
          })
          setUserLists(listNames);
          let newTasks = [];
          console.log("AllTasks")
          console.log(userData.AllTasks)
          // There are some weird edge cases with variables not existing, so we have to check that the user has tasks first
          if (userData.AllTasks[0]) {
            userData.AllTasks.forEach( someList =>
              someList.forEach(task => {
                let parentList = null;

                // parent lists in the db are stored by id, not name, so we get their name for display purposes by cross-referencing the user's lists
                listsFromDb.forEach(list => {
                  if (list.id == task.parent_id) {
                    parentList = list.list_name
                  }
                })
                console.log(task.date)
                //dates are stored in the db in a different format than the datepickers can use, so they are converted here
                task.date = moment(task.date).toDate();
                task.date = moment(task.date).add(7, 'h').toDate();
                console.log(task.date)
                task.list = parentList;
                console.log("parent list info");
                console.log(task.list + selectedList)
                task.subTasks = [];
                // we check if the user is on the same list for each task before adding it to the display.  This allows switching between separate lists.
                if (task.list == selectedList) {
                  newTasks.push(task);
                }
              })
            );
          }
          console.log(newTasks)
          setTasks(newTasks);
        }
      )
    );
  }

  // useEffect makes this basically happen on load, but not get called whenever React decides the page needs to refresh
  useEffect(() => {
    refreshTasks();
    if (!getToken()) {
      console.log('/home token does not exist');
      return (<Redirect to="/login" />);
    } else {
      console.log('/home token exists');
      if (email === '') {
        setEmail(sessionStorage.getItem('email'));
        
      }
    }
  }, []);

  const [showListNav, setListNav] = useState(false);
  const [showOptions, setOptions] = useState(false);
  const [showAddTask, setAddTask] = useState(false);
  const [showChangeTask, setChangeTask] = useState(false);
  const [changingTask, setChangingTask] = useState(0);
  const [email, setEmail] = useState('');

  // We update the database entry for a task here.  
  function updateTask(taskObject) {
    console.log("updating")
    setChangeTask(false);
    let parentId;
    userLists.forEach(([name, id]) => {
      if (name == taskObject.list) {
        parentId = id;
      }
    })

    taskObject.end_repeat = moment(taskObject.end_repeat).format("MM/DD/YYYY hh:MM:ss A")
    console.log(taskObject.end_repeat)

    console.log(taskObject)
    fetch('http://localhost:3003/api/update/'+userId, {
                    method: 'POST',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        
                        update: 'taskSettings', 
                        taskId: taskObject.id,
                        date: taskObject.date,
                        discordSelected: taskObject.discordSelected,
                        emailSelected: taskObject.emailSelected,
                        end_repeat: taskObject.end_repeat,
                        parent_id: parentId,
                        remind: taskObject.remind,
                        reminder_time: taskObject.date,
                        repeatFrequency: taskObject.repeatFrequency,
                        willRepeat: taskObject.willRepeat,
                        text: taskObject.text,
                        
                    })
            
                }).then(response => {
                    if(response.status===404){
                        return "Error: 404"
                    }else{
                        return response
                      }
                }).then(data=>{ console.log(JSON.stringify(data)); refreshTasks(); });
  }

// very similar to above, but we don't pass in an id because the db will create one
  function createTask(taskObject) {
    console.log("creating")
    setAddTask(false); 
    let parentId;
    userLists.forEach(([name, id]) => {
      console.log(id + name)
      if (name == taskObject.list) {
        parentId = id;
      }
      console.log(parentId)
    })

    taskObject.end_repeat = moment(taskObject.end_repeat).format("MM/DD/YYYY hh:MM:ss A")

    console.log(taskObject)
    fetch('http://localhost:3003/api/create/'+userId, {
                    method: 'POST',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        
                        create: 'task', 
                        date: taskObject.date,
                        discordSelected: taskObject.discordSelected,
                        emailSelected: taskObject.emailSelected,
                        end_repeat: taskObject.end_repeat,
                        parentId: parentId,
                        remind: taskObject.remind,
                        reminder_time: taskObject.date,
                        repeatFrequency: taskObject.repeatFrequency,
                        sub_task: true,
                        willRepeat: taskObject.willRepeat,
                        task_name: taskObject.text,
                        
                    })
            
                }).then(response => {
                    if(response.status===404){
                        return "Error: 404"
                    }else{
                        return response
                      }
                }).then(data=>{ console.log(JSON.stringify(data)); refreshTasks(); });
  }

  // again, similar to above, but for user settings
  function updateUserSettings(userObject) {
    fetch('http://localhost:3003/api/update/'+userId, {
                    method: 'POST',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        
                        update: 'userSettings', 
                        default_list: userObject.default_list,
                        discord_reminder: userObject.discord_reminder,
                        email_reminder: userObject.email_reminder
                        
                    })
            
                }).then(response => {
                    if(response.status===404){
                        return "Error: 404"
                    }else{
                        return response
                      }
                }).then(data=>{ console.log(JSON.stringify(data)); refreshTasks(); });
  }

  // Right now, this is called whenever a task is marked as completed.  In the future, we'd like to make this be separate from completion marking
  function deleteTask(id) {
    console.log("delete this:")
    console.log(id)
    console.log('http://localhost:3003/api/delete/'+userId)
    fetch('http://localhost:3003/api/delete/'+userId, {
                    method: 'DELETE',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        
                        delete: 'task', 
                        taskId: id
                        
                    })
            
                }).then(response => {
                    if(response.status===404){
                        return "Error: 404"
                    }else{
                        return response
                      }
                }).then(data=>{ console.log(JSON.stringify(data)); refreshTasks(); });
  }

  function choseAList(chosen) {
    setListNav(false); 
    setSelectedList(chosen);
    
  }

  // this is called whenever selectedList is updated (by React)
  useEffect(() => {
    refreshTasks();
 }, [selectedList]);
  
  return (
    <>
      <div className="mainContainer">
        {/* for all these overlay windows, we check whether or not to show them with a boolean */}
        {showAddTask && <AddTask userLists={userLists} list={defaultList} onAdd={createTask} defaultReminders={{ "discord": true, "email": false }} onCancel={() => setAddTask(false)} />}
        {/* if a task is being changed, we show the task adder, but with a bunch of extra data from that task */}
        {showChangeTask && <AddTask userLists={userLists} onAdd={updateTask} defaultReminders={{ "discord": true, "email": false }} onCancel={() => setChangeTask(false)}
          id={changingTask.id}
          date={changingTask.date}
          time={changingTask.date}
          text={changingTask.text}
          list={changingTask.list}
          willRepeat={changingTask.willRepeat}
          reminder={changingTask.remind}
          repeatFrequency={changingTask.repeatFrequency}
          emailSelected={changingTask.emailSelected}
          discordSelected={changingTask.discordSelected}
          subtasks={changingTask.subTasks}
        />}
        {showListNav && <ListNav onChooseList={choseAList} lists={[{ name: "Main" }, { name: "Shared" }]} />}
        {showOptions && <Options onChooseOption={updateUserSettings} userLists={userLists} defaultList={defaultList} defaultReminders={{ "discord": discordDefault, "email": emailDefault }} />}
        <Header />
        <div className='listContainer'>
          {/* displays placeholder list and title */}
          {tasks.length > 0 ? (<Tasks tasks={tasks} listTitle={selectedList} markCompleted={deleteTask} changeTask={
            (id) => {
              for (let i = 0; i < tasks.length; i++) {
                if (tasks[i].id === id) {
                  setChangingTask(tasks[i])
                }
              }
              setChangeTask(!showChangeTask);
              setListNav(false);
              setOptions(false);
              setAddTask(false);
          }
          }
           />) : ('No tasks to show')}
        </div>
      </div>

      {/* this could be simplified greatly by abstracting/modularizing these functions, but hey it works. */}
      <BottomNavBar onAddTask={() => { setAddTask(!showAddTask); setChangeTask(false); setListNav(false); setOptions(false) }} onListNav={() => { setListNav(!showListNav); setChangeTask(false); setAddTask(false); setOptions(false) }} onOptions={() => { setListNav(false); setChangeTask(false); setAddTask(false); setOptions(!showOptions) }} />
      {/* </Container> */}
    </>
  )
}

export default Home
