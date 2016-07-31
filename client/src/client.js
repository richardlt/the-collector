import React from 'react'
import ReactDOM from 'react-dom'
import { Router, Route, IndexRoute, browserHistory } from 'react-router'

import Store from './store.js'
import Collections from './pages/collections.page.js'
import Collection from './pages/collection.page.js'

ReactDOM.render(
    <Router history={browserHistory}>
        <Route path="/">
            <IndexRoute component={Collections} />
            <Route path="/:collectionUUID" component={Collection} />
        </Route>
    </Router>,
    document.getElementById('root')
)
