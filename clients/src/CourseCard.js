import React from 'react'; //import React Component
import { useHistory } from 'react-router-dom';
// import { useParams } from "react-router-dom";
// import { useEffect } from "react";
import axios from 'axios';

export default function CourseCard(props) {
  let history = useHistory();


  const handleClick = () => {
    history.push("/class/" + props.courseID);
  }

  const handleStudentRegistration = () => {
    console.log("we're registering!")
      axios({
        "method": "POST",
        "url": "https://studybuddy-api.kaylalee.me/register-class",
        "headers": {
          "Content-Type": "application/json",
          "Authorization": props.authToken
        },
        "data": {
          courseID: props.courseID,
          studentID: props.studentID
        }
      })
      .then((response) => {
        console.log(response.data)
      })
      .catch((error) => {
        console.log(error)
      })
  }

  const handleExpertRegistration = () => {
    axios({
      "method": "POST",
      "url": "https://studybuddy-api.kaylalee.me/register-expert",
      "headers": {
        "Content-Type": "application/json",
        "Authorization": props.authToken
      },
      "data": {
        courseID: props.courseID,
        ExpertID: props.studentID
      }
    })
    .then((response) => {
      console.log(response.data)
    })
    .catch((error) => {
      console.log(error)
    })
}

  return (
    <div className="col-sm-12 col-md-6 col-xl-4">
      <div className="card">
        <div className="card-body"  onClick={handleClick}>
          <h5 className="card-title">{props.course}</h5>
          <h6 className="card-subtitle mb-2 text-muted">Professor: &nbsp; {props.professor}</h6>
          <p className="card-text">Department: &nbsp;{props.department}</p>
          <p className="card-text">Quarter: &nbsp;{props.quarter}</p>
        </div>
        <div className="registration row justify-content-center">
          <button type="button" className="btn btn-outline-primary col-5" >Enroll as Student</button>
          <button type="button" className="btn btn-outline-primary col-5">Enroll as Expert</button>  
        </div>
      </div>
    </div> 
  )
}