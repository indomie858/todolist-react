//component for individual tasks
import SubTasks from './SubTasks'
import { useState } from 'react';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExpandLessIcon from '@material-ui/icons/ExpandLess';
import CheckBoxOutlineBlankIcon from '@material-ui/icons/CheckBoxOutlineBlank';
import CheckBoxIcon from '@material-ui/icons/CheckBox';
import moment from 'moment'

const Task = (props) => {
    const [showSubTasks, setShowSubTasks] = useState(false);
    const [subTasks, setSubTasks] = useState(props.task.subTasks);
    const [taskComplete, setTaskComplete] = useState(false);

    //function for adding subtasks. currently pushing subtask to array in task object
    //need to handle sending subtask to the backend
    const addSubTask = (subTask) => {
        setSubTasks([...subTasks, subTask]);
    }

    //onClick={() => setShowCompleteButton(!showCompleteButton)}
    return (
        <>
            <div className='task' >
                <div className="task-flex-left">
                    <h3 onClick={() => { props.changeTask(props.id); }}>{props.task.text}{' '}</h3>
                    <p>{moment(props.task.date).format("M/D h:MM A")}</p>
                </div>
                <div className="task-flex-mid" onClick={() => setShowSubTasks(!showSubTasks)}>
                    {subTasks.length > 0 ? (!showSubTasks ? <ExpandMoreIcon /> : <ExpandLessIcon />) : ''}
                </div>
                <div className="task-flex-right" onClick={() => {
                    setTaskComplete(!taskComplete);
                    console.log(props.task)
                    }}>
                    {!taskComplete ? <CheckBoxOutlineBlankIcon /> : <CheckBoxIcon />}
                </div>
            </div>
            {/* displays list of subtasks when individual task is clicked */}
            {showSubTasks && <SubTasks subTasks={subTasks} addSubTask={addSubTask} />}
        </>
    )
}

export default Task