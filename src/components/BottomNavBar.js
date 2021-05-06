// component for bottom nav bar
import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import BottomNavigation from '@material-ui/core/BottomNavigation';
import BottomNavigationAction from '@material-ui/core/BottomNavigationAction';
import VisibilityIcon from '@material-ui/icons/Visibility';
import SettingsIcon from '@material-ui/icons/Settings';
import ControlPointIcon from '@material-ui/icons/ControlPoint';

const useStyles = makeStyles({
    root: {
        width: '100%',
        position: 'sticky',
        bottom: 0,
        left: 0,
    },
});

const BottomNavBar = (props) => {
    const classes = useStyles();
    const [value, setValue] = React.useState(0);

    return (
        <BottomNavigation
            value={value}
            onChange={(event, newValue) => {
                setValue(newValue);
            }}
            showLabels
            className={classes.root}
        >
            <BottomNavigationAction label="" icon={<VisibilityIcon />} />
            <BottomNavigationAction label="" icon={<SettingsIcon />} />
            <BottomNavigationAction label="" icon={<ControlPointIcon />} onClick={props.onAddTask}/>
        </BottomNavigation>
    );
}

export default BottomNavBar
