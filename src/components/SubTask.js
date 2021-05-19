//component for individual subtask
import CheckBoxOutlineBlankIcon from '@material-ui/icons/CheckBoxOutlineBlank';
import CheckBoxIcon from '@material-ui/icons/CheckBox';
import { useState } from 'react';

const SubTask = ({ subTask }) => {
    const [taskComplete, setTaskComplete] = useState(subTask.isCompleted);

    return (
        // outputs subtask from subtasks array
        <div className="subtask">
            <p>{subTask.text}</p>
            <div className="task-flex-right" onClick={() => setTaskComplete(!taskComplete)} >
                {!taskComplete ? <CheckBoxOutlineBlankIcon /> : <CheckBoxIcon />}
            </div>
        </div>
    )
}

export default SubTask
