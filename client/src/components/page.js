import React from "react";
import { connect } from "react-redux";
import { push } from "connected-react-router";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import ViewModuleIcon from "@material-ui/icons/ViewModule";
import SettingsIcon from "@material-ui/icons/SettingsOutlined";
import Divider from "@material-ui/core/Divider";
import Hidden from "@material-ui/core/Hidden";
import Drawer from "@material-ui/core/Drawer";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import CssBaseline from "@material-ui/core/CssBaseline";

const drawerWidth = 240;

const styles = theme => ({
  root: {
    position: "relative",
    display: "flex",
    flexDirection: "row",
    width: "100%",
    height: "100%"
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
    [theme.breakpoints.up("md")]: {
      display: "none"
    }
  },
  drawerWrapper: { width: drawerWidth },
  drawerPaper: {
    width: drawerWidth,
    [theme.breakpoints.up("md")]: {
      position: "fixed"
    }
  },
  contentWrapper: {
    flex: 1,
    height: "100%",
    display: "flex",
    flexDirection: "column"
  },
  appBar: { position: "static" },
  content: {
    flex: 1,
    overflow: "hidden"
  }
});

class Page extends React.Component {
  constructor(props) {
    super(props);
    this.state = { menuOpen: false };
    this.handleOpenMenu = this.handleOpenMenu.bind(this);
  }

  handleOpenMenu() {
    this.setState(state => ({ menuOpen: !state.menuOpen }));
  }

  render() {
    const { classes, goCollections, goSettings } = this.props;

    const drawer = (
      <div>
        <div className={classes.toolbar} />
        <ListItem button onClick={goCollections}>
          <ListItemIcon>
            <ViewModuleIcon />
          </ListItemIcon>
          <ListItemText primary="My collections" />
        </ListItem>
        <ListItem button onClick={goSettings}>
          <ListItemIcon>
            <SettingsIcon />
          </ListItemIcon>
          <ListItemText primary="Settings" />
        </ListItem>
        <Divider />
      </div>
    );

    return (
      <React.Fragment>
        <CssBaseline />
        <div className={classes.root}>
          <Hidden mdUp className={classes.drawerWrapper}>
            <Drawer
              variant="temporary"
              anchor={"left"}
              open={this.state.menuOpen}
              onClose={this.handleOpenMenu}
              classes={{ paper: classes.drawerPaper }}
              ModalProps={{ keepMounted: true }}
            >
              {drawer}
            </Drawer>
          </Hidden>
          <Hidden smDown implementation="css" className={classes.drawerWrapper}>
            <Drawer
              variant="permanent"
              open
              classes={{ paper: classes.drawerPaper }}
            >
              {drawer}
            </Drawer>
          </Hidden>
          <div className={classes.contentWrapper}>
            <AppBar className={classes.appBar}>
              <Toolbar>
                <IconButton
                  className={classes.menuButton}
                  color="inherit"
                  aria-label="Menu"
                  onClick={this.handleOpenMenu}
                >
                  <MenuIcon />
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

Page.propTypes = {
  classes: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired,
  title: PropTypes.string.isRequired
};

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  goCollections: _ => {
    dispatch(push("/collections"));
  },
  goSettings: _ => {
    dispatch(push("/settings"));
  }
});

export default withStyles(styles, { withTheme: true })(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(Page)
);
