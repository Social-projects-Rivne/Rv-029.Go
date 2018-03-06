import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Card, { CardActions, CardContent } from 'material-ui/Card';
import Paper from 'material-ui/Paper';
import Button from 'material-ui/Button';
import TextField from 'material-ui/TextField'
import Dialog, {
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
} from 'material-ui/Dialog';
import Typography from 'material-ui/Typography';
import { Link, browserHistory } from 'react-router'


class ProjectCard extends Component {


    state = {
        AddUserOpen: false,
    }

    handleClose = () => {
        this.setState({ AddUserOpen: false })
    }

    handleClickOpenAddUser = () => {
        this.setState({ AddUserOpen: true })
    };

  render() {
      const { classes, board } = this.props;
    return (
        <div>
            <Card className={classes.card}>
                <CardContent>
                    <Typography variant="headline" component="h2">{board.name}</Typography>
                    <Typography className={classes.title}>{board.description}</Typography>
                </CardContent>
                <CardActions>
                    <Link
                        to={`board/${board.id}`}
                        className={classes.link}>
                        <Button
                            size="small"
                            color={'secondary'}
                            >
                            View
                        </Button>
                    </Link>
                    <Button variant="raised" color="primary" className={classes.button} onClick={this.handleClickOpenAddUser}>
                        Add user
                    </Button>

                </CardActions>
            </Card>

            {/*/!* #################### MODAL ISSUE #################### *!/*/}
            <Dialog
                open={this.state.AddUserOpen}
                onClose={this.handleClose}
                aria-labelledby="form-dialog-title" >
                <DialogTitle id="form-dialog-title">Add user</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        Please, fill required fields.
                    </DialogContentText>

                </DialogContent>
                <DialogActions>
                    <Button onClick={this.handleClose} color="primary">
                        Cancel
                    </Button>
                </DialogActions>
            </Dialog>
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

ProjectCard.propTypes = {
  board: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ProjectCard);
