import React from 'react'

import { useState } from "react";
import { Redirect, useHistory } from 'react-router-dom';

export default function StudentCard(props) {
  let history = useHistory();

  const handleClick = () => {
    history.push("/profiles/" + props.id)
  }

  return (
    <div className="col-9">
      <div className="card" onClick={handleClick}>
        <div className="card-body">
          <h5 className="card-title">{props.person}</h5>
          <h6 className="card-subtitle mb-2 text-muted">Username: {props.username} Major: {props.major}</h6>
          <p className="card-text">Phone number: {props.phonenumber}</p> 
          <p> Email: {props.email}</p>
          Click to view profile
        </div>
      </div>
    </div>
  );
}
