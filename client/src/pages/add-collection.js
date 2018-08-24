import React from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";

import Details from "./../components/details";
import { addCollection } from "./../actions/collection";

const styles = theme => ({
  textField: {
    width: "100%"
  },
  buttons: {
    display: "flex",
    justifyContent: "flex-end"
  },
  create: {
    marginTop: theme.spacing.unit * 3,
    marginLeft: theme.spacing.unit
  }
});

class AddCollection extends React.Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.state = { name: "" };
  }

  handleChange(event) {
    this.setState({ name: event.target.value });
  }

  handleSubmit(event) {
    this.props.addCollection(this.state.name);
    event.preventDefault();
  }

  render() {
    const { classes } = this.props;

    return (
      <Details name="Add collection">
        <form onSubmit={this.handleSubmit}>
          <TextField
            id="name"
            label="Name"
            className={classes.textField}
            margin="normal"
            value={this.state.name}
            onChange={this.handleChange}
          />
          <div className={classes.buttons}>
            <Button
              type="submit"
              variant="raised"
              color="primary"
              className={classes.create}
            >
              Create
            </Button>
          </div>
        </form>
      </Details>
    );
  }
}

AddCollection.propTypes = {
  classes: PropTypes.object.isRequired
};

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  addCollection: name => {
    dispatch(addCollection(name));
  }
});

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(AddCollection)
);
