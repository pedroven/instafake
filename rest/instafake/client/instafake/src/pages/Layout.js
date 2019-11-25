import React, { Component } from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Profile from "./profile/Profile";

class Layout extends Component {
  render() {
    return (
      <div>
        <Router>
          <Switch>
            <Route path="/user/:id" exact component={Profile}></Route>
          </Switch>
        </Router>
      </div>
    );
  }
}

export default Layout;
