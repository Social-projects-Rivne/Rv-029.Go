import React, { Component } from 'react';
import Users from '../components/scrumPocker/Users'
import Issue from '../components/scrumPocker/Issue'
import Estimation from '../components/scrumPocker/Estimation'
import {withStyles} from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import * as projectsActions from "../actions/ProjectsActions";
import * as defaultPageActions from "../actions/DefaultPageActions";
import * as sprintsActions from "../actions/SprintsActions";
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import Card, { CardActions, CardContent } from 'material-ui/Card';

class EstimationRoom extends Component {

  render() {
    const { classes } = this.props

    return (
      <div className={classes.root}>
        <Grid
          container
          spacing={40}
          justify={'center'}
          className={classes.container} >

          <Grid item xs={4}>
            <Issue issueID={this.props.ownProps.params.id} />
            <Users />
          </Grid>

          <Grid item xs={8}>
            <Card>
              <Estimation/>
            </Card>
          </Grid>

        </Grid>
      </div>
    )
  }
}

const styles = {
  container: {
    padding: '2em 6em',
  },
  root: {
    width: '100%',
    overflow: 'hidden'
  },
  pos: {
    marginBottom: 12,
  },
}

const mapStateToProps = (state, ownProps) => {
  return {
    sprints: state.sprints,
    projects: state.projects,
    defaultPage: state.defaultPage,
    ownProps
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    sprintsActions: bindActionCreators(sprintsActions, dispatch),
    projectsActions: bindActionCreators(projectsActions, dispatch),
    defaultPageActions: bindActionCreators(defaultPageActions, dispatch)
  }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(EstimationRoom)
)
