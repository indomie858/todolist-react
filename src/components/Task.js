//component for individual tasks

import SubTasks from './SubTasks'
import { useState } from 'react';

const Task = ({ task }) => {
    const [showSubTasks, setShowSubTasks] = useState(false);

    return (
        <>
            <div className='task' onClick={() => setShowSubTasks(!showSubTasks)}>
                <h3>{task.text}{' '}</h3>
                <p>{task.day}</p>
            </div>
            {/* displays list of subtasks when individual task is clicked */}
            {showSubTasks && <SubTasks subTasks={task.subTasks}/>}
        </>
    )
}

export default Task