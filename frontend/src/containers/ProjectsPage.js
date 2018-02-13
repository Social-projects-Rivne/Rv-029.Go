import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import ProjectCard from '../components/ProjectCard';
import * as topBarActions from "../actions/TopBarActions";
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import {API_URL} from "../constants/global";
import axios from "axios/index";

const pageTitle = "Projects"

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

const ProjectsPage = ({classes, topBar, action, ownProps, ...decorator}) => {

    if (topBar.pageTitle !== pageTitle) {
        action.changePageTitle(pageTitle)
    }
    let projects = [];

    axios.get(API_URL + 'project/list')
        .then((response) => {
            projects = response
        })
        .catch((response) => {
            action.setStatus(response.response.data.status)
            if (response.response.data.Message) {
                action.setErrorMessage(response.response.data.Message)
            } else {
                action.setErrorMessage("Server error occured")
            }
        });


    return (
        <Grid container className={classes.root}>
            <Grid item xs={12}>
                <Grid container className={classes.list} justify="center" spacing={24}>
                    {projects.map((value, index) => (
                        <Grid key={index} item>
                            <ProjectCard project={value}/>
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

const mapStateToProps = (state, ownProps) => {
    return {
        topBar: state.topBar,
        ownProps
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        action: bindActionCreators(topBarActions, dispatch)
    }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(ProjectsPage)
)
