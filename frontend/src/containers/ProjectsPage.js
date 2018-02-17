import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {withStyles} from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import ProjectCard from '../components/ProjectCard';
import * as projectsActions from "../actions/ProjectsActions";
import * as defaultPageActions from "../actions/DefaultPageActions";
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import {API_URL} from "../constants/global";
import messages from "../services/messages";
import axios from "axios";

const pageTitle = "Projects"

const styles = theme => ({
    root: {
        // flexGrow: 1,
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

class ProjectsPage extends Component {

    constructor(props) {
        super(props)
        if (props.defaultPage.pageTitle !== pageTitle) {
            this.props.defaultPageActions.changePageTitle(pageTitle)
        }
    }

    componentDidMount() {
        axios.get(API_URL + 'project/list')
            .then((response) => {
                this.props.projectsActions.setProjects(response.data)
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
        const {classes, projects, } = this.props


        return (
            <Grid className={classes.root}>
                <Grid item xs={12}>
                    <Grid container className={classes.list}>
                        {projects.currentProjects.map((value, index) => (
                            <Grid key={index} item>
                               <ProjectCard project={value}/>
                            </Grid>
                         ))}
                    </Grid>
                </Grid>
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
    connect(mapStateToProps, mapDispatchToProps)(ProjectsPage)
)
