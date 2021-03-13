import React from "react";
import { useParams } from "react-router-dom";

import { useState, useEffect } from "react";

export default function UserProfile(props) {
  const urlParams = useParams();

  let person = urlParams.person;

  return (
    <div>
      <div className="students">{person}'s profile</div>
      <div className="students">{person}'s Registered Classes:  </div>
      <div className="students">{person}'s Expert Classes: </div>
      <div className="students">{person}'s E-mail: </div>
      <div className="students">{person}'s Phone Number: </div>
    </div>
  );
}