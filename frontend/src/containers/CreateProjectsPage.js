import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {withStyles} from 'material-ui/styles';
import { Link, browserHistory } from 'react-router'
import Grid from 'material-ui/Grid';
import * as projectsActions from "../actions/ProjectsActions";
import * as defaultPageActions from "../actions/DefaultPageActions";
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import Paper from 'material-ui/Paper'
import Typography from 'material-ui/Typography'
import Input from 'material-ui/Input'
import {API_URL} from "../constants/global";
import messages from "../services/messages";
import axios from "axios";
import Button from 'material-ui/Button'
import FormInput from "../components/FormInput";

const pageTitle = "Projects"

const styles = theme => ({
    root: {
        minHeight: '100vh',
        backgroundColor: '#2B2D42',
    },
    control: {
        padding: theme.spacing.unit * 2,
    },
    list: {
        paddingTop: 20
    },
    paper: {
        padding: '4em 3em',
    }
});

class CreateProjectsPage extends Component {

    constructor(props) {
        super(props)
        if (props.defaultPage.pageTitle !== pageTitle) {
            this.props.defaultPageActions.changePageTitle(pageTitle)
        }
    }

    static propTypes = {
        classes: PropTypes.object.isRequired,
    }

    //TODO:
    createProject = () => {
        return true
    }


    //TODO:
    validateProjectName = (name) => {
        return true
    }

    //TODO:
    validateProjectNameInput = (name) => {

    }

    render () {
        const {classes, projects } = this.props

        return (
            <Grid container
                  className={classes.root}
                  spacing={0}
                  alignItems={'center'}
                  justify={'center'} >
                <Paper
                    className={classes.paper}
                    elevation={8}
                    component='form'>

                    <Typography
                        type='headline'
                        component='h3'>
                        Create Project
                    </Typography>

                    <FormInput type="text" name="Project Name" isValid={this.validateProjectName()} onUserInput={this.validateProjectNameInput}/>

                    <Grid
                        container
                        alignItems={'center'}
                        justify={'space-around'}
                        className={classes.buttons}>
                        <Button
                            type='submit'
                            color='primary'
                            onClick={this.createProject}>
                            Create
                        </Button>
                        <Link to={'/projects'}
                              className={classes.link}>
                            <Button
                                color={'secondary'}>
                                Back
                            </Button>
                        </Link>
                    </Grid>

                </Paper>
            </Grid>
        )
    }
}

const mapStateToProps = (state, ownProps) => {
    return {
        projects: state.projects,
        defaultPage: state.defaultPage,
        ownProps
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        projectsActions: bindActionCreators(projectsActions, dispatch),
        defaultPageActions: bindActionCreators(defaultPageActions, dispatch)
    }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(CreateProjectsPage)
)
