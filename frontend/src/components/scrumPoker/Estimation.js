import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Stepper, { Step, StepLabel, StepContent } from 'material-ui/Stepper';
import Card from 'material-ui/Card';
import Button from 'material-ui/Button';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';

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
    estimation: null
  };

  componentWillReceiveProps(nextProps) {

    if (this.state.activeStep !== nextProps.activeStep) {
      this.setState({
        activeStep: nextProps.activeStep
      })
    }

  }

  handleNext = () => {
    this.props.actions.increaseStep()
  };

  handleBack = () => {
    this.props.actions.decreaseStep()
  };

  join = () => {
    this.props.registerClient()
  }

  handleActiveButton = index => () => {
    this.setState({ estimation: index })
  }

  sendEstimate = (est) => () => {
    this.props.sendEstimate(est)
  }

  getStepContent = step => {
    const { classes } = this.props;
    const { activeStep } = this.state;
    const steps = getSteps();

    switch (step) {
      case 0:
        return (
          <div>

            <Typography>
              For each ad campaign that you create, you can control how much you're willing to spend on clicks and conversions, which networks and geographical locations you want your ads to show on, and more.
            </Typography>

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
                Next
              </Button>
            </div>

          </div>
        )
      case 1:
        return (
          <div>
            <Typography>
              An ad group contains one or more ads which target a shared set of keywords.
            </Typography>

            <div>
              <Button
                disabled={activeStep === 0}
                onClick={this.handleBack}
                className={classes.button} >
                Back
              </Button>

              <Button
                raised
                className={classes.button}
                onClick={this.join}
                color={'secondary'}>
                Join
              </Button>
            </div>
          </div>
        )
      case 2:
        let estimation = [1, 3, 5 ,8]

        return (
          <div>
            <Typography>
              Try out different ad text to see what brings in the most customers, and learn how to enhance your ads using features like ad extensions. If you run into any problems with your ads, find out how to tell if they're running and how to resolve approval issues.
            </Typography>

            {estimation.map((item, i) => {
              return <Button
                key={i}
                color="primary"
                raised={this.state.estimation === item}
                onClick={this.handleActiveButton(item)}
                className={classes.estimateButton}>
                  { item }
                </Button>
            })}

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
                onClick={this.sendEstimate(this.state.estimation)}
                className={classes.button} >
                Estimate
              </Button>
            </div>

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
      <Card>
        <div className={classes.root}>
          <Stepper activeStep={activeStep} orientation="vertical">
            {steps.map((label, index) => {
              return (
                <Step key={label}>
                  <StepLabel>{label}</StepLabel>
                  <StepContent>
                    <div className={classes.actionsContainer}>

                      {this.getStepContent(index)}

                    </div>
                  </StepContent>
                </Step>
              );
            })}
          </Stepper>
          {activeStep === steps.length && (
            <Paper square elevation={0} className={classes.resetContainer}>
              <Typography>Please wait until all yuor teammates complete estimation</Typography>
              <Button onClick={() => {this.props.actions.setStep(0)}} className={classes.button}>
                Reset
              </Button>
            </Paper>
          )}
        </div>
      </Card>
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