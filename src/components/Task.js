//component for individual tasks
import SubTasks from './SubTasks'
import { useState } from 'react';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExpandLessIcon from '@material-ui/icons/ExpandLess';
import moment from 'moment'

const Task = (props) => {
    const [showSubTasks, setShowSubTasks] = useState(false);
    const [subTasks, setSubTasks] = useState(props.task.subTasks);

    //function for adding subtasks. currently pushing subtask to array in task object
    //need to handle sending subtask to the backend
    const addSubTask = (subTask) => {
        setSubTasks([...subTasks, subTask]);
    }

    return (
        <>
            <div className='task'>
                <div className="task-flex-left">
                    <h3 onClick={() => {props.changeTask(props.id);}}>{props.task.text}{' '}</h3>
                    <p>{moment(props.task.date).format("M/D h:MM A")}</p>
                </div>
                <div className="task-flex-right" onClick={() => setShowSubTasks(!showSubTasks)}>
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