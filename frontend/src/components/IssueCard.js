import React from 'react'
import ExpansionPanel, {
  ExpansionPanelSummary,
  ExpansionPanelDetails,
} from 'material-ui/ExpansionPanel'
import ExpandMoreIcon from 'material-ui-icons/ExpandMore'
import IconButton from 'material-ui/IconButton';
import DeleteIcon from 'material-ui-icons/Delete';
import Icon from 'material-ui/Icon';
import Typography from 'material-ui/Typography'
import Grid from 'material-ui/Grid'

const IssueCard = () => {
  return(
    <ExpansionPanel>
      <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
        <Typography>Issue 1</Typography>
      </ExpansionPanelSummary>
      <ExpansionPanelDetails>
        <Typography>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse malesuada lacus ex,
          sit amet blandit leo lobortis eget.
        </Typography>
      </ExpansionPanelDetails>
      <ExpansionPanelDetails>
        <Grid
          container
          justify={'flex-end'}>
          <Grid item>
            <IconButton aria-label="Delete">
              <DeleteIcon />
            </IconButton>
            <IconButton>
              <Icon>edit_icon</Icon>
            </IconButton>
          </Grid>
        </Grid>
      </ExpansionPanelDetails>
    </ExpansionPanel>
  )
}

export default IssueCard
