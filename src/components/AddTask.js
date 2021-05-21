import React, {useState} from "react";
import DatePicker from 'react-date-picker';
import TimePicker from 'react-time-picker';
import RepeatPicker from "./RepeatPicker";
import AddSubForm from "./AddSubForm"
import discordImage from "./discord.png";
import MailOutlineIcon from '@material-ui/icons/MailOutline';
import moment from "moment";


const AddTask = (props) => {

    const [userLists, setUserLists] = useState(props.userLists)
    const [templateValue, setTemplateValue] = useState('');
    const [id, setId] = useState(props.id ? props.id : '');
    const [taskValue, setTaskValue] = useState(props.text ? props.text : "");
    const [dateValue, setDateValue] = useState(props.date ? new Date(props.date) : new Date());
    const [listValue, setListValue] = useState(props.list ? props.list : "Main");
    const [willRepeat, setWillRepeat] = useState(props.willRepeat ? props.willRepeat : false);
    const [showTime, toggleShowTime] = useState(props.reminder ? props.reminder : false);
    const [time, setTime] = useState(props.time ? moment(props.time).format("HH:mm") : moment(new Date()).format("HH:mm"));
    const [repeatFrequency, setRepeatFrequency] = useState(props.repeatFrequency ? props.repeatFrequency : "Never");
    const [emailSelected, setEmailSelected] = useState(props.emailSelected ? props.emailSelected : props.defaultReminders['email']);
    const [discordSelected, setDiscordSelected] = useState(props.discordSelected !== undefined ? props.discordSelected : props.defaultReminders['discord']);
    const [endRepeat, setEndRepeat] = useState(props.endRepeat ? new Date(props.endRepeat) : new Date());
    const [showSubtasks, setShowSubtasks] = useState(false);
    const [subtaskValue, setSubtaskValue] = useState('');
    
    //array for subtasks to be pushed with new task
    const [subtaskArr, setSubtaskArr] = useState(props.subtasks ? props.subtasks : []);

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
                {showTime && <TimePicker format="h:mm a" onChange={setTime} value={time} className="addTaskInput" disableClock={true} clearIcon={null} clockIcon={null}/>}
                {showTime && <MailOutlineIcon className={emailSelected ? "outlined icon" : "icon"} onClick={() => setEmailSelected(!emailSelected)}/>}
                {showTime && <img alt="discord icon" className={discordSelected ? "discord outlined icon" : "discord icon"} onClick={() => setDiscordSelected(!discordSelected)} src={discordImage}/>}
                {willRepeat && <RepeatPicker 
                    frequency={repeatFrequency} changeRepeats={(e) => setRepeatFrequency(e.target.value)} endDate={endRepeat}
                    changeEndDate={setEndRepeat}/>}
                <label className="addTaskInput">List: &nbsp;
                    <select value={listValue} onChange={(e) => setListValue(e.target.value)}>
                        {userLists.map((name, id) => (
                            <option value={name[0]} key={id[1]}>{name[0]}</option>
                        ))}
                    </select>
                </label>
                <div>
                    Subtasks? &nbsp;
                    <input type="checkbox" 
                        checked={showSubtasks}
                        onChange={() => setShowSubtasks(!showSubtasks)}/> &nbsp;
                </div>
                <div className="subtaskSectionAdd">
                    {showSubtasks && <AddSubForm setSubtaskValue={setSubtaskValue} subtaskValue=    {subtaskValue} subtaskArr={subtaskArr} setSubtaskArr={setSubtaskArr}/>}
                    {(subtaskArr && showSubtasks) ? subtaskArr.map((subTask) => (<p>{subTask}</p>)) : ''}
                </div>
                <div className="addTaskInput">
                    <button className="addTaskInput addTaskButton" onClick={() => { 
                        console.log("time:")
                        console.log(time);
                        console.log("date:")
                        console.log(dateValue)
                        let newTime;
                        if (time.slice(-2) == "AM") {
                            newTime = time.slice(0, -3);
                        } else if (time.slice(-2) == "PM") {
                            newTime = time.slice(0, -3);
                            newTime = parseInt(newTime.slice(0,2) + 12).toString + newTime.slice(2)
                        } else {
                            newTime = time;
                        }
                        let timeString;
                        if (parseInt(newTime.slice(0,2)) == 0) {
                            timeString = "12" + newTime.slice(2) + ":00 AM"
                        } else if (parseInt(newTime.slice(0,2)) == 12) {
                            timeString = (parseInt(newTime.slice(0,2))).toString() + newTime.slice(2) + ":00 PM"
                        }else if (parseInt(newTime.slice(0,2)) > 12) {
                            timeString = (parseInt(newTime.slice(0,2))-12).toString() + newTime.slice(2) + ":00 PM"
                        } else {
                            timeString = newTime + ":00 AM"
                        }

                        timeString = timeString.padStart(11, '0')

                        console.log("updated Time:")
                        console.log(timeString);
                        props.onAdd({
                            id: id,
                            date: moment(dateValue).format("MM/DD/YYYY") + " " + timeString,
                            discordSelected: discordSelected,
                            emailSelected: emailSelected,
                            end_repeat: endRepeat,
                            list: listValue,
                            remind: showTime,
                            reminder_time: time,
                            repeatFrequency: repeatFrequency,
                            text: taskValue,
                            willRepeat: willRepeat
                        }
                    )}}>Add</button>
                    <button className="addTaskInput addTaskButton" onClick={() => props.onCancel()}>Cancel</button>
                </div>
            </div>
            <div className="popoverTag1 popoverRight1"></div>
            <div className="popoverTag2 popoverRight2"></div>
        </div>
    );
}

export default AddTask
