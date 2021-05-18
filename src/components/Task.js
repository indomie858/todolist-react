//component for individual tasks

import SubTasks from './SubTasks'
import { useState } from 'react';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExpandLessIcon from '@material-ui/icons/ExpandLess';

const Task = ({ task }) => {
    const [showSubTasks, setShowSubTasks] = useState(false);
    const [subTasks, setSubTasks] = useState(task.subTasks);

    //function for adding subtasks. currently pushing subtask to array in task object
    //need to handle sending subtask to the backend
    const addSubTask = (subTask) => {
        setSubTasks([...subTasks, subTask]);
    }

    return (
        <>
            <div className='task' onClick={() => setShowSubTasks(!showSubTasks)}>
                <div className="task-flex-left">
                    <h3>{task.text}{' '}</h3>
                    <p>{task.day}</p>
                </div>
                <div className="task-flex-right">
                    {/* oh lawd this is nasty, but oh whale */}
                    {subTasks.length > 0 ? (!showSubTasks ? <ExpandMoreIcon /> : <ExpandLessIcon />) : ''}
                </div>
            </div>
            {/* displays list of subtasks when individual task is clicked */}
            {showSubTasks && <SubTasks subTasks={subTasks} addSubTask={addSubTask} />}
        </>
    )
}

export default Task