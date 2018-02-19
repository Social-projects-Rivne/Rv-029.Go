import React, { Component } from 'react'
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Card, { CardActions, CardContent } from 'material-ui/Card';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';
import { Link, browserHistory } from 'react-router'
import * as boardsActions from "../actions/BoardsActions";
import * as defaultPageActions from "../actions/DefaultPageActions";
import * as projectsActions from "../actions/ProjectsActions";
import {bindActionCreators} from "redux";
import {connect} from "react-redux";

class ProjectCard extends Component {

    static propTypes = {
        project: PropTypes.object.isRequired,
        classes: PropTypes.object.isRequired,
    }

    viewProject = (projectID) => {
        const results = this.props.projects.currentProjects.filter((item) => { return item.UUID === projectID });
        if (results.length > 0) {
            this.props.projectsActions.setCurrentProject(results[0])
        }
        browserHistory.push('/project/' + projectID)
    }

    render = () => {
        const { classes, project } = this.props;

        return (
            <div>
                <Card className={classes.card}>
                    <CardContent>
                        <Typography variant="headline" component="h2">{project.Name}</Typography>
                    </CardContent>
                    <CardActions>
                        <Button
                            onClick={() => { this.viewProject(project.UUID) }}
                            size="small"
                            color={'secondary'}>
                            View
                        </Button>
                    </CardActions>
                </Card>
            </div>
        )
    }
}


const styles = theme => ({
    card: {
        minWidth: 275,
    },
    bullet: {
        display: 'inline-block',
        margin: '0 2px',
        transform: 'scale(0.8)',
    },
    title: {
        marginBottom: 16,
        fontSize: 14,
        color: theme.palette.text.secondary,
    },
    pos: {
        marginBottom: 12,
        color: theme.palette.text.secondary,
    },
    link: {
        textDecoration: 'none'
    }
});

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
    connect(mapStateToProps, mapDispatchToProps)(ProjectCard)
)
