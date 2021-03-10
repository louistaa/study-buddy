import React from "react";
import ClassPage from "./ClassPage";
import Courses from "./Courses";
import Chats from "./Chats"
import { Route, Switch, Redirect } from "react-router-dom";
import ChatSpecific from "./ChatSpecific";

export default function App() {
  return (
    <div className="container-fluid">
      <div className="row">
        <a href="/" className="col studyBuddy">iSTUDY BUDDY</a>
        <a href="/chats" className="col chats">Chats</a>
      </div>

      <Switch>
        <Route exact path="/" component={Courses} />
        <Route exact path="/chats" component={Chats} />
        <Route exact path="/:courseName" component={ClassPage} />
        <Route exact path="/chats/:person" component={ChatSpecific} />
        <Redirect to="/" />
      </Switch>

      <div className="row justify-content-center">
        &copy; 2021 iStudy Buddy, YuYu Madigan, Kayla Lee, Saatvik Arya, Louis Ta
      </div>
    </div>
  );
}