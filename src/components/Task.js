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
    const [taskComplete, setTaskComplete] = useState(props.task.isComplete);
    const [subtaskFlag, setSubtaskFlag] = useState(props.task.sub_task);

    //function for adding subtasks.
    //need to handle sending subtask to the backend
    const addSubTask = (subTask) => {
        setSubTasks([...subTasks, subTask]);
        console.log(subTask);
    }

    //onClick={() => setShowCompleteButton(!showCompleteButton)}
    return (
        <>
            <div className='task' >
                <div className="task-flex-left">
                    <h3 onClick={() => { props.changeTask(props.id); }}>{props.task.text}{' '}</h3>
                    <p>{
                        moment(props.task.date).format("M/D h:mm A")
                    }</p>
                </div>
                <div className="task-flex-mid" onClick={() => setShowSubTasks(!showSubTasks)}>
                    {subTasks ? (!showSubTasks ? <ExpandMoreIcon /> : <ExpandLessIcon />) : ''}
                </div>
                <div className="task-flex-right" onClick={() => {
                    setTaskComplete(!taskComplete);
                    console.log(props.task)
                }}>
                    {!taskComplete ? <CheckBoxOutlineBlankIcon onClick={() => props.markCompleted(props.id)} /> : <CheckBoxIcon />}
                </div>
            </div>
            {/* displays list of subtasks when individual task is clicked */}
            {showSubTasks && <SubTasks subTasks={subTasks} addSubTask={addSubTask} />}
        </>
    )
}

export default Task