import React, {useState} from "react";
import ListNavItem from "./ListNavItem.js"


const ListNav = (props) => {



    return ( 
        <div>
            <div className="popover" >
                <span>Change List:</span>
                {props.lists.map((name) => <ListNavItem key={Math.random().toString(36).substr(2, 9)} list={name}/>)}
            </div>
            <div className="popoverTag1 popoverLeft1"></div>
            <div className="popoverTag2 popoverLeft2"></div>
        </div>
    );
}

export default ListNav
