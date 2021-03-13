import React, { Component } from "react";
import Auth from "./Components/Auth/Auth";
import PageTypes from "./Constants/PageTypes/PageTypes";
import Main from "./Components/Main/Main";
import "./Styles/App.css";
import api from "./Constants/APIEndpoints/APIEndpoints";

import ClassPage from "./ClassPage";
import Courses from "./Courses";
import MyProfile from "./MyProfile";
import { Route, Switch, Redirect } from "react-router-dom";
import NewCourse from "./NewCourse";

class App extends Component {
  constructor() {
    super();
    this.state = {
      page: localStorage.getItem("Authorization")
        ? PageTypes.signedInMain
        : PageTypes.signIn,
      authToken: localStorage.getItem("Authorization") || null,
      user: null,
    };

    this.getCurrentUser();
  }

  /**
   * @description Gets the users
   */
  getCurrentUser = async () => {
    if (!this.state.authToken) {
      return;
    }
    const response = await fetch(api.base + api.handlers.myuser, {
      headers: new Headers({
        Authorization: this.state.authToken,
      }),
    });
    if (response.status >= 300) {
      alert("Unable to verify login. Logging out...");
      localStorage.setItem("Authorization", "");
      this.setAuthToken("");
      this.setUser(null);
      return;
    }
    const user = await response.json();
    this.setUser(user);
  };

  /**
   * @description sets the page type to sign in
   */
  setPageToSignIn = (e) => {
    e.preventDefault();
    this.setState({ page: PageTypes.signIn });
  };

  /**
   * @description sets the page type to sign up
   */
  setPageToSignUp = (e) => {
    e.preventDefault();
    this.setState({ page: PageTypes.signUp });
  };

  setPage = (e, page) => {
    e.preventDefault();
    this.setState({ page });
  };

  /**
   * @description sets auth token
   */
  setAuthToken = (authToken) => {
    this.setState({
      authToken,
      page: authToken === "" ? PageTypes.signIn : PageTypes.signedInMain,
    });
  };

  /**
   * @description sets the user
   */
  setUser = (user) => {
    this.setState({ user });
  };

  render() {
    const { page, user } = this.state;
    return (
      <div>
        {user ? (
          <div>
            <Main
              page={page}
              setPage={this.setPage}
              setAuthToken={this.setAuthToken}
              user={user}
              setUser={this.setUser}
            />

            <div className="row">
              <a href="/" className="col studyBuddy">
                iSTUDY BUDDY
              </a>
              <a href="/myprofile" className="col chats">
                My Profile
              </a>
              <a href="/newCourse" className="myProfile">
                Add Course
              </a>
            </div>

            <Switch>
              <Route exact path="/" component={Courses} />
              <Route exact path="/myprofile" component={MyProfile} />
              <Route exact path="/newCourse" component={NewCourse} />
              <Route exact path="/:courseName" component={ClassPage} />
              <Route exact path="/students/:profile" component={ClassPage} />
              <Redirect to="/" />
            </Switch>
          </div>
        ) : (
          <div>
            <a href="/" className="col studyBuddy">
              iSTUDY BUDDY
            </a>

            <div className="row justify-content-center pleaseSignIn">Please Sign In or Sign Up</div>

            <Auth
              page={page}
              setPage={this.setPage}
              setAuthToken={this.setAuthToken}
              setUser={this.setUser}
            />
          </div>
        )}

        <div className="row justify-content-center">
          &copy; 2021 iStudy Buddy, YuYu Madigan, Kayla Lee, Saatvik Arya, Louis
          Ta
        </div>
      </div>
    );
  }
}

export default App;
