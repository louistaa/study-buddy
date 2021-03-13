import React from "react";
import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from 'axios';

export default function UserProfile(props) {
  let { person } = useParams();

  const [student, setStudent] = useState([]);


  useEffect(() => {
    axios({
      "method": "GET",
      "url": "https://studybuddy-api.kaylalee.me/students/" + person,
      "headers": {
        "Authorization": props.authToken
      }
    })
    .then((response) => {
      setStudent(response.data);
    })
    .catch((error) => {
      console.log(error)
    })
  }, [])

  return (
    <div>
      <div className="students">{student.firstName}'s profile</div>
      <div className="students">Name: {student.firstName} {student.lastName}</div>
      <div className="students">Major: {student.major}</div>
      <div className="students">E-mail: {student.email}</div>
      <div className="students">Phone Number: {student.phoneNumber}</div>
      <div className="students">Registered Classes:  </div>
      <div className="students">Expert Classes: </div>
      
      
    </div>
  );
}