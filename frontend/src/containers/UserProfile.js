import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {withStyles} from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import ProjectCard from '../components/ProjectCard'
import BoardCard from '../components/BoardCard'
import * as defaultPageActions from "../actions/DefaultPageActions"
import {connect} from 'react-redux'
import {bindActionCreators} from 'redux'
import {API_URL} from "../constants/global"
import messages from "../services/messages"
import axios from "axios"
import * as userActions from "../actions/UserActions"
import Paper from 'material-ui/Paper'
import Typography from 'material-ui/Typography'

class ViewUserProfile extends Component {

  // componentDidMount() {
  //   this.getUserInfo()
  // }

  getUserInfo = () => {
    axios.get(API_URL + `profile/${this.props.ownProps.params.id}`)
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

    const {classes, boards} = this.props
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
                  <li>
                    <Grid Item>
                        <Typography variant="headline" gutterBottom component="h2">
                          Nigga Shit
                        </Typography>
                    </Grid>
                  </li>
                    <br/>
                  <li>
                    <Grid Item>
                        <Typography variant="headline" gutterBottom component="h2">
                          owner@gmail.com
                        </Typography>
                    </Grid>
                  </li>
                    <br/>
                  <li>
                    <Grid Item>
                        <Typography variant="headline" gutterBottom component="h2">
                          User
                        </Typography>
                    </Grid>
                  </li>
                    <br/>
                  <li>
                    <Grid Item>
                        <Typography variant="headline" gutterBottom component="h2">
                          Projects:
                        </Typography>
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
  wrapper: {
    maxWidth: 400,
  },
  paper: {
    padding: theme.spacing.unit * 2,
    height: "100%",
    width: "300px",
  },
  greed:{
    width:"100%",
    height: "100vh",
    paddingTop:"2em",
  },
  pos: {
    marginBottom: 12,
    color: theme.palette.text.secondary,
  },
  list: {
    listStyleType: "none",  
  },
});

const mapStateToProps = (state, ownProps) => {
  return {
    profiles: state.profiles,
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

export default withStyles(styles)(
  connect(mapStateToProps, mapDispatchToProps)(ViewUserProfile)
)
