import React from 'react'
import PropTypes from 'prop-types'
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
import Chip from 'material-ui/Chip';

const IssueCard = ({ title, desc, status, estimate }) => {
  return(
    <ExpansionPanel>
      <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
        <Grid
          container
          alignItems={'center'}>
          <Grid style={{ marginRight: '1em' }}>
            <Chip label={ status } />
          </Grid>
          <Grid>
            <Typography>{ title }</Typography>
          </Grid>
        </Grid>
      </ExpansionPanelSummary>
      <ExpansionPanelDetails>
        <Grid container >
          <Grid item xs={12}>
            <Typography type={'title'}>{ `Estimate: ${estimate}` }</Typography>
            <Typography> { desc } </Typography>
          </Grid>
        </Grid>
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


IssueCard.propTypes = {
  title: PropTypes.string.isRequired,
  desc: PropTypes.string.isRequired,
  status: PropTypes.string.isRequired,
  estimate: PropTypes.number.isRequired
}

export default IssueCard
