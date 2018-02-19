import React, { Component } from 'react'
import Card, { CardHeader, CardActions, CardContent } from 'material-ui/Card'
import Button from 'material-ui/Button'
import { Link, browserHistory } from 'react-router'
import Typography from 'material-ui/Typography'
import { withStyles } from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import Chip from 'material-ui/Chip';
import IconButton from 'material-ui/IconButton';
import DeleteIcon from 'material-ui-icons/Delete';
import Icon from 'material-ui/Icon';

class SprintCard extends Component {

  state = {
    open: false,
    anchorEl: null
  }

  handleClick = () => {
    this.setState({ open: true });
  };

  handleClose = () => {
    this.setState({ open: false });
  };

  render() {
    const { classes } = this.props;
    const { open } = this.state;
    const { anchorEl } = this.state;

    return (

      <Card className={classes.root}>
        <CardHeader
          avatar={
            <Chip label={this.props.status} />
          }
          action={
            <Grid item>
              {/* FIXME: horizontal scroll cause of this btn WTF? */}
              <IconButton>
                <Icon>edit_icon</Icon>
              </IconButton>
              <IconButton>
                <DeleteIcon />
              </IconButton>
            </Grid>
          }
          title={this.props.title}
          subheader={this.props.date} />
        <CardContent>
          <Typography>{this.props.desc}</Typography>
        </CardContent>
        <CardActions>
          <Link
            // to="view_board"
            className={this.props.classes.link}>
            <Button
              size="small"
              color={'secondary'}>
              View
            </Button>
          </Link>
        </CardActions>
      </Card>
    )
  }
}

const styles = {
  root: {
    marginBottom: '1em'
  },
  link: {
    textDecoration: 'none'
  },
}

export default withStyles(styles)(SprintCard)

