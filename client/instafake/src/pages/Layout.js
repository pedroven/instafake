import React, { Component } from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Profile from "./profile/Profile";
import Header from "../components/Header";

class Layout extends Component {
  render() {
    return (
      <>
      <Header />
      <div>
        <Router>
          <Switch>
            <Route path="/user/:id" exact component={Profile}></Route>
          </Switch>
        </Router>
      </div>
      <footer className="footer"></footer>
      </>
    );
  }
}

export default Layout;
