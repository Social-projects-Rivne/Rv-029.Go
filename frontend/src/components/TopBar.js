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
import List, { ListItem, ListItemIcon, ListItemText, ListItemAvatar } from 'material-ui/List';
import Avatar from 'material-ui/Avatar';
import PersonIcon from 'material-ui-icons/Person';
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
import {API_URL} from "../constants/global";
import messages from "../services/messages";
import axios from "axios/index";
import FormInput from "./FormInput";


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
    },
    exampleImageInput: {
        cursor: 'pointer',
        position: 'absolute',
        top: '0',
        bottom: '0',
        right: '0',
        left: '0',
        width: '100%',
        opacity: '0'
    }
};

//TODO: rewrite to component
const TopBar = ({ classes, defaultPage, projects, defaultPageActions, ownProps, ...decorator }) => {

    if (!sessionStorage.getItem('token')) {
        browserHistory.push("/authorization/login")
    }

    if (defaultPage.permissions.length == 0) {
        axios.get(API_URL + `permissions/list`)
            .then((response) => {
                defaultPageActions.setPermissionsList(response.data.Data)
            })
            .catch((error) => {
                if (error.response && error.response.data.Message) {
                    messages(error.response.data.Message)
                } else {
                    messages("Server error occured")
                }
            });
    }

    if (defaultPage.roles.length == 0) {
        axios.get(API_URL + `roles/list`)
            .then((response) => {
                defaultPageActions.setRolesList(response.data.Data)
            })
            .catch((error) => {
                if (error.response && error.response.data.Message) {
                    messages(error.response.data.Message)
                } else {
                    messages("Server error occured")
                }
            });
    }

    const toggleDrawer = () => {
        defaultPageActions.toggleDrawer(!defaultPage.isDrawerOpen);
    };

    //
    const handleClickOpenAddUser = () => {
        defaultPageActions.toggleAddUserToProject(true);
    };
    const handleUserToProjectClose = () => {
        defaultPageActions.toggleAddUserToProject(false);
    };

    const handleUserToProjectWithPermissionsClose = () => {
        defaultPageActions.togglePermissionsDialog(false);
    };

    const handleListItemClick = (item) => {
        defaultPageActions.togglePermissionsDialog(true, item);
    };

    const handleFileSelected = (e) => {
        defaultPageActions.setUsersFileToImport(e.target.files[0]);
    };

    const handleImportUsers = () => {
        const formData = new FormData();

        formData.append('import', defaultPage.file);

        const config = {
            headers: {
                'content-type': 'multipart/form-data',
                'Authorization': 'Bearer ' + sessionStorage.getItem('token'),
            }
        };

        return axios.post(API_URL + "user/import", formData, config)
    };

    const handleImportUsersOpen = () => {
        defaultPageActions.toggleImportUsersDialog(true);
    };

    const handleImportUsersClose = () => {
        defaultPageActions.toggleImportUsersDialog(false);
    };

    const handleAddUserToProject = (role) => {
        handleUserToProjectClose()
        handleUserToProjectWithPermissionsClose()
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
                            <Link onClick={handleImportUsersOpen} className={classes.link}>
                                <ListItem button>
                                    <ListItemIcon>
                                        <Icon color="primary">add</Icon>
                                    </ListItemIcon>
                                    <ListItemText primary="Import Users From CSV" />
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
                onClose={handleUserToProjectClose}
                aria-labelledby="form-dialog-title" >
                <DialogTitle id="form-dialog-title">Add user</DialogTitle>
                <DialogContent>
                    <List>
                        {(projects.users.length > 0) ? (
                            projects.users.map((item, i) => (
                                <ListItem button onClick={() => handleListItemClick(item)} key={i}>
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
                    <Button onClick={handleUserToProjectClose} color="primary">
                        Cancel
                    </Button>
                </DialogActions>
            </Dialog>
            {/* ADD USER TO PROJECT WITH PERMISSIONS DIALOG */}
            <Dialog
                open={defaultPage.isUserToProjectPermissionsOpen}
                onClose={handleUserToProjectWithPermissionsClose}
                aria-labelledby="form-dialog-title" >
                <DialogTitle id="form-dialog-title">Add user</DialogTitle>
                <DialogContent>
                    <List>
                        {(defaultPage.roles.length > 0) ? (
                            defaultPage.roles.map((item, i) => (
                                <ListItem button onClick={() => handleAddUserToProject(item)} key={i}>
                                    { item }
                                </ListItem>
                            ))
                        ) : ("No Roles")}
                    </List>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleUserToProjectWithPermissionsClose} color="primary">
                        Cancel
                    </Button>
                </DialogActions>
            </Dialog>
            {/* IMPORT USERS */}
            <Dialog
                open={defaultPage.isImportUsersOpened}
                onClose={handleImportUsersClose}
                aria-labelledby="form-dialog-title" >
                <DialogTitle id="form-dialog-title">Select *.csv file in appropriate format</DialogTitle>
                <DialogContent>
                    <div>
                        <input
                            type="file" name="file" id="file"
                            className="input-file"
                            onChange={handleFileSelected}
                            accept="text/csv"
                        />
                    </div>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleImportUsersClose} color="default">
                        Cancel
                    </Button>
                    <Button onClick={handleImportUsers} color="primary">
                        Import
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
