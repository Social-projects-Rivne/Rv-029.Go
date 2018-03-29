import React, { Component } from 'react';
import Typography from 'material-ui/Typography';
import List, { ListItem, ListItemIcon, ListItemText } from 'material-ui/List';
import Card, { CardContent } from 'material-ui/Card';
import Avatar from 'material-ui/Avatar';
import ImageIcon from 'material-ui-icons/Image';
import RadioButtonUnchecked from 'material-ui-icons/RadioButtonUnchecked';
import RadioButtonChecked from 'material-ui-icons/RadioButtonChecked';

export default class Users extends Component {
  render() {
      console.log("sdfsdf"+this.props.users)
    return (
      <Card>
        <CardContent>
          <Typography type='subheading' component="h2">
            Active Users
          </Typography>

          <List>
            {(this.props.users) ?
              (this.props.users.map((item, i) => (
                <ListItem key={i}>
                  <Avatar>
                      <ImageIcon />
                  </Avatar>
                  <ListItemText primary={item.Email} secondary="ready"/>
                  <ListItemIcon>
                      <RadioButtonChecked />
                  </ListItemIcon>
                </ListItem>
              ))) : (<ListItem>Room is empty</ListItem>)}
          </List>

        </CardContent>
      </Card>
    )
  }
}

