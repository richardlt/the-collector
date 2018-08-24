import React from "react";
import { connect } from "react-redux";
import { Switch, Route, Redirect } from "react-router";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import Paper from "@material-ui/core/Paper";
import Avatar from "@material-ui/core/Avatar";
import CssBaseline from "@material-ui/core/CssBaseline";
import LockIcon from "@material-ui/icons/Lock";

import CollectionsPage from "./collections.js";
import AddCollectionPage from "./add-collection.js";
import SettingsPage from "./settings.js";
import NoMatchPage from "./nomatch.js";
import { fetchMe } from "./../actions/user";

const styles = theme => ({
  layout: {
    width: "auto",
    marginLeft: theme.spacing.unit * 3,
    marginRight: theme.spacing.unit * 3,
    [theme.breakpoints.up(400 + theme.spacing.unit * 3 * 2)]: {
      width: 400,
      marginLeft: "auto",
      marginRight: "auto"
    }
  },
  paper: {
    marginTop: theme.spacing.unit * 8,
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    padding: `${theme.spacing.unit * 2}px ${theme.spacing.unit * 3}px ${theme
      .spacing.unit * 3}px`
  },
  avatar: {
    margin: theme.spacing.unit,
    backgroundColor: theme.palette.secondary.main
  },
  form: {
    marginTop: theme.spacing.unit
  },
  submit: {
    marginTop: theme.spacing.unit * 3
  }
});

class Login extends React.Component {
  componentDidMount() {
    this.props.fetchMe();
  }

  render() {
    const { classes } = this.props;
    return this.props.me ? (
      <Switch>
        <Redirect exact path="/" to="/collections" />
        <Route path="/addCollection" component={AddCollectionPage} />
        <Route path="/collections" component={CollectionsPage} />
        <Route path="/settings" component={SettingsPage} />
        <Route component={NoMatchPage} />
      </Switch>
    ) : (
      <React.Fragment>
        <CssBaseline />
        <main className={classes.layout}>
          <Paper className={classes.paper}>
            <Avatar className={classes.avatar}>
              <LockIcon />
            </Avatar>
            <Typography variant="headline">Sign in</Typography>
            <form
              className={classes.form}
              action="/api/auth/login"
              method="get"
            >
              <Button
                type="submit"
                fullWidth
                variant="raised"
                color="primary"
                className={classes.submit}
              >
                With Facebook
              </Button>
            </form>
          </Paper>
        </main>
      </React.Fragment>
    );
  }
}

Login.propTypes = {
  classes: PropTypes.object.isRequired
};

const mapStateToProps = state => ({
  me: state.users.me
});

const mapDispatchToProps = dispatch => ({
  fetchMe: _ => {
    dispatch(fetchMe());
  }
});

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(Login)
);
