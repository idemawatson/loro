import { Switch, Route } from "react-router-dom";

import App from "App";
import Error from "pages/Error";
import User from "pages/User";

export const Path = {
  app: "/",
  user: "/user",
  error: "*",
};

const routes = (
  <Switch>
    <Route exact path={Path.app} component={App}></Route>
    <Route exact path={Path.user} component={User}></Route>
    <Route exact path={Path.error} component={Error}></Route>
  </Switch>
);

export default routes;
