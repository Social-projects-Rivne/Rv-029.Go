import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {withStyles} from 'material-ui/styles';
import { Link, browserHistory } from 'react-router'
import Grid from 'material-ui/Grid';
import * as projectsActions from "../actions/ProjectsActions";
import * as boardsActions from "../actions/BoardsActions";
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

const pageTitle = "Add Board"

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

class CreateBoardPage extends Component {

    getInitialState = () => {
        return {
            inputBoardName: '',
            inputBoardNameValid: false,
            inputBoardDesc: '',
            inputBoardDescValid: false,
        }
    }

    constructor(props) {
        super(props)

        this.state = {
            inputBoardName: '',
            inputBoardNameValid: true,
            inputBoardDesc: '',
            inputBoardDescValid: true,
        };

        if (props.defaultPage.pageTitle !== pageTitle) {
            this.props.defaultPageActions.changePageTitle(pageTitle)
        }
    }

    static propTypes = {
        classes: PropTypes.object.isRequired,
    }

    componentDidMount() {
        if (this.props.projects.currentProject == null) {
            axios.get(API_URL + `project/show/${this.props.ownProps.params.id}`)
                .then((response) => {
                    this.props.projectsActions.setCurrentProject(response.data.Data)
                    this.props.defaultPageActions.changePageTitle("Project " + this.props.projects.currentProject.Name)
                })
                .catch((error) => {
                    if (error.response && error.response.data.Message) {
                        messages(error.response.data.Message)
                    } else {
                        messages("Server error occured")
                    }
                })
        }
    }

    createBoard = (e) => {
        e.preventDefault();
        if (this.validateBoardName() && this.validateBoardDesc()) {
            axios.post(API_URL + 'project/'+ this.props.projects.currentProject.UUID +'/board/create', {
                name: this.state.inputBoardName,
                description: this.state.inputBoardDesc,
            })
                .then((response) => {
                    if (response.data.Status) {
                        messages(response.data.Message, response.data.Status)
                        browserHistory.push('/project/' + this.props.projects.currentProject.UUID)
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

    validateBoardName = () => {
        if (this.state.inputBoardNameValid && this.state.inputBoardName.length < 3) {
            this.setState({
                inputBoardNameValid: false,
            });

            return false;
        }

        return this.state.inputBoardNameValid
    }

    validateBoardDesc = () => {
        if (this.state.inputBoardDescValid && this.state.inputBoardDesc.length < 3) {
            this.setState({
                inputBoardDescValid: false,
            });

            return false;
        }

        return this.state.inputBoardDescValid
    }

    changeBoardNameInput = (e) => {
        this.setState({
            inputBoardName: e.target.value,
            inputBoardNameValid: (e.target.value.length >= 3),
        });
    }

    changeBoardDescInput = (e) => {
        this.setState({
            inputBoardDesc: e.target.value,
            inputBoardDescValid: (e.target.value.length >= 3),
        });
    }

    render () {

        const {classes, boards } = this.props
        const {currentProject} = this.props.projects

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
                        Create Board
                    </Typography>

                    <FormInput type="text" value={this.state.inputBoardName} name="Board Name" isValid={this.state.inputBoardNameValid} onUserInput={this.changeBoardNameInput}/>
                    <FormInput type="text" value={this.state.inputBoardDesc} name="Board Description" isValid={this.state.inputBoardDescValid} onUserInput={this.changeBoardDescInput}/>

                    <Grid
                        container
                        alignItems={'center'}
                        justify={'space-around'}
                        className={classes.buttons}>
                        <Button
                            type='submit'
                            color='primary'
                            onClick={this.createBoard}>
                            Create
                        </Button>
                        {(currentProject !== null) ? (
                            <Link to={'/project/' + currentProject.UUID }
                                  className={classes.link}>
                                <Button
                                    color={'secondary'}>
                                    Back
                                </Button>
                            </Link>
                        ) : ('')}
                    </Grid>

                </Paper>
            </Grid>
        )
    }
}

const mapStateToProps = (state, ownProps) => {
    return {
        projects: state.projects,
        boards: state.boards,
        defaultPage: state.defaultPage,
        ownProps
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        boards: bindActionCreators(boardsActions, dispatch),
        projectsActions: bindActionCreators(projectsActions, dispatch),
        defaultPageActions: bindActionCreators(defaultPageActions, dispatch)
    }
}

export default withStyles(styles)(
    connect(mapStateToProps, mapDispatchToProps)(CreateBoardPage)
)
