import React from "react";

import { useState, useEffect } from "react";
import { Redirect } from "react-router-dom";

export default function MyProfile(props) {
  // fetch a list of the people in a course
  // const [data, setData] = useState([]);
  // const [redirectTo, setRedirectTo] = useState(undefined)

  // useEffect(() => {
  //   fetch("./mockProfile.json")
  //     .then((res) => res.json())
  //     .then((data) => {
  //       setData(data.people); //change the state and re-render
  //     });
  // }, []);

  return (
    <div>
      <div className="students">Your profile</div>
      <div className="students">First Name: {props.user.firstName} </div>
      <div className="students">Last Name: {props.user.lastName} </div>
      <div className="students">Your Registered Classes:  </div>
      <div className="students">Your Expert Classes: </div>
      <div className="students">Your E-mail: {props.user.email} </div>
      <div className="students">Your Phone Number : {props.user.phoneNumber} </div>
    </div>
  );
}
