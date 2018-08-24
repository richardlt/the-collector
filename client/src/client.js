import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { applyMiddleware, compose, createStore } from "redux";
import { Route, Switch } from "react-router-dom";
import { routerMiddleware, connectRouter } from "connected-react-router";
import thunk from "redux-thunk";

import LoginPage from "./pages/login.js";
import { createBrowserHistory } from "history";
import { ConnectedRouter } from "connected-react-router";
import rootReducer from "./reducers";

const history = createBrowserHistory();

const composeEnhancer = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store = createStore(
  connectRouter(history)(rootReducer),
  applyMiddleware(thunk),
  composeEnhancer(applyMiddleware(routerMiddleware(history)))
);

ReactDOM.render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <Switch>
        <Route path="/" component={LoginPage} />
      </Switch>
    </ConnectedRouter>
  </Provider>,
  document.getElementById("root")
);
