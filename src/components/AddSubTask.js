// component with form for adding subtasks

import Button from '@material-ui/core/Button';
import AddCircleIcon from '@material-ui/icons/AddCircle';
import TextField from '@material-ui/core/TextField';
import { useState } from 'react';


const AddSubTask = ({ addSubTask, setShowSubTaskInput }) => {
    //state that listens for changes in textfield
    const [subTaskText, setSubTaskText] = useState('');

    //this function will handle form submission for adding tasks
    const handleSubmit = (e) => {
        //prevents page from reloading on submit
        e.preventDefault();
        //this function is located in Task.js
        addSubTask(subTaskText);
        //hide form after submission
        setShowSubTaskInput(false);
    }

    return (
        <>
            <form onSubmit={handleSubmit} className="subtaskform">
                <TextField required type="text" label="Subtask" onChange={e => setSubTaskText(e.target.value)} ></TextField>
                <Button type="submit"><AddCircleIcon /></Button>
            </form>
        </>
    )
}

export default AddSubTask
