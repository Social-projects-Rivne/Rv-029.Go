import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Card, { CardActions, CardContent } from 'material-ui/Card';
import Paper from 'material-ui/Paper';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';
import { Link, browserHistory } from 'react-router'

function ProjectCard(props) {
    const { classes, project } = props;

    return (
        <div>
            <Card className={classes.card}>
                <CardContent>
                    <Typography className={classes.title}>{project.user_name}</Typography>
                    <Typography variant="headline" component="h2">{project.name}</Typography>
                    <Typography component="p">{project.desc}</Typography>
                </CardContent>
                <CardActions>
                    <Link
                          to={"project/" + project.id}
                          className={classes.link}>
                        <Button
                          size="small"
                          color={'secondary'}>
                            View
                        </Button>
                    </Link>
                </CardActions>
            </Card>
        </div>
    )
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

ProjectCard.propTypes = {
    project: PropTypes.object.isRequired,
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ProjectCard);