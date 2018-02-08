import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import { FormLabel, FormControlLabel } from 'material-ui/Form';
import Radio, { RadioGroup } from 'material-ui/Radio';
import Paper from 'material-ui/Paper';
import ProjectCard from '../components/ProjectCard';

const styles = theme => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        height: 140,
        width: 100,
    },
    control: {
        padding: theme.spacing.unit * 2,
    },
});

const ProjectsPage = (props) => {
  const projects = [{
      id: 1,
      title: "Test Project",
      author: "SomeBody",
      desc: "Some Project Description"
  }];

  return (
      <Paper>
          <Grid container className={props.classes.root}>
              <Grid item xs={12}>
                  <Grid container justify="center" spacing={24}>
                      {projects.map(value => (
                          <ProjectCard project={value} />
                      ))}
                  </Grid>
              </Grid>
          </Grid>
      </Paper>
  )
}

ProjectsPage.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ProjectsPage);