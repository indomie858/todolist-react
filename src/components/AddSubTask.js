import Button from '@material-ui/core/Button';
import AddCircleIcon from '@material-ui/icons/AddCircle';
import TextField from '@material-ui/core/TextField';

const AddSubTask = () => {
    return (
        <>
            <TextField type="text"></TextField>
            <Button><AddCircleIcon /></Button>
        </>
    )
}

export default AddSubTask
