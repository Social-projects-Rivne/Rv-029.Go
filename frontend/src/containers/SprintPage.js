import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {withStyles} from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import Paper from 'material-ui/Paper';
import IssueCard from '../components/IssueCard';
import * as projectsActions from "../actions/ProjectsActions";
import * as defaultPageActions from "../actions/DefaultPageActions";
import * as sprintsActions from "../actions/SprintsActions";
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import {API_URL} from "../constants/global";
import messages from "../services/messages";
import axios from "axios";

const pageTitle = "Active Sprint";

const styles = theme => ({
    root: {
        // flexGrow: 1,
        minHeight: '100vh',
        backgroundColor: '#2B2D42',
    },
    paper: {
        marginLeft: 20,
    },
    columnTitle: {
        paddingTop: 20,
        textAlign: 'center',
        fontWeight: 'bold',
        fontSize: 18,
    },
    status: {
        padding: 40
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

class SprintPage extends Component {

    constructor(props) {
        super(props)
        if (props.defaultPage.pageTitle !== pageTitle) {
            this.props.defaultPageActions.changePageTitle(pageTitle)
        }
    }

    componentWillMount() {
        axios.get(API_URL + `project/board/sprint/show/${this.props.ownProps.params.id}`)
            .then((response) => {
                this.props.sprintsActions.setActiveSprint(response.data.Data)
            })
            .catch((error) => {
                if (error.response && error.response.data.Message) {
                    messages(error.response.data.Message)
                } else {
                    messages("Server error occured")
                }
            });

        axios.get(API_URL + `project/board/sprint/${this.props.ownProps.params.id}/issue/list`)
            .then((response) => {
                this.props.sprintsActions.setSprintIssues(response.data.Data)
            })
            .catch((error) => {
                if (error.response && error.response.data.Message) {
                    messages(error.response.data.Message)
                } else {
                    messages("Server error occured")
                }
            });
    }

    groupByStatus = () => {
        console.log(this.props.sprints);
        let i = 0, val, index, values = [], result = [];

        for (; i < this.props.sprints.issues.length; i++) {
            val = this.props.sprints.issues[i]['Status'];
            index = values.indexOf(val);
            if (index > -1)
                result[index].push(this.props.sprints.issues[i]);
            else {
                values.push(val);
                result.push([this.props.sprints.issues[i]]);
            }
        }

        return result;
    };

    static propTypes = {
        classes: PropTypes.object.isRequired,
    }

    render () {
        const {classes, projects, } = this.props
        const issues = this.groupByStatus();
        return (
            <Grid className={classes.root}>
                <Grid item xs={12} container>
                    <Grid item xs={4} className={classes.status}>
                        <Paper className={classes.paper} elevation={4}>
                            <h5 className={classes.columnTitle}>TODO:</h5>
                            <hr/>
                            <Grid container className={classes.list}>
                                {issues.map((value, index) => (
                                    <IssueCard key={value.id}
                                        assigned={value.user_id}
                                        status={value.Status}
                                        description={value.desc}
                                    />
                                ))}
                            </Grid>
                        </Paper>
                    </Grid>
                    <Grid item xs={4} className={classes.status}>
                        <Paper className={classes.paper} elevation={4}>
                            <h5 className={classes.columnTitle}>In Progress:</h5>
                            <hr/>
                            <Grid container className={classes.list}>
                            </Grid>
                        </Paper>
                    </Grid>
                    <Grid item xs={4} className={classes.status}>
                        <Paper className={classes.paper} elevation={4}>
                            <h5 className={classes.columnTitle}>Done:</h5>
                            <hr/>
                            <Grid container className={classes.list}>
                            </Grid>
                        </Paper>
                    </Grid>
                </Grid>
            </Grid>
        )
    }
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
    connect(mapStateToProps, mapDispatchToProps)(SprintPage)
)
