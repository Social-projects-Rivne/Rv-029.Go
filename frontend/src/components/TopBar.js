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
import * as topBarActions from '../actions/TopBarActions'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

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

const TopBar = ({ classes, topBarState, action, ownProps, ...decorator }) => {

    if (!sessionStorage.getItem('token')) {
        browserHistory.push("/authorization/login")
    }

    const logOut = () => {
        sessionStorage.removeItem('token')
        browserHistory.push("/authorization/login")
    }

    const toggleDrawer = () => {
        action.toggleDrawer(!topBarState.isDrawerOpen)
    }

    return (
        <div className={classes.root}>
            <Drawer open={topBarState.isDrawerOpen} onClose={toggleDrawer}>
                <div
                    tabIndex={0}
                    role="button"
                    onClick={toggleDrawer}
                    onKeyDown={toggleDrawer}
                >
                    Some menu should be here
                </div>
            </Drawer>
            <AppBar position="static">
                <Toolbar>
                    <IconButton className={classes.menuButton} color="inherit" aria-label="Menu" onClick={toggleDrawer}>
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


const mapStateToProps = (state, ownProps) => {
    return {
        topBarState: state.topBar,
        ownProps
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        action: bindActionCreators(topBarActions, dispatch)
    }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(TopBar)
)
