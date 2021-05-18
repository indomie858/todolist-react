import React, {useState} from "react";
import DatePicker from 'react-date-picker';
import TimePicker from 'react-time-picker';
import RepeatPicker from "./RepeatPicker";
import discordImage from "./discord.png";
import MailOutlineIcon from '@material-ui/icons/MailOutline';


const AddTask = (props) => {

    const [taskValue, setTaskValue] = useState("");
    const [dateValue, setDateValue] = useState(new Date());
    const [listValue, setListValue] = useState("Main");
    const [willRepeat, setWillRepeat] = useState(false);
    const [showTime, toggleShowTime] = useState(false);
    const [time, setTime] = useState("1:00");
    const [repeatFrequency, setRepeatFrequency] = useState("Every");
    const [numDays, setNumDays] = useState(1);
    const [emailSelected, setEmailSelected] = useState(props.defaultReminders['email']);
    const [discordSelected, setDiscordSelected] = useState(props.defaultReminders['discord']);

    return ( 
        <div>
            <div className="popover" >
                <span>Add New Task:</span>
                <input type="text" value={taskValue} onChange={(e) => setTaskValue(e.target.value)}/>
                <DatePicker value={dateValue}  onChange={setDateValue} className="addTaskInput"/>
                <div>Repeat? &nbsp;
                    <input type="checkbox" checked={willRepeat} onChange={() => setWillRepeat(!willRepeat)}/>&nbsp;
                    Remind? &nbsp;
                    <input type="checkbox"checked={showTime} onChange={() => toggleShowTime(!showTime)}/>
                </div>
                {showTime && <TimePicker onChange={setTime} value={time} className="addTaskInput" disableClock={true} clearIcon={null} clockIcon={null}/>}
                {showTime && <MailOutlineIcon className={emailSelected ? "outlined icon" : "icon"} onClick={() => setEmailSelected(!emailSelected)}/>}
                {showTime && <img className={discordSelected ? "discord outlined icon" : "discord icon"} onClick={() => setDiscordSelected(!discordSelected)} src={discordImage}/>}
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
                <div className="addTaskInput">
                    <button className="addTaskInput addTaskButton" onClick={props.onAdd}>Add</button>
                </div>
            </div>
            <div className="popoverTag1 popoverRight1"></div>
            <div className="popoverTag2 popoverRight2"></div>
        </div>
    );
}

export default AddTask
