import React from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";

import { fetchItem } from "./../actions/item";
import Details from "./../components/details.js";

const styles = theme => ({});

class Item extends React.Component {
  constructor(props) {
    super(props);

    this.refreshItem = this.refreshItem.bind(this);
  }

  componentDidMount() {
    this.refreshItem(this.props);
  }

  componentWillReceiveProps(nextProps) {
    this.refreshItem(nextProps);
  }

  refreshItem(props) {
    const uuid = props.match.params.itemUUID;
    if (props.collection && (!props.item || props.item.uuid != uuid)) {
      this.props.fetchItem(props.collection, uuid);
    }
  }

  render() {
    return (
      <Details title="Item details">
        {this.props.item ? (
          <img width="100%" src={this.props.item.picture} />
        ) : (
          ""
        )}
      </Details>
    );
  }
}

Item.propTypes = {
  classes: PropTypes.object.isRequired,
  match: PropTypes.object.isRequired
};

const mapStateToProps = state => ({
  collection: state.collections.current,
  item: state.items.current
});

const mapDispatchToProps = dispatch => ({
  fetchItem: (collection, itemUUID) => {
    dispatch(fetchItem(collection, itemUUID));
  }
});

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(Item)
);
