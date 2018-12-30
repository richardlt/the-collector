import React from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import { Grid, Button } from "@material-ui/core";
import { push } from "connected-react-router";
import { fetchItem, deleteItem } from "./../actions/item";
import Details from "./../components/details.js";

const styles = theme => ({});

class Item extends React.Component {
  constructor(props) {
    super(props);

    this.refreshItem = this.refreshItem.bind(this);
    this.deleteItem = this.deleteItem.bind(this);
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

  deleteItem() {
    this.props.deleteItem(this.props.collection, this.props.item.uuid);
  }

  render() {
    return (
      <Details title="Item details">
        <Grid container spacing={16}>
          <Grid item xs={12} sm={8}>
            {this.props.item ? (
              <img width="100%" src={this.props.item.picture} />
            ) : (
                ""
              )}
          </Grid>
          <Grid item xs={12} sm={4}>
            <Grid container direction="column" alignItems="center">
              <Grid item>
                <Button variant="contained" color="secondary" onClick={this.deleteItem}>Delete</Button>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
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
  },
  deleteItem: (collection, itemUUID) => {
    dispatch(deleteItem(collection, itemUUID)).then(_ => {
      dispatch(push("/collections/" + collection.slug));
    });
  }
});

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(Item)
);
