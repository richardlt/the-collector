import React from 'react';
import ReactDOM from 'react-dom';
import { connect } from 'react-redux';
import { Route, Switch } from 'react-router-dom';
import Modal from '@material-ui/core/Modal';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import GridList from '@material-ui/core/GridList';
import GridListTile from '@material-ui/core/GridListTile';
import Button from '@material-ui/core/Button';
import AddAPhotoIcon from '@material-ui/icons/AddAPhoto';
import SettingsIcon from '@material-ui/icons/Settings';
import { push } from 'connected-react-router';
import Page from './../components/page.js';
import ItemPage from './item.js';
import { fetchCollection } from './../actions/collection';
import { fetchItems, addItem } from './../actions/item';

const styles = theme => ({
  gridWrapper: {
    width: '100%',
    height: '100%',
    overflow: 'auto'
  },
  grid: {
    width: '100%',
    overflow: 'hidden'
  },
  tile: {
    padding: '2px'
  },
  fab: {
    position: 'absolute',
    bottom: theme.spacing.unit * 2,
    right: theme.spacing.unit * 2
  },
  inputFile: {
    width: '0.1px',
    height: '0.1px',
    opacity: 0,
    overflow: 'hidden',
    position: 'absolute',
    zIndex: -1
  },
  paper: {
    position: 'absolute',
    width: theme.spacing.unit * 50,
    top: '50%',
    left: '50%',
    transform: `translate(-50%, -50%)`
  },
  preview: {
    width: '100%'
  },
  icon: {
    color: 'white'
  }
});

const BASE_COLUMN_SIZE = 100;

class List extends React.Component {
  constructor(props) {
    super(props);

    this.gridRef = React.createRef();
    this.state = {
      columns: 10,
      cellHeight: BASE_COLUMN_SIZE,
      modalOpen: false
    };
    this.inputFileRef = React.createRef();
  }

  componentDidMount() {
    window.addEventListener('resize', this.resize.bind(this));
    this.resize();
    this.props.fetchCollectionAndItems(this.props.match.params.collectionSlug);
  }

  componentWillReceiveProps() {
    // resize when receiving props, handle navigate from item to collection
    this.resize();
  }

  resize = () => {
    if (this.gridRef.current) {
      const { width } = ReactDOM.findDOMNode(
        this.gridRef.current
      ).getBoundingClientRect();
      const columns = Math.floor(width / BASE_COLUMN_SIZE);
      this.setState({
        columns,
        cellHeight: Math.floor(width / columns)
      });
    }
  };

  handleChange = () => {
    if (this.props.collection) {
      this.props.addItem(
        this.props.collection,
        this.inputFileRef.current.files[0]
      );
    }
  };

  openModal = selectedItem => {
    this.setState({
      ...this.state,
      modalOpen: true,
      selectedItem
    });
  };

  closeModal = item => {
    this.setState({
      ...this.state,
      modalOpen: false
    });
  };

  openItem = () => {
    this.props.openItem(this.props.collection, this.state.selectedItem);
    this.setState({
      ...this.state,
      modalOpen: false
    });
  };

  render() {
    const { classes, match } = this.props;

    const list = (
      <GridList
        ref={this.gridRef}
        cellHeight={this.state.cellHeight}
        cols={this.state.columns}
        className={classes.grid}
        spacing={0}
      >
        {this.props.items.map(item => (
          <GridListTile
            key={item.uuid}
            classes={{
              tile: classes.tile
            }}
            onClick={_ => {
              this.openModal(item);
            }}
          >
            <img src={item.picture + '?size=medium'} alt={item.uuid} />
          </GridListTile>
        ))}
      </GridList>
    );

    const modal =
      this.state.modalOpen && this.state.selectedItem ? (
        <Modal open={this.state.modalOpen} onClose={this.closeModal}>
          <div>
            <div className={classes.paper}>
              <img
                className={classes.preview}
                src={this.state.selectedItem.picture + '?size=large'}
                alt={this.state.selectedItem.uuid}
              />
            </div>
            <Button
              variant="fab"
              className={classes.fab}
              color="primary"
              onClick={this.openItem}
            >
              <SettingsIcon />
            </Button>
          </div>
        </Modal>
      ) : (
        ''
      );

    return (
      <Switch>
        <Route
          exact
          path={match.url}
          render={() => (
            <Page
              title={this.props.collection ? this.props.collection.name : ''}
            >
              {modal}
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
        <Route path={`${match.url}/:itemUUID`} component={ItemPage} />
      </Switch>
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
    dispatch(push('/collections/' + collection.slug + '/' + item.uuid));
  },
  fetchCollectionAndItems: slug => {
    dispatch(fetchCollection(slug)).then(collection => {
      dispatch(fetchItems(collection));
    });
  },
  addItem: (collection, file) => {
    dispatch(addItem(collection, file));
  }
});

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(List)
);
