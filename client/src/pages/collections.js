import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { withStyles } from "@material-ui/core/styles";
import { Route } from "react-router-dom";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemAvatar from "@material-ui/core/ListItemAvatar";
import ListItemText from "@material-ui/core/ListItemText";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import FolderIcon from "@material-ui/icons/Folder";
import AddIcon from "@material-ui/icons/Add";
import { push } from "connected-react-router";

import Page from "./../components/page.js";
import CollectionPage from "./collection.js";
import { fetchCollections } from "./../actions/collection";

const styles = theme => ({
  list: { height: "100%", overflow: "auto" },
  fab: {
    position: "absolute",
    bottom: theme.spacing.unit * 2,
    right: theme.spacing.unit * 2
  }
});

class Collections extends React.Component {
  constructor(props) {
    super(props);
    this.handleOpen = this.handleOpen.bind(this);
  }

  componentDidMount() {
    this.props.fetchCollections();
  }

  handleOpen(slug) {
    return _ => {
      this.props.openCollection(slug);
    };
  }

  render() {
    const { classes, match } = this.props;
    
    return (
      <React.Fragment>
        <Route
          exact
          path={match.url}
          render={() => (
            <Page title="My collections">
              <List className={classes.list}>
                {this.props.collections.map(collection => {
                  return (
                    <ListItem
                      button
                      key={collection.uuid}
                      onClick={this.handleOpen(collection.slug)}
                    >
                      <ListItemAvatar>
                        <Avatar>
                          <FolderIcon />
                        </Avatar>
                      </ListItemAvatar>
                      <ListItemText primary={collection.name} />
                    </ListItem>
                  );
                })}
              </List>
              <Button
                variant="fab"
                className={classes.fab}
                color="primary"
                onClick={this.props.addCollection}
              >
                <AddIcon />
              </Button>
            </Page>
          )}
        />
        <Route
          path={`${match.url}/:collectionSlug`}
          component={CollectionPage}
        />
      </React.Fragment>
    );
  }
}

Collections.propTypes = {
  classes: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired,
  match: PropTypes.object.isRequired
};

const mapStateToProps = state => ({
  collections: state.collections.all,
  loading: state.collections.loading,
  error: state.collections.error
});

const mapDispatchToProps = dispatch => ({
  openCollection: slug => {
    dispatch(push("/collections/" + slug));
  },
  fetchCollections: _ => {
    dispatch(fetchCollections());
  },
  addCollection: _ => {
    dispatch(push("/addCollection"));
  }
});

export default withStyles(styles, { withTheme: true })(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(Collections)
);
