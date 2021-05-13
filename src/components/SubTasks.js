// component for list of subtasks
import SubTask from './SubTask'
import AddSubTask from './AddSubTask'
import AddCircleIcon from '@material-ui/icons/AddCircle';
import { useState } from 'react';

const SubTasks = ({ subTasks }) => {
    const [showSubTaskInput, setShowSubTaskInput] = useState(false);

    return (
        <>
            {/* iterates through subtask array and passes each subtask to SubTask component */}
            {subTasks ? subTasks.map((subTask) => (<SubTask subTask={subTask} />)) : ''}
            <div className="addsubtask">
                {showSubTaskInput ? <AddSubTask /> : <button className="addsubtaskicon" onClick={() => setShowSubTaskInput(true)} ><AddCircleIcon fontSize="large" /></button>}
            </div>
        </>
    )
}

export default SubTasks
