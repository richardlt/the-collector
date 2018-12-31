import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { Grid, Button } from '@material-ui/core';
import { Slider } from '@material-ui/lab';
import { push } from 'connected-react-router';
import { fetchItem, updateItemFile, deleteItem } from './../actions/item';
import Details from './../components/details.js';
import AvatarEditor from 'react-avatar-editor';

const styles = theme => ({
  picture: {
    width: '50%'
  }
});

const defaultSettings = {
  scaleValue: 50,
  scaleRate: 1,
  rotateValue: 0,
  rotateAngle: 0
};

class Item extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      ...defaultSettings
    };
  }

  componentWillReceiveProps() {
    this.setState({
      ...this.state,
      ...defaultSettings
    });
  }

  deleteItem = () => {
    this.props.deleteItem(this.props.collection, this.props.item.uuid);
  };

  handleScaleChange = (_, value) => {
    this.setState({
      ...this.state,
      scaleValue: value,
      scaleRate: 1 + (value / 100 - 0.5)
    });
  };

  handleRotateChange = (_, value) => {
    this.setState({
      ...this.state,
      rotateValue: value,
      rotateAngle: Math.floor((value * 360) / 100)
    });
  };

  saveItem = () => {
    if (this.editor) {
      const canvas = this.editor.getImage();
      canvas.toBlob(
        blob => {
          const filename = this.props.item.picture.substring(
            this.props.item.picture.lastIndexOf('/') + 1
          );
          const file = new File([blob], filename.split('.')[0] + '.jpg', {
            type: 'image/jpeg'
          });
          this.props.updateItemFile(
            this.props.collection,
            this.props.item,
            file
          );
        },
        'image/jpeg',
        1
      );
    }
  };

  setEditorRef = editor => (this.editor = editor);

  render() {
    const { classes } = this.props;

    return (
      <Details title="Item details">
        <Grid container spacing={16}>
          <Grid item xs={12} sm={6}>
            <Grid container direction="column" alignItems="center">
              {this.props.item ? (
                <div>
                  <AvatarEditor
                    ref={this.setEditorRef}
                    image={this.props.item.picture}
                    className={classes.picture}
                    border={20}
                    color={[255, 255, 255, 0.6]} // RGBA
                    scale={this.state.scaleRate}
                    rotate={this.state.rotateAngle}
                  />
                  <Slider
                    classes={{ container: classes.slider }}
                    value={this.state.scaleValue}
                    aria-labelledby="label"
                    onChange={this.handleScaleChange}
                  />
                  <br />
                  <Slider
                    classes={{ container: classes.slider }}
                    value={this.state.rotateValue}
                    aria-labelledby="label"
                    onChange={this.handleRotateChange}
                  />
                </div>
              ) : (
                ''
              )}
            </Grid>
          </Grid>
          <Grid item xs={12} sm={6}>
            <Grid container direction="column" alignItems="center">
              <Grid item>
                <Button
                  variant="contained"
                  color="primary"
                  onClick={this.saveItem}
                >
                  Save
                </Button>
              </Grid>
              <br />
              <Grid item>
                <Button
                  variant="contained"
                  color="secondary"
                  onClick={this.deleteItem}
                >
                  Delete
                </Button>
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

const mapStateToProps = (state, props) => {
  const uuid = props.match.params.itemUUID;
  return {
    collection: state.collections.current,
    item: state.items ? state.items.all.find(i => i.uuid == uuid) : null
  };
};

const mapDispatchToProps = dispatch => ({
  fetchItem: (collection, itemUUID) => {
    dispatch(fetchItem(collection, itemUUID));
  },
  deleteItem: (collection, itemUUID) => {
    dispatch(deleteItem(collection, itemUUID)).then(_ => {
      dispatch(push('/collections/' + collection.slug));
    });
  },
  updateItemFile: (collection, item, file) => {
    dispatch(updateItemFile(collection, item, file));
  }
});

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(Item)
);
