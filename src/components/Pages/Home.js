import React from 'react'
import { useState } from 'react';
import { Redirect } from 'react-router-dom';
import Header from '../Header';
import Tasks from '../Tasks';
import BottomNavBar from '../BottomNavBar';
import AddTask from '../AddTask.js'

const Home = () => {

  const getToken = () => {
    //tokens are stored locally so user doesn't have to keep logging in
    const token = sessionStorage.getItem('token');
    return token
  };

  //state for displaying tasks. currently has placeholder objects. will replace with tasks from database
  const [tasks, setTasks] = useState(
    [
      {
        id: 1,
        text: 'Doctors appt',
        day: 'Feb 5th at 2:30pm',
        reminder: true,
        subTasks: [
          'Get in car',
          'Drive to doctor'
        ],
      },
      {
        id: 2,
        text: 'School meeting',
        day: 'Feb 6th at 1:30pm',
        reminder: true,
        subTasks: [
          'Take two shots',
          'Put on some pants',
          'Get a ride to school'
        ],
      },
      {
        id: 3,
        text: 'Food shopping',
        day: 'Feb 7th at 5:30pm',
        reminder: false,
        subTasks: [
          'Make shopping list',
          'Drive to Costco',
          'Buy some shit'
        ],
      },
      {
        id: 4,
        text: 'something something darkside',
        day: 'Feb 8th at 2:30pm',
        reminder: true,
        subTasks: [
        ],
      },
    ]
  )



  const [showAddTask, setAddTask] = useState(false);
  const [email, setEmail] = useState('');

  if (!getToken()){
    console.log('/home token does not exist');
    return (<Redirect to="/login" />);
  } else {
    console.log('/home token exists');
    if(email === ''){
      setEmail(sessionStorage.getItem('email'));
    }
  }
  return (
    <>
      {/* <Container maxWidth="xs"> */}
      <p>Welcome {email}</p>
      <div className="mainContainer">
        <Header />
        <div className='listContainer'>
          {/* displays placeholder list and title "Today" */}
          {tasks.length > 0 ? (<Tasks tasks={tasks} listTitle='Today' />) : ('No tasks to show')}
        </div>
        <div className='listContainer'>
          {/* displays same placeholder list except with title "Tomorrow" */}
          {tasks.length > 0 ? (<Tasks tasks={tasks} listTitle='Tomorrow' />) : ('No tasks to show')}
        </div>
        {showAddTask && <AddTask onAdd={() => setAddTask(false)} />}
      </div>
      <BottomNavBar onAddTask={() => setAddTask(true)} />
      {/* </Container> */}
    </>
  )
}

export default Home
