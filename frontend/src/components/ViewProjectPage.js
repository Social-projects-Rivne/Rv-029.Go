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


class ViewProjectPage extends Component {

  componentWillMount() {
    if (this.props.projects.currentProject == null) {
        axios.get(API_URL + `project/show/${this.props.ownProps.params.id}`)
            .then((response) => {
                this.props.projectsActions.setCurrentProject(response.data)
                this.props.defaultPageActions.changePageTitle("Project " + this.props.projects.currentProject.Name)
            })
            .catch((error) => {
                if (error.response && error.response.data.Message) {
                    messages(error.response.data.Message)
                } else {
                    messages("Server error occured")
                }
            })
    } else {
        this.props.defaultPageActions.changePageTitle("Project " + this.props.projects.currentProject.Name)
    }

    axios.get(API_URL + `project/${this.props.ownProps.params.id}/board/list`)
    .then((response) => {
      this.props.boardsActions.setBoards(response.data)
    })
    .catch((error) => {
      if (error.response && error.response.data.Message) {
        messages(error.response.data.Message)
      } else {
        messages("Server error occured")
      }
    })
  }

  componentDidMount() {
    if (this.props.projects.currentProjects.length > 0) {
      let project = this.props.projects.currentProjects.filter((value, index) => {
        return value.UUID == this.props.ownProps.params.id;
      })
      if (project.length > 0) {
        this.props.projectsActions.setCurrentProject(project[0])
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
    boards: state.boards,
    projects: state.projects,
    defaultPage: state.defaultPage,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    boardsActions: bindActionCreators(boardsActions, dispatch),
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch),
    projectsActions: bindActionCreators(projectsActions, dispatch)
  }
}

export default withStyles(styles)(
  connect(mapStateToProps, mapDispatchToProps)(ViewProjectPage)
)
