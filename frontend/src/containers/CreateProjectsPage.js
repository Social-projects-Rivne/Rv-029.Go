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
import {API_URL} from "../constants/global";
import messages from "../services/messages";
import axios from "axios";
import Button from 'material-ui/Button'
import FormInput from "../components/FormInput";
import auth from "../services/auth";

const pageTitle = "Add Project"

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

    getInitialState = () => {
        return {
            inputProjectName: '',
            inputProjectNameValid: false,
        }
    }

    constructor(props) {
        super(props)

        this.state = {
            inputProjectName: '',
            inputProjectNameValid: true,
        };

        if (props.defaultPage.pageTitle !== pageTitle) {
            this.props.defaultPageActions.changePageTitle(pageTitle)
        }
    }

    static propTypes = {
        classes: PropTypes.object.isRequired,
    }

    createProject = (e) => {
        e.preventDefault()
        if (this.validateProjectName()) {
            axios.post(API_URL + 'project/create', {
                name: this.state.inputProjectName,
            })
                .then((response) => {
                    if (response.data.Status) {
                        messages(response.data.Message, response.data.Status)
                        browserHistory.push('/projects')
                    }
                })
                .catch((error) => {
                    if (error.response && error.response.data.Message) {
                        messages(error.response.data.Message, false)
                    } else {
                        messages("Server error occured", false)
                    }
                })
        }
    }

    validateProjectName = () => {
        if (this.state.inputProjectNameValid && this.state.inputProjectName.length < 3) {
            this.setState({
                inputProjectNameValid: false,
            });

            return false;
        }

        return this.state.inputProjectNameValid
    }

    changeProjectNameInput = (e) => {
        this.setState({
            inputProjectName: e.target.value,
            inputProjectNameValid: (e.target.value.length >= 3)
        });
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

                    <FormInput type="text" value={this.state.inputProjectName} name="Project Name" isValid={this.state.inputProjectNameValid} onUserInput={this.changeProjectNameInput}/>

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
