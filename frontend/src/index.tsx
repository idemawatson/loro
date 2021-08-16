import ReactDOM from "react-dom";
import { createBrowserHistory } from "history";
import { Router } from "react-router-dom";

import routes from "router";

import "assets/css/index.css";

const history = createBrowserHistory();

ReactDOM.render(
  <Router history={history}>{routes}</Router>,
  document.getElementById("root")
);
