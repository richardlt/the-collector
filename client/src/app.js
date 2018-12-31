import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { connect } from 'react-redux';
import { Route, Switch, Redirect } from 'react-router-dom';
import CircularProgress from '@material-ui/core/CircularProgress';
import CollectionsPage from './pages/collections';
import AddCollectionPage from './pages/add-collection';
import SettingsPage from './pages/settings';
import LoginPage from './pages/login';
import NoMatchPage from './pages/nomatch';
import { fetchMe } from './actions/user';

const styles = theme => ({
  progress: {
    position: 'absolute',
    bottom: 0,
    right: 0,
    margin: theme.spacing.unit * 2
  }
});

class App extends React.Component {
  constructor(props) {
    super(props);
  }
  componentDidMount() {
    this.props.fetchMe();
  }

  render() {
    const { classes } = this.props;

    return this.props.me ? (
      <Switch>
        <Redirect exact path="/" to="/collections" />
        <Route path="/collections" component={CollectionsPage} />
        <Route path="/addCollection" component={AddCollectionPage} />
        <Route path="/settings" component={SettingsPage} />
        <Route component={NoMatchPage} />
      </Switch>
    ) : this.props.error ? (
      <Switch>
        <Route exact path="/" component={LoginPage} />
        <Redirect to="/" />
      </Switch>
    ) : (
      <CircularProgress size={50} className={classes.progress} />
    );
  }
}

const mapStateToProps = state => ({
  me: state.users.me,
  loading: state.users.loading,
  error: state.users.error
});

const mapDispatchToProps = dispatch => ({
  fetchMe: _ => {
    dispatch(fetchMe());
  }
});

App.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(App)
);
