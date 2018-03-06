import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {withStyles} from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import ProjectCard from '../components/ProjectCard'
import BoardCard from '../components/BoardCard'
import * as defaultPageActions from "../actions/DefaultPageActions"
import * as boardsActions from "../actions/BoardsActions"
import * as projectsActions from "../actions/ProjectsActions"
import {connect} from 'react-redux'
import {bindActionCreators} from 'redux'
import {API_URL} from "../constants/global"
import messages from "../services/messages"
import axios from "axios"

class ViewUserProfile extends Component {

  componentWillMount() {
    if (this.props.profiles.currentUser == null) {
        axios.get(API_URL + `profile/${this.props.ownProps.params.id}`)
            .then((response) => {
                this.props.ProfileActions.setCurrentUser(response.data.Data)
                this.props.defaultPageActions.changePageTitle("Profile " + this.props.profiles.currentUser.Name)
            })
            .catch((error) => {
                if (error.response && error.response.data.Message) {
                    messages(error.response.data.Message)
                } else {
                    messages("Server error occured")
                }
            })
    } else {
        this.props.defaultPageActions.changePageTitle("Profile " + this.props.profiles.currentUser.Name)
    }
  }

  componentDidMount() {
    if (this.props.profiles.currentProfiles.length > 0) {
      let profile = this.props.projects.currentProfiles.filter((value, index) => {
        return value.UUID == this.props.ownProps.params.id;
      })
      if (profile.length > 0) {
        this.props.ProfileActions.setCurrentProfile(profile[0])
      }
    }
  }

  static propTypes = {
    classes: PropTypes.object.isRequired,
  }

  render () {
    const {classes, boards} = this.props

    return (
      <Grid className={classes.root}>
        <Grid item xs={12}>
          <Grid container className={classes.list} justify="center">
            {boards.currentBoards.map((item, i) => (
              <Grid key={i} item>
                <BoardCard board={item}/>
              </Grid>
            ))}
          </Grid>
        </Grid>
      </Grid>
    )
  }
}

const styles = theme => ({
  root: {
    minHeight: '100vh',
    backgroundColor: '#2B2D42',
  },
  paper: {
    height: 140,
    width: 100,
  },
  control: {
    padding: theme.spacing.unit * 2,
  },
  list: {
    paddingTop: 20,
    paddingLeft: 80,
    paddingRight: 80,
  }
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
    ProfileActions: bindActionCreators(ProfileActions, dispatch)
  }
}

export default withStyles(styles)(
  connect(mapStateToProps, mapDispatchToProps)(ViewUserProfile)
)
