// component for bottom nav bar
//uses material ui components - Material UI - https://material-ui.com/

import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import BottomNavigation from '@material-ui/core/BottomNavigation';
import BottomNavigationAction from '@material-ui/core/BottomNavigationAction';
import ListIcon from '@material-ui/icons/List';
import SettingsIcon from '@material-ui/icons/Settings';
import ControlPointIcon from '@material-ui/icons/ControlPoint';

const useStyles = makeStyles({
    root: {
        background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
        width: '100%',
        position: 'fixed',
        bottom: 0,
        left: 0,
    },
});

const BottomNavBar = (props) => {
    const classes = useStyles();
    const [value, setValue] = React.useState(0);

    return (
        // bottom navigation bar component from material UI https://material-ui.com/components/bottom-navigation/
        <BottomNavigation
            value={value}
            onChange={(event, newValue) => {
                setValue(newValue);
            }}
            showLabels
            className={classes.root}
        >
            <BottomNavigationAction label="" icon={<ListIcon />} onClick={props.onListNav}/>
            <BottomNavigationAction label="" icon={<SettingsIcon />} onClick={props.onOptions}/>
            <BottomNavigationAction label="" icon={<ControlPointIcon />} onClick={props.onAddTask}/>
        </BottomNavigation>
    );
}

export default BottomNavBar
