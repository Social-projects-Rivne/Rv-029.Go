import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Stepper, { Step, StepLabel, StepContent } from 'material-ui/Stepper';
import Button from 'material-ui/Button';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import IconButton from 'material-ui/IconButton';
import Filter1 from 'material-ui-icons/Filter1';
import Filter3 from 'material-ui-icons/Filter3';
import Filter5 from 'material-ui-icons/Filter5';
import Filter8 from 'material-ui-icons/Filter8';

function getSteps() {
  return [
    'You are in estimation room',
    'Join to estimation',
    'Estimate'
  ]
}

class VerticalLinearStepper extends React.Component {
  state = {
    activeStep: 0,
  };

  handleNext = () => {
    this.setState({
      activeStep: this.state.activeStep + 1,
    });
  };

  handleBack = () => {
    this.setState({
      activeStep: this.state.activeStep - 1,
    });
  };

  handleReset = () => {
    this.setState({
      activeStep: 0,
    });
  };

  getStepContent = step => {
    const { classes } = this.props;

    switch (step) {
      case 0:
        return ( <Typography>For each ad campaign that you create, you can control how much you're willing to spend on clicks and conversions, which networks and geographical locations you want your ads to show on, and more. </Typography> )
      case 1:
        return (
          <div>
            <Typography>An ad group contains one or more ads which target a shared set of keywords.</Typography>
              <Button
                className={classes.button}
                raised
                color={'secondary'}>
                Join
              </Button>
          </div>
        )
      case 2:
        return (
          <div>
            <Typography>Try out different ad text to see what brings in the most customers, and learn how to enhance your ads using features like ad extensions. If you run into any problems with your ads, find out how to tell if they're running and how to resolve approval issues.</Typography>
            {/*<Button color="primary" className={classes.estimateButton}>*/}
              {/*1*/}
            {/*</Button>*/}
            {/*<Button color="primary" className={classes.estimateButton}>*/}
              {/*3*/}
            {/*</Button>*/}
            {/*<Button color="primary" className={classes.estimateButton}>*/}
              {/*5*/}
            {/*</Button>*/}
            {/*<Button color="primary" className={classes.estimateButton}>*/}
              {/*8*/}
            {/*</Button>*/}

            <IconButton>
              <Filter1 />
            </IconButton>
            <IconButton>
              <Filter3 />
            </IconButton>
            <IconButton>
              <Filter5 />
            </IconButton>
            <IconButton>
              <Filter8 />
            </IconButton>
          </div>
        )
      default:
        return 'Unknown step';
    }
  }

  render() {
    const { classes } = this.props;
    const { activeStep } = this.state;
    const steps = getSteps();

    return (
      <div className={classes.root}>
        <Stepper activeStep={activeStep} orientation="vertical">
          {steps.map((label, index) => {
            return (
              <Step key={label}>
                <StepLabel>{label}</StepLabel>
                <StepContent>
                  <div className={classes.actionsContainer}>

                    {this.getStepContent(index)}

                    <div>
                      <Button
                        disabled={activeStep === 0}
                        onClick={this.handleBack}
                        className={classes.button} >
                        Back
                      </Button>
                      <Button
                        raised
                        color="primary"
                        onClick={this.handleNext}
                        className={classes.button} >
                        {activeStep === steps.length - 1 ? 'Finish' : 'Next'}
                      </Button>
                    </div>
                  </div>
                </StepContent>
              </Step>
            );
          })}
        </Stepper>
        {activeStep === steps.length && (
          <Paper square elevation={0} className={classes.resetContainer}>
            <Typography>All steps completed - you&quot;re finished</Typography>
            <Button onClick={this.handleReset} className={classes.button}>
              Reset
            </Button>
          </Paper>
        )}
      </div>
    );
  }
}

VerticalLinearStepper.propTypes = {
  classes: PropTypes.object,
};

const styles = theme => ({
  root: {
    width: '90%',
  },
  button: {
    marginTop: theme.spacing.unit,
    marginRight: theme.spacing.unit,
  },
  estimateButton: {
    marginLeft: 0,
    marginRight: 0,
    borderBottom: '1px solid #c4c4c4',
    borderRadius: '0',
    marginTop: theme.spacing.unit
  },
  actionsContainer: {
    marginBottom: theme.spacing.unit * 2,
  },
  resetContainer: {
    padding: theme.spacing.unit * 3,
  },
});

export default withStyles(styles)(VerticalLinearStepper);