import React from 'react'

import { useState } from "react";
import { Redirect } from 'react-router-dom';

export default function StudentCard(props) {
  const [redirectTo, setRedirectTo] = useState(undefined)

  const handleClick = () => {
    setRedirectTo(props.person)
  }

<<<<<<< HEAD
  if (redirectTo) {
    return <Redirect push to={'/profiles/' + redirectTo} />
  }
=======
  // if (redirectTo) {
  //   return <Redirect push to={'/chats/' + redirectTo} />
  // }
>>>>>>> saatvik

  return (
    <div className="col-9">
      <div className="card" onClick={handleClick}>
        <div className="card-body">
          <h5 className="card-title">{props.person}</h5>
<<<<<<< HEAD
          <h6 className="card-subtitle mb-2 text-muted">{props.status}</h6>
          Click to view {props.person}'s profile
=======
          <h6 className="card-subtitle mb-2 text-muted">Username: {props.username} Major: {props.major}</h6>
          <p className="card-text">Phone number: {props.phonenumber} Email: {props.email}</p>
        </div>
        <div className="contact">
          <button type="button" class="btn btn-outline-primary">View Profile</button>
>>>>>>> saatvik
        </div>
      </div>
    </div>
  );
}
