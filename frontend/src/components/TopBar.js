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
import SnackBar from './SnackBar'
import ModalNotification from './ModalNotification'
import { Link, browserHistory } from 'react-router';
import * as defaultAction from '../actions/DefaultPageActions';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import auth from '../services/auth'
import Dialog, {
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
} from 'material-ui/Dialog';


const styles = {
    root: {
        width: '100%'
    },
    a: {
        textDecoration: 'none',
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

//TODO: rewrite to component
const TopBar = ({ classes, defaultPage, projects, defaultPageActions, ownProps, ...decorator }) => {

    if (!sessionStorage.getItem('token')) {
        browserHistory.push("/authorization/login")
    }

    const toggleDrawer = () => {
        defaultPageActions.toggleDrawer(!defaultPage.isDrawerOpen)
    };

    //
    const handleClickOpenAddUser = () => {
        defaultPageActions.toggleAddUserToProject(true)
    };
    const handleClose = () => {
        defaultPageActions.toggleAddUserToProject(false);
    };

    let projectBoardsList = null
    if (projects.currentProject !== null) {
        projectBoardsList = <List component="nav">
                <Link
                    to={'project/'+ projects.currentProject.UUID +'/board/create'}
                    className={classes.link}>
                    <ListItem button>
                        <ListItemIcon>
                            <Icon color="primary">add</Icon>
                        </ListItemIcon>
                        <ListItemText primary="Add Board" />
                    </ListItem>
                </Link>
                <Link onClick={handleClickOpenAddUser}
                    className={classes.link}>
                    <ListItem button>
                        <ListItemIcon>
                            <Icon color="primary">add</Icon>
                        </ListItemIcon>
                        <ListItemText primary="Add User To Project" />
                    </ListItem>
                </Link>
            </List>
    }

    return (
        <div className={classes.root}>
            <Drawer open={defaultPage.isDrawerOpen} onClose={toggleDrawer}>
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
                            <Link
                                to="project/create"
                                className={classes.link}>
                                <ListItem button>
                                    <ListItemIcon>
                                        <Icon color="primary">add</Icon>
                                    </ListItemIcon>
                                    <ListItemText primary="Add Project" />
                                </ListItem>
                            </Link>
                        </List>
                        <Divider />
                        {projectBoardsList}
                        <Divider />
                        <List component="nav">
                            <Link className={classes.a} to={'/profile/'}>
                                <ListItem button>
                                    <ListItemText primary="Profile" />
                                </ListItem>
                            </Link>
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
                        {defaultPage.pageTitle}
                    </Typography>
                    <Button color="inherit" onClick={auth.logOut} >Logout</Button>
                </Toolbar>
            </AppBar>
            <SnackBar
                errorMessage={defaultPage.errorMessage}
                setErrorMessage={defaultPageActions.setErrorMessage}/>

            <ModalNotification
                title='Notification'
                content={defaultPage.notificationMessage}
                setNotificationMessage={defaultPageActions.setNotificationMessage}/>

            {/* ADD USER TO PROJECT DIALOG */}
            <Dialog
                open={defaultPage.isUserToProjectOpen}
                onClose={handleClose}
                aria-labelledby="form-dialog-title" >
                <DialogTitle id="form-dialog-title">Add user</DialogTitle>
                <DialogContent>
                    <List>
                        {(projects.currentProjectUsers.length > 0) ? (
                            projects.currentProjectUsers.map((item, i) => (
                                <ListItem button onClick={() => this.handleListItemClick(item)} key={i}>
                                    <ListItemAvatar>
                                        <Avatar className={classes.avatar}>
                                            <PersonIcon />
                                        </Avatar>
                                    </ListItemAvatar>
                                    <ListItemText primary={item.Email}  />
                                </ListItem>
                            ))
                        ) : ("No Users")}
                    </List>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose} color="primary">
                        Cancel
                    </Button>
                </DialogActions>
            </Dialog>
        </div>


    );
}

TopBar.propTypes = {
    classes: PropTypes.object.isRequired,
};


const mapStateToProps = (state, ownProps) => {
    return {
        defaultPage: state.defaultPage,
        projects: state.projects,
        ownProps
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        defaultPageActions: bindActionCreators(defaultAction, dispatch),
    }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(TopBar)
)
