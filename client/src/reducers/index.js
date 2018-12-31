import { combineReducers } from 'redux';
import collections from './collections';
import items from './items';
import users from './users';

const rootReducer = combineReducers({
  collections,
  items,
  users
});

export default rootReducer;
