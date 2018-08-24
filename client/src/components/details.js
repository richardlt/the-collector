import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { goBack } from "connected-react-router";
import { withStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import ArrowBackIcon from "@material-ui/icons/ArrowBack";
import CssBaseline from "@material-ui/core/CssBaseline";

const styles = theme => ({
  root: {
    position: "relative",
    display: "flex",
    flexDirection: "row",
    width: "100%"
  },
  backButton: {
    marginLeft: -12,
    marginRight: 20
  },
  appBar: { position: "sticky" },
  contentWrapper: { flex: 1 },
  content: { padding: theme.spacing.unit }
});

class Details extends React.Component {
  render() {
    const { classes, goBack } = this.props;

    return (
      <React.Fragment>
        <CssBaseline />
        <div className={classes.root}>
          <div className={classes.contentWrapper}>
            <AppBar className={classes.appBar}>
              <Toolbar>
                <IconButton
                  className={classes.backButton}
                  color="inherit"
                  aria-label="Back"
                  onClick={goBack}
                >
                  <ArrowBackIcon />
                </IconButton>
                <Typography variant="title" color="inherit">
                  {this.props.title}
                </Typography>
              </Toolbar>
            </AppBar>
            <div className={classes.content}>{this.props.children}</div>
          </div>
        </div>
      </React.Fragment>
    );
  }
}

Details.propTypes = {
  classes: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired,
  title: PropTypes.string.isRequired
};

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  goBack: _ => {
    dispatch(goBack());
  }
});

export default withStyles(styles, { withTheme: true })(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(Details)
);
