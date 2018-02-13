import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import ProjectCard from '../components/ProjectCard';

const styles = theme => ({
    root: {
        flexGrow: 1,
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
        paddingTop: 20
    }
});

const ProjectsPage = (props) => {

  const projects = [{
      id: 1,
      title: "Test Project 1",
      author: "SomeBody",
      desc: "Some Project Description"
  },{
      id: 2,
      title: "Test Project 2",
      author: "SomeBody",
      desc: "Some Project Description"
  },{
      id: 3,
      title: "Test Project 3",
      author: "SomeBody",
      desc: "Some Project Description"
  },{
      id: 4,
      title: "Test Project 4",
      author: "SomeBody",
      desc: "Some Project Description"
  },{
      id: 5,
      title: "Test Project 5",
      author: "SomeBody",
      desc: "Some Project Description"
  }];

  return (
      <Grid container className={props.classes.root} >
          <Grid item xs={12}>
              <Grid container className={props.classes.list} justify="center" spacing={24}>
                  {projects.map((value, index) => (
                      <Grid key={index} item>
                          <ProjectCard project={value} />
                      </Grid>
                  ))}
              </Grid>
          </Grid>
      </Grid>
  )
}

ProjectsPage.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ProjectsPage);