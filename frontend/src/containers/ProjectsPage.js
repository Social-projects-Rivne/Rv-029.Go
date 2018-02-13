import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {withStyles} from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import ProjectCard from '../components/ProjectCard';
import * as topBarActions from "../actions/TopBarActions";
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import {API_URL} from "../constants/global";
import axios from "axios";

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

class ProjectsPage extends Component {

    constructor(props) {
        super(props)
        if (props.topBar.pageTitle !== pageTitle) {
            this.props.action.changePageTitle(pageTitle)
        }
    }

    componentDidMount() {
        axios.get(API_URL + 'project/list')
            .then((response) => {
                this.props.action.setProjects(response.data)
            })
            .catch((error) => {
                if (error.response && error.response.data.Message) {
                    this.props.action.setErrorMessage(error.response.data.Message)
                } else {
                    this.props.action.setErrorMessage("Server error occured")
                }
            });
    }

    static propTypes = {
        classes: PropTypes.object.isRequired,
    }

    render () {
        const {classes, topBar} = this.props

        return (
            <Grid container className={classes.root}>
                <Grid item xs={12}>
                    <Grid container className={classes.list} justify="center" spacing={24}>
                        {topBar.currentProjects.map((value, index) => (
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
