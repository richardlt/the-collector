import { createStore, combineReducers } from 'redux'

import collections from './reducers/collections.reducer.js'
import items from './reducers/items.reducer.js'

const app = combineReducers({
    collections,
    items
})

export default createStore(app)
