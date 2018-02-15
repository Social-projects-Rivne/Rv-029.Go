import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import Typography from 'material-ui/Typography';
import Button from 'material-ui/Button';
import IconButton from 'material-ui/IconButton';
import MenuIcon from 'material-ui-icons/Menu';
import Drawer from 'material-ui/Drawer';
import {browserHistory} from "react-router";

const styles = {
    root: {
        width: '100%',
    },
    flex: {
        flex: 1,
    },
    menuButton: {
        marginLeft: -12,
        marginRight: 20,
    },
};

function TopBar(props) {
    const { classes } = props;

    if (!sessionStorage.getItem('token')) {
        browserHistory.push("/authorization/login")
    }

    const logOut = () => {
        sessionStorage.removeItem('token')
        browserHistory.push("/authorization/login")
    }

    const toggleDrawer = (state) => {
        return !state
    }

    return (
        <div className={classes.root}>
            <Drawer open={false} onClose={toggleDrawer()}>
                <div
                    tabIndex={0}
                    role="button"
                    onClick={toggleDrawer()}
                    onKeyDown={toggleDrawer()}
                >
                    Something
                </div>
            </Drawer>
            <AppBar position="static">
                <Toolbar>
                    <IconButton className={classes.menuButton} color="inherit" aria-label="Menu" onClick={toggleDrawer()}>
                        <MenuIcon />
                    </IconButton>
                    <Typography variant="title" color="inherit" className={classes.flex}>
                        Title
                    </Typography>
                    <Button color="inherit" onClick={logOut} >Logout</Button>
                </Toolbar>
            </AppBar>
        </div>
    );
}

TopBar.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(TopBar);