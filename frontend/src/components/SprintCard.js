import React, { Component } from 'react'
import Card, { CardHeader, CardActions, CardContent } from 'material-ui/Card'
import Button from 'material-ui/Button'
import { Link, browserHistory } from 'react-router'
import Typography from 'material-ui/Typography'
import { withStyles } from 'material-ui/styles'
import { Manager, Target, Popper } from 'react-popper'
import ClickAwayListener from 'material-ui/utils/ClickAwayListener'
import Grow from 'material-ui/transitions/Grow'
import Paper from 'material-ui/Paper';
import { MenuItem, MenuList } from 'material-ui/Menu';
import classNames from 'classnames';
import IconButton from 'material-ui/IconButton';
import MoreVertIcon from 'material-ui-icons/MoreVert';

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
          action={
            <Manager className={this.props.classes.button}>
              <Target>
                <IconButton
                  aria-label="More"
                  aria-owns={anchorEl ? 'long-menu' : null}
                  aria-haspopup="true"
                  onClick={this.handleClick} >
                  <MoreVertIcon />
                </IconButton>
              </Target>
              <Popper
                placement="bottom-start"
                eventsEnabled={open}
                className={classNames({ [classes.popperClose]: !open })} >
                <ClickAwayListener onClickAway={this.handleClose}>
                  <Grow in={open} id="menu-list" style={{ transformOrigin: '0 0 0' }}>
                    <Paper>
                      <MenuList role="menu">
                        <MenuItem onClick={this.handleClose}>Edit</MenuItem>
                        <MenuItem onClick={this.handleClose}>Remove</MenuItem>
                      </MenuList>
                    </Paper>
                  </Grow>
                </ClickAwayListener>
              </Popper>
            </Manager>
          }
          title={this.props.title}
          subheader={this.props.date}
        />
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
