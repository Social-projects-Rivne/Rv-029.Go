import React, { Component } from 'react'
import PropTypes from 'prop-types'
import InjectObjSize from '../decorators/objSize'
import {withStyles} from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import ProjectCard from '../components/ProjectCard'
import BoardCard from '../components/BoardCard'
import * as defaultPageActions from "../actions/DefaultPageActions"
import {connect} from 'react-redux'
import { Link, browserHistory} from 'react-router'
import {bindActionCreators} from 'redux'
import {API_URL} from "../constants/global"
import messages from "../services/messages"
import axios from "axios"
import * as userActions from "../actions/UsersActions"
import Paper from 'material-ui/Paper'
import Typography from 'material-ui/Typography'
import Button from 'material-ui/Button'
import AppBar from 'material-ui/AppBar'
import Tabs, { Tab } from 'material-ui/Tabs'
import Divider from 'material-ui/Divider'
import GridList, { GridListTile, GridListTileBar } from 'material-ui/GridList'
import Avatar from 'material-ui/Avatar'
import Chip from 'material-ui/Chip'
import FaceIcon from 'material-ui-icons/Face'
import initialState from '../reducers/users'
import Dialog, {
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from 'material-ui/Dialog'

class ViewUserProfile extends Component {



  state = {
    file: null,
    avatar: null
  }

  handleFileSelected = (e) => {
    this.props.userActions.setUsersPhotoToImport(e.target.files[0]);
    // Assuming only image
      var file = e.target.files[0];

      this.setState({ file })


      console.log(url) // Would see a path?
      // TODO: concat files
  };
  
  handleImportUsers = () => {
    const formData = new FormData();
    formData.append('image', this.props.user.file);
  
    const config = {
        headers: {
            'content-type': 'multipart/form-data',
            'Authorization': 'Bearer ' + sessionStorage.getItem('token'),
        }
    };
    
    axios.post(API_URL + `profile/photo`, formData, config)
    .then((response) => {
      this.handleUsersClose()
      let user = this.props.user.userInfo
      user.Photo = response.data.Data      
      this.props.userActions.setCurrentUser(user)
    })
    .catch((error) => {
      if (error.response && error.response.data.Message) {
        messages(error.response.data.Message)
      } else {
        messages("Server error occured")
      }
    })

    if (this.state.file) {
      let reader = new FileReader();
      let url = reader.readAsDataURL(this.state.file);
    
       reader.onloadend = function(e) {
          this.setState({
              avatar: [reader.result]
          })
        }.bind(this);
    }
  };
  
  handleUsersOpen = () => {
    console.log('method')
    this.props.userActions.toggleUsersDialog(true);
  };
  
  handleUsersClose = () => {
    this.props.userActions.toggleUsersDialog(false);
  };  

  componentDidMount() {
    this.getUserInfo()

      

    console.log(this.props.user.userInfo)
    this.props.defaultPageActions.changePageTitle("Profile")
  }

  sortUserProjects = (userInfo) => {

    let { Projects } = userInfo,
        projectsArr = []

    for (let key in Projects) {
      projectsArr.push({ID: key, name: Projects[key]})
    }

    projectsArr.sort((a, b) => {
      if(a.name < b.name) return -1;
      if(a.name > b.name) return 1;
      return 0;
    })

    userInfo.Projects = projectsArr

    return userInfo
  }

  getUserInfo = () => {
    axios.get(API_URL + `profile/${this.props.ownProps.params.id}`)
      .then((response) => {
        this.props.userActions.setCurrentUser(this.sortUserProjects( response.data.Data )) 
        this.setState({
          avatar: response.data.Data.Photo
        })
      })
      .catch((error) => {
        if (error.response && error.response.data.Message) {
          messages(error.response.data.Message)
        } else {
          messages("Server error occured")
        }
      });
  }

  static propTypes = {
    classes: PropTypes.object.isRequired,
  }

  render () {
    console.log(userActions)
    const {classes, user} = this.props

    return (
      <Grid container
            allignitems={"center"}
            justify={"center"}
           >
            <Paper className={classes.paper}>
                    <Grid container>
                        <Grid item>
                          <img  className={classes.img} src={(this.state.avatar) ? (this.state.avatar) : ('static/default.png')}/>
                        </Grid>
                        <Grid item>
                            <Typography variant="headline" gutterBottom component="h2">
                              {(user.userInfo) ? (user.userInfo.FirstName) : ('')}    {(user.userInfo) ? (user.userInfo.LastName) : ('')}  
                            </Typography>
                            <Typography variant="headline" gutterBottom component="h2">
                              {(user.userInfo) ? (user.userInfo.Email) : ('')}
                            </Typography>
                            <Typography variant="headline" gutterBottom component="h2">
                              {(user.userInfo) ? (user.userInfo.Role) : ('')}
                            </Typography>
                            {(this.props.ownProps.params.id == 'own') ? (

                            <div>  
                              <Link className={classes.link} to={'/profile/own/update'}>         
                                <Button variant="raised" color="primary">
                                  Edit Info
                                </Button>
                              </Link>

                              <Link onClick={this.handleUsersOpen} className={classes.link}>
                                <Button variant="raised" color="primary">
                                  Change Avatar
                                </Button>
                              </Link>
                            </div>

                            ) : (null)}

                        </Grid>   
                      </Grid>

                  <Grid container  
                   allignitems={"center"}
                    justify={"center"}>

                    <Grid item>
                    <Typography variant="headline" gutterBottom component="h2">
                        Projects: 
                    </Typography>
                    </Grid>

                      <Grid item className={classes.chipContainer}>

                        {(user.userInfo && user.userInfo.Projects.length) ? (null) : (
                          <Typography>
                            No projects found
                          </Typography>    
                        )}

                        {

                          (user.userInfo) ? (
                            user.userInfo.Projects.map((item, i) => {

                              
                              return (
                              <Link className={classes.link} key={i} to={'/project/'+item.ID}>             
                                <Chip
                                  className={classes.chip}
                                  label={item.name}
                              />
                              </Link>
                              )
                        })
                      ) : (null)
                      
                      }
                  </Grid>
                  </Grid>
                          
                        
            </Paper>

            <Dialog
                open={user.isUsersOpened}
                onClose={this.handleUsersClose}
                aria-labelledby="form-dialog-title" >
                <DialogTitle id="form-dialog-title">Select image file</DialogTitle>
                    <DialogContent>
                        <div>
                            <input
                                type="file" name="file" id="file"
                                className="input-file"
                                onChange={this.handleFileSelected}
                                accept="image/jpeg, image/png"
                            />
                        </div>
                    </DialogContent>
                <DialogActions>
                    <Button onClick={this.handleUsersClose} color="default">
                        Cancel
                    </Button>
                    <Button 
                    onClick={this.handleImportUsers} 
                    color="primary">
                        Import
                    </Button>
                </DialogActions>
            </Dialog>
      </Grid>
    );
  }
}


const styles = {
  chipContainer: {
    width: 500,
  },
  chip: {
    margin: '.1em .5em',
    // display: 'flex'
  },
  paper: {
    margin: '2em 0',
    padding: '1em'
  },
  img: {
    width: '200px',
    height: '200px',    
  },  
  link: {
    textDecoration: 'none',
  },
}

const mapStateToProps = (state, ownProps) => {
  return {
    user: state.users,
    defaultPage: state.defaultPage,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch),
    userActions: bindActionCreators(userActions, dispatch)
  }
}

export default InjectObjSize(
  withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(ViewUserProfile)
  )
) 