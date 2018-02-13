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
import List, { ListItem, ListItemIcon, ListItemText } from 'material-ui/List';
import Divider from 'material-ui/Divider';
import Icon from 'material-ui/Icon';
import { Link, browserHistory } from 'react-router';
import * as topBarActions from '../actions/TopBarActions';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import auth from '../services/auth'

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
    listWrapper: {
        minWidth: 250,
        maxWidth: 360,
    },
    link: {
        textDecoration: 'none'
    }
};

const TopBar = ({ classes, topBarState, action, ownProps, ...decorator }) => {

    if (!sessionStorage.getItem('token')) {
        browserHistory.push("/authorization/login")
    }

    const toggleDrawer = () => {
        action.toggleDrawer(!topBarState.isDrawerOpen)
    }

    let projectBoardsList = null
    if (topBarState.currentProject !== null) {
        projectBoardsList = <List component="nav">
            {topBarState.currentBoardProjects.map((value, index) => (
                <ListItem button key={value.id}>
                    <ListItemText primary={value.name} />
                </ListItem>
            ))}
        </List>
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
                    <div className={classes.listWrapper}>
                        <List component="nav">
                            <Link
                                to="projects"
                                className={classes.link}>
                                <ListItem button>
                                    <ListItemIcon>
                                        <Icon color="primary">dashboard</Icon>
                                    </ListItemIcon>
                                    <ListItemText primary="Projects" />
                                </ListItem>
                            </Link>
                        </List>
                        <Divider />
                        {projectBoardsList}
                        <Divider />
                        <List component="nav">
                            <ListItem button>
                                <ListItemText primary="Logout" onClick={auth.logOut} />
                            </ListItem>
                        </List>
                    </div>
                </div>
            </Drawer>
            <AppBar position="static">
                <Toolbar>
                    <IconButton className={classes.menuButton} color="inherit" aria-label="Menu" onClick={toggleDrawer}>
                        <MenuIcon />
                    </IconButton>
                    <Typography variant="title" color="inherit" className={classes.flex}>
                        {topBarState.pageTitle}
                    </Typography>
                    <Button color="inherit" onClick={auth.logOut} >Logout</Button>
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
