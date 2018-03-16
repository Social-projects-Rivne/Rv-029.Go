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
import * as userActions from "../actions/UserActions"
import Paper from 'material-ui/Paper'
import Typography from 'material-ui/Typography'
import Button from 'material-ui/Button'
import AppBar from 'material-ui/AppBar'
import Tabs, { Tab } from 'material-ui/Tabs'
import Divider from 'material-ui/Divider'
import List, { ListItem, ListItemText } from 'material-ui/List'

class ViewUserProfile extends Component {

  componentDidMount() {
    this.getUserInfo()
    this.props.defaultPageActions.changePageTitle("Profile")
  }

  getUserInfo = () => {
    axios.get(API_URL + `profile`)
      .then((response) => {
        this.props.userActions.setCurrentUser(response.data.Data)
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

    const {classes, user, objSize} = this.props
    let projectsArr = []

    if (user.userInfo) {
      let projects = user.userInfo.Projects

      for (let key in user.userInfo.Projects) {
        projectsArr.push({ID: key, name: projects[key]})
      }

      console.log(projectsArr)
    }

    return (
      <Grid container
            allignitems={"center"}
            justify={"center"}
            className={classes.greed}>
        <div className={classes.root}>
          <div className={classes.wrapper}>
            <Paper className={classes.paper}>
              <Grid container wrap="nowrap">
                <ul className={classes.list}>
                  <div>
                    <li>
                      <Grid item>
                          <img className={classes.img}src={(user.userInfo) ? (user.userInfo.Photo) : ('')}/>
                      </Grid>
                    </li>
                    <li>
                      <Grid item>
                          <Typography variant="headline" gutterBottom component="h2">
                            {(user.userInfo) ? (user.userInfo.FirstName) : ('')}    {(user.userInfo) ? (user.userInfo.LastName) : ('')}  
                          </Typography>
                      </Grid>
                    </li>
                  </div>                  
                    <br/>
                  <li>
                    <Grid item>
                        <Typography variant="headline" gutterBottom component="h2">
                        {(user.userInfo) ? (user.userInfo.Email) : ('')}
                        </Typography>
                    </Grid>
                  </li>
                    <br/>
                  <li>
                    <Grid item>
                        <Typography variant="headline" gutterBottom component="h2">
                        {(user.userInfo) ? (user.userInfo.Role) : ('')}
                        </Typography>
                    </Grid>
                  </li>
                    <br/>
                  <li>
                    <Typography variant="headline" gutterBottom component="h2">
                        Projects: 
                    </Typography>
                    <div className={classes.projects}>
                      <List component="nav">

                        {
                          (user.userInfo) ? (
                            projectsArr.map((item, i) => (
                              <Link className={classes.a} to={'/project/'+item.ID}>     
                                <ListItem button  key={i}>
                                  <ListItemText primary={item.name} />
                                </ListItem>
                              </Link>
                            ))
                          ) : (<h1>loh</h1>)
                        }

                      </List>
                    </div>
                  </li>
                    <br/>
                  <li>
                    <Grid item>
                      <Link className={classes.a} to={'/profile/update'}>                    
                        <Button variant="raised" color="primary">
                          Change information
                        </Button>
                      </Link>
                    </Grid>
                  </li>
                    <br/>
                </ul>
              </Grid>
            </Paper>
          </div>
        </div>
      </Grid>
    );
  }
}


const styles = theme => ({
  root: {
    overflow: 'hidden',
    padding: `0 ${theme.spacing.unit * 3}px`,
  },
  projects: {
    width: '100%',
    maxWidth: '360px',
    backgroundColor: theme.palette.background.paper,
  },
  wrapper: {
    maxWidth: 600,
  },
  paper: {
    padding: theme.spacing.unit * 2,
    height: "100%",
    width: "600px",
  },
  a: {
    textDecoration: 'none',
  },
  greed:{
    width:"100%",
    minHeight: "100vh",
    paddingTop:"2em",
    paddingBottom: '2em',
  },
  pos: {
    marginBottom: 12,
    color: theme.palette.text.secondary,
  },
  photo: {
    paddingTop:"2em",
    paddingBottom: '2em',
    float: 'left',
  },
  list: {
    listStyleType: "none",  
  },
  img: {
    height: '20vh', 
  },
});

const mapStateToProps = (state, ownProps) => {
  return {
    user: state.user,
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