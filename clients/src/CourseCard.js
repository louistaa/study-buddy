import React from 'react'; //import React Component
import { useHistory } from 'react-router-dom';
// import { useParams } from "react-router-dom";
// import { useEffect } from "react";
// import axios from 'axios';

export default function CourseCard(props) {
  let history = useHistory();


  const handleClick = () => {
    history.push("/class/" + props.courseID);
  }


  
  // useEffect(() => {
  //   axios({
  //     "method": "POST",
  //     "url": "https://studybuddy-api.kaylalee.me/register-student/" + props.courseID,
  //     "headers": {
  //       "Authorization": props.authToken
  //     }
  //   })
  //   .then((response) => {
  //     setStudents(response.data);
  //   })
  //   .catch((error) => {
  //     console.log(error)
  //   })
  // }, []);

  return (
    <div className="col-sm-12 col-md-6 col-xl-4">
      <div className="card" onClick={handleClick}>
        <div className="card-body">
          <h5 className="card-title">{props.course}</h5>
          <h6 className="card-subtitle mb-2 text-muted">Professor: &nbsp; {props.professor}</h6>
          <p className="card-text">Department: &nbsp;{props.department}</p>
          <p className="card-text">Quarter: &nbsp;{props.quarter}</p>
        </div>
        <div className="registration">
          <button type="button" className="btn btn-outline-primary" >Enroll as Student</button>
          <button type="button" className="btn btn-outline-primary">Enroll as Expert</button>  
        </div>
      </div>
    </div> 
  )
}