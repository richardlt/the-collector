import React from "react";
import ReactDOM from "react-dom";
import { connect } from "react-redux";
import { Route } from "react-router-dom";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import GridList from "@material-ui/core/GridList";
import GridListTile from "@material-ui/core/GridListTile";
import Button from "@material-ui/core/Button";
import AddAPhotoIcon from "@material-ui/icons/AddAPhoto";
import { push } from "connected-react-router";

import Page from "./../components/page.js";
import ItemPage from "./item.js";
import { fetchCollection } from "./../actions/collection";
import { fetchItems, addItem } from "./../actions/item";

const styles = theme => ({
  gridWrapper: {
    width: "100%",
    height: "100%",
    overflow: "auto"
  },
  grid: {
    width: "100%"
  },
  link: {
    display: "inline-block"
  },
  titleBar: {
    background:
      "linear-gradient(to bottom, rgba(0,0,0,0.7) 0%, " +
      "rgba(0,0,0,0.3) 70%, rgba(0,0,0,0) 100%)"
  },
  icon: {
    color: "white"
  },
  fab: {
    position: "absolute",
    bottom: theme.spacing.unit * 2,
    right: theme.spacing.unit * 2
  },
  inputFile: {
    width: "0.1px",
    height: "0.1px",
    opacity: 0,
    overflow: "hidden",
    position: "absolute",
    zIndex: -1
  }
});

const BASE_COLUMN_SIZE = 100;

class List extends React.Component {
  constructor(props) {
    super(props);
    this.handleEdit = this.handleEdit.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.gridRef = React.createRef();
    this.state = { columns: 10 };
    this.inputFileRef = React.createRef();
  }

  componentDidMount() {
    window.addEventListener("resize", this.resize.bind(this));
    this.resize();
    this.props.fetchCollectionAndItems(this.props.match.params.collectionSlug);
  }

  resize() {
    if (this.gridRef.current) {
      const { width } = ReactDOM.findDOMNode(
        this.gridRef.current
      ).getBoundingClientRect();
      this.setState({ columns: Math.floor(width / BASE_COLUMN_SIZE) });
    }
  }

  handleChange() {
    if (this.props.collection) {
      this.props.addItem(
        this.props.collection,
        this.inputFileRef.current.files[0]
      );
    }
  }

  render() {
    const { classes, match } = this.props;

    const list = (
      <GridList
        ref={this.gridRef}
        cellHeight={BASE_COLUMN_SIZE}
        cols={this.state.columns}
        className={classes.grid}
      >
        {this.props.items.map(item => (
          <GridListTile key={item.uuid}>
            <img src={item.picture} alt={item.uuid} />
          </GridListTile>
        ))}
      </GridList>
    );

    return (
      <React.Fragment>
        <Route
          exact
          path={match.url}
          render={() => (
            <Page
              title={this.props.collection ? this.props.collection.name : ""}
            >
              <div className={classes.gridWrapper}>{list}</div>
              <input
                type="file"
                id="file"
                accept="image/*"
                className={classes.inputFile}
                onChange={this.handleChange}
                ref={this.inputFileRef}
              />
              <Button variant="fab" className={classes.fab} color="primary">
                <label htmlFor="file">
                  <AddAPhotoIcon />
                </label>
              </Button>
            </Page>
          )}
        />
        <Route path={`${match.url}/:itemID`} component={ItemPage} />
      </React.Fragment>
    );
  }
}

List.propTypes = {
  classes: PropTypes.object.isRequired,
  match: PropTypes.object.isRequired
};

const mapStateToProps = state => ({
  collection: state.collections.current,
  items: state.items.all,
  loading: state.collections.loading,
  error: state.collections.error
});

const mapDispatchToProps = dispatch => ({
  openItem: (collection, item) => {
    dispatch(push("/collections/" + collection.slug + "/" + item.uuid));
  },
  fetchCollectionAndItems: slug => {
    dispatch(fetchCollection(slug)).then(collection => {
      dispatch(fetchItems(collection));
    });
  },
  fetchItems: href => {
    dispatch(fetchItems(href));
  },
  addItem: (collection, file) => {
    dispatch(addItem(collection, file)).then(_ => {
      dispatch(fetchItems(collection));
    });
  }
});

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(List)
);
