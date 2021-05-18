import React, {useState} from "react";
import discordImage from "./discord.png";
import MailOutlineIcon from '@material-ui/icons/MailOutline';


const Options = (props) => {

    const [listValue, setListValue] = useState(props.defaultList);
    const [emailSelected, setEmailSelected] = useState(props.defaultReminders['email']);
    const [discordSelected, setDiscordSelected] = useState(props.defaultReminders['discord']);

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
                        <MailOutlineIcon className={emailSelected ? "outlined icon" : "icon"} onClick={() => setEmailSelected(!emailSelected)}/>
                        <img className={discordSelected ? "discord outlined icon" : "discord icon"} onClick={() => setDiscordSelected(!discordSelected)} src={discordImage}/>
                    </div>
                </div>
            </div>
            <div className="popoverTag1 popoverCenter1"></div>
            <div className="popoverTag2 popoverCenter2"></div>
        </div>
    );
}

export default Options
