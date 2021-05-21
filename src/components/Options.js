import React, {useState} from "react";
import discordImage from "./discord.png";
import MailOutlineIcon from '@material-ui/icons/MailOutline';
import firebase from "firebase";
import { useHistory } from "react-router-dom";


  


async function setEmailGlobal(emailState){
    let text = "a3a1hWUx5geKB8qeR6fbk5LZZGI2"
    // let options = "emailNotifications="+emailState
    fetch('http://localhost:3003/api/update/'+text,{
                    method: 'POST',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        //pass in value of input text in body of request
                        update: 'userSettings',
                        emailNotifications: {emailState},
                        
                    })
            
                }).then(response => {
                    if(response.status===404){
                        return "Error: 404"
                    }else{
                        return response.json()}
                }).then(data=>JSON.stringify(data));
    }
async function setDefaultListGlobal(listValue){
    let text = "a3a1hWUx5geKB8qeR6fbk5LZZGI2"
    // let options = "emailNotifications="+emailState
    fetch('http://localhost:3003/api/update/'+text,{
                    method: 'POST',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        //pass in value of input text in body of request
                        update: 'userSettings',
                        default_list: listValue
                        
                    })
            
                }).then(response => {
                    if(response.status===404){
                        return "Error: 404"
                    }else{
                        return response.json()}
                }).then(data=>JSON.stringify(data));
    }


const Options = (props) => {
    const [showShareInput, setShowShareInput] = useState(false);
    const [shareListInput, setShareListInput] = useState('');
    const [listValue, setListValue] = useState(props.defaultList);
    const [emailSelected, setEmailSelected] = useState(props.defaultReminders['email']);
    const [discordSelected, setDiscordSelected] = useState(props.defaultReminders['discord']);
    const history = useHistory();

    function returnLogin() {
        history.push("/login");
    }
    // const logOut = logOut();

    const logOut = () => {
        firebase.auth().signOut().then(() => {
            // Sign-out successful.
            console.log("The User has been logged out.")
            sessionStorage.setItem('token', "");
            returnLogin()
        }).catch((error) => {
            // An error happened.
            console.log(error)
        });
    }

    function sendOptions(partialUserObject) {
        console.log("sneding")
        console.log(partialUserObject.discord_reminder)

        props.onChooseOption({
            default_list: partialUserObject.hasOwnProperty('default_list') ? partialUserObject.default_list : listValue,
            discord_reminder: partialUserObject.hasOwnProperty('discord_reminder') ? partialUserObject.discord_reminder : discordSelected,
            email_reminder: partialUserObject.hasOwnProperty('email_reminder') ? partialUserObject.email_reminder : emailSelected
        })
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        //handle share submit here
        //variable for input is shareListInput
        alert('Shared list')
    }

    return ( 
        <div>
            <div className="popover" >
                <div className="listHeader">List Options:</div>
                    <div className="optionsOption"><span className="clickableText green" onClick={() => setShowShareInput(!showShareInput)}>Share</span> this list</div>
                    {showShareInput && 
                       <form className="addSubForm" onSubmit={handleSubmit}>
                           <input type="text" placeholder="User email" onChange={(e)=> setShareListInput(e.target.value)} />
                           <button type="submit">Share</button>
                       </form> 
                    }
                    <div className="optionsOption"><span className="clickableText red">Delete</span> this list</div>

                <div className="listHeader">Global Options:</div>
                    <div className="optionsOption">Default List: &nbsp;
                        <select value={listValue} onChange={(e) => {setListValue(e.target.value); sendOptions({default_list: e.target.value});}}>
                            {props.userLists.map((name, id) => (
                                <option value={name[0]} key={id[1]}>{name[0]}</option>
                            ))}
                            
                        </select>
                    </div>
                    <div className="optionsOption">Default Reminder Method: 
                        <div className="remindersContainer">
                            <MailOutlineIcon className={emailSelected ? "outlined icon" : "icon"} onClick={() => { const newValue = !emailSelected; setEmailSelected(newValue); sendOptions({email_reminder: newValue});}}/>
                            <img alt="discord icon" className={discordSelected ? "discord outlined icon" : "discord icon"} onClick={() => { const newValue = !discordSelected; setDiscordSelected(newValue); sendOptions({discord_reminder: newValue});}} src={discordImage}/>
                        </div>
                    </div>
                <div className="listHeader">Account Options:</div>
                    <div className="optionsOption">
                        <span className="clickableText blue">Reset Password</span></div>
                        <button className="addTaskInput addTaskButton" onClick={() => logOut()}>Logout</button>
            </div>
            <div className="popoverTag1 popoverCenter1"></div>
            <div className="popoverTag2 popoverCenter2"></div>
        </div>
    );
}

export default Options
