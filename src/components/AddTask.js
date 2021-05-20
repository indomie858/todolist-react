import React, {useState} from "react";
import DatePicker from 'react-date-picker';
import TimePicker from 'react-time-picker';
import RepeatPicker from "./RepeatPicker";
import AddSubForm from "./AddSubForm"
import discordImage from "./discord.png";
import MailOutlineIcon from '@material-ui/icons/MailOutline';


const AddTask = (props) => {
    const [templateValue, setTemplateValue] = useState('');
    const [taskValue, setTaskValue] = useState(props.text ? props.text : "");
    const [dateValue, setDateValue] = useState(props.date ? new Date(props.date) : new Date());
    const [listValue, setListValue] = useState(props.list ? props.list : "Main");
    const [willRepeat, setWillRepeat] = useState(props.willRepeat ? props.willRepeat : false);
    const [showTime, toggleShowTime] = useState(props.reminder ? props.reminder : false);
    const [time, setTime] = useState(props.date ? new Date(props.date) : new Date());
    const [repeatFrequency, setRepeatFrequency] = useState(props.repeatFrequency ? props.repeatFrequency : "Every");
    const [numDays, setNumDays] = useState(props.numDays ? props.numDays : 1);
    const [emailSelected, setEmailSelected] = useState(props.emailSelected ? props.emailSelected : props.defaultReminders['email']);
    const [discordSelected, setDiscordSelected] = useState(props.discordSelected !== undefined ? props.discordSelected : props.defaultReminders['discord']);
    const [showSubtasks, setShowSubtasks] = useState(false);
    const [subtaskValue, setSubtaskValue] = useState('');

    return ( 
        <div>
            <div className="popover" >
                <span>Add New Task:</span>
                <label className="templateInput">Template: &nbsp;
                    <select value={templateValue} onChange={(e) => {
                        //sorry for how nasty this select element is
                        setTemplateValue(e.target.value);
                        if (e.target.value !== ''){
                            toggleShowTime(true);
                            setTaskValue(e.target.value);
                        } else {
                            toggleShowTime(false);
                            setTaskValue('');
                        }                        
                        }}>
                        <option value=""></option>
                        <option value="Study">Study</option>
                        <option value="Clean Room">Clean Room</option>
                        <option value="Do Laundry">Do Laundry</option>
                        <option value="Buy Groceries">Buy Groceries</option>
                        <option value="Relax">Relax</option>
                    </select>
                </label>
                <input 
                    className="textInput"
                    type="text" 
                    value={taskValue} 
                    onChange={(e) => setTaskValue(e.target.value)} autoFocus/>
                <DatePicker value={dateValue}  onChange={setDateValue} className="addTaskInput"/>
                <div>Repeat? &nbsp;
                    <input type="checkbox" checked={willRepeat} onChange={() => setWillRepeat(!willRepeat)}/>&nbsp;
                    Remind? &nbsp;
                    <input type="checkbox"
                        checked={showTime} 
                        onChange={() => toggleShowTime(!showTime)}/>
                </div>
                {showTime && <TimePicker onChange={setTime} value={time} className="addTaskInput" disableClock={true} clearIcon={null} clockIcon={null}/>}
                {showTime && <MailOutlineIcon className={emailSelected ? "outlined icon" : "icon"} onClick={() => setEmailSelected(!emailSelected)}/>}
                {showTime && <img alt="discord icon" className={discordSelected ? "discord outlined icon" : "discord icon"} onClick={() => setDiscordSelected(!discordSelected)} src={discordImage}/>}
                {willRepeat && <RepeatPicker 
                    isEvery={repeatFrequency} changeRepeats={(e) => setRepeatFrequency(e.target.value)}
                    numDays={numDays} onChange={(e) => e.target.value > 0 ? setNumDays(e.target.value) : setNumDays(numDays)}/>}
                <label className="addTaskInput">List: &nbsp;
                    <select value={listValue} onChange={(e) => setListValue(e.target.value)}>
                        <option value="Main">Main</option>
                        <option value="Shared">Shared</option>
                        <option value="Other List">Other List</option>
                    </select>
                </label>
                <div>
                    Subtask? &nbsp;
                    <input type="checkbox" 
                        checked={showSubtasks}
                        onChange={() => setShowSubtasks(!showSubtasks)}/> &nbsp;
                </div>
                {showSubtasks && <AddSubForm setSubtaskValue={setSubtaskValue} />}
                <div className="addTaskInput">
                    <button className="addTaskInput addTaskButton" onClick={props.onAdd}>Add</button>
                    <button className="addTaskInput addTaskButton" onClick={props.onCancel}>Cancel</button>
                </div>
            </div>
            <div className="popoverTag1 popoverRight1"></div>
            <div className="popoverTag2 popoverRight2"></div>
        </div>
    );
}

export default AddTask
