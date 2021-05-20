import React, {useState} from "react";
import discordImage from "./discord.png";
import MailOutlineIcon from '@material-ui/icons/MailOutline';
import firebase from "firebase";


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


const Options = (props) => {

    const [listValue, setListValue] = useState(props.defaultList);
    const [emailSelected, setEmailSelected] = useState(props.defaultReminders['email']);
    const [discordSelected, setDiscordSelected] = useState(props.defaultReminders['discord']);
    // const logOut = logOut();

    const logOut = () => {
        
    }

    return ( 
        <div>
            <div className="popover" >
                <div className="listHeader">List Options:</div>
                    <div className="optionsOption"><span className="clickableText green">Share</span> this list</div>
                    <div className="optionsOption"><span className="clickableText red">Delete</span> this list</div>

                <div className="listHeader">Global Options:</div>
                    <div className="optionsOption">Default List: &nbsp;
                        <select value={listValue} onChange={(e) => setListValue(e.target.value)}>
                            <option value="Main">Main</option>
                            <option value="Shared">Shared</option>
                            <option value="Other List">Other List</option>
                        </select>
                    </div>
                    <div className="optionsOption">Default Reminder Method: 
                        <div className="remindersContainer">
                            <MailOutlineIcon className={emailSelected ? "outlined icon" : "icon"} onClick={() => {setEmailGlobal(!emailSelected); setEmailSelected(!emailSelected);}}/>
                            <img alt="discord icon" className={discordSelected ? "discord outlined icon" : "discord icon"} onClick={() => setDiscordSelected(!discordSelected)} src={discordImage}/>
                        </div>
                    </div>
                <div className="listHeader">Account Options:</div>
                    <div className="optionsOption">
                        <button onClick={logOut} >Reset Password</button>
                    </div>
            </div>
            <div className="popoverTag1 popoverCenter1"></div>
            <div className="popoverTag2 popoverCenter2"></div>
        </div>
    );
}

export default Options
