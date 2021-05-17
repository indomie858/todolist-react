import React from 'react'
import { useState } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Header from '../Header';
import Tasks from '../Tasks';
import BottomNavBar from '../BottomNavBar';
import AddTask from '../AddTask.js'
import ListNav from '../ListNav.js'
import Container from '@material-ui/core/Container';

const Home = () => {
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
  const [showListNav, setListNav] = useState(false);
  return (
    <>
      {/* <Container maxWidth="xs"> */}
      <div className="mainContainer">
      {showAddTask && <AddTask onAdd={() => setAddTask(false)}/>}
      {showListNav && <ListNav onChooseList={() => setListNav(false)} lists={['Main', 'Shared', 'Other']}/>}
        <Header />
        <div className='listContainer'>
          {/* displays placeholder list and title "Today" */}
          {tasks.length > 0 ? (<Tasks tasks={tasks} listTitle='Today' />) : ('No tasks to show')}
        </div>
        <div className='listContainer'>
          {/* displays same placeholder list except with title "Tomorrow" */}
          {tasks.length > 0 ? (<Tasks tasks={tasks} listTitle='Tomorrow' />) : ('No tasks to show')}
        </div>
      </div>
      <BottomNavBar onAddTask={() => {setAddTask(!showAddTask); setListNav(false)}} onListNav={() => {setListNav(!showListNav); setAddTask(false)}} />
      {/* </Container> */}
    </>
  )
}

export default Home
