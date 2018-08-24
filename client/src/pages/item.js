import React from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";

import Details from "./../components/details.js";

const styles = theme => ({
  textField: {
    width: "100%"
  },
  formControl: {
    width: "100%"
  },
  chips: {
    display: "flex",
    flexWrap: "wrap"
  },
  chip: {
    margin: theme.spacing.unit / 4
  }
});

class Item extends React.Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
  }

  render() {
    return (
      <Details title="Item details">
        {this.props.item ? <img src={this.props.item.picture} /> : ""}
      </Details>
    );
  }
}

Item.propTypes = {
  classes: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired
};

const mapStateToProps = state => ({
  item: state.items.current,
  loading: state.items.loading,
  error: state.items.error
});

const mapDispatchToProps = dispatch => ({});

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(Item)
);
