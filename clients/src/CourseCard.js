import React, { useState } from 'react'; //import React Component
import { Redirect } from 'react-router-dom';

export default function CourseCard(props) {
  const [redirectTo, setRedirectTo] = useState(undefined)

  const handleClick = () => {
    // make INFO XXX -> INFOXXX
    setRedirectTo(props.course.replace(/\s/g, ""))
  }

  if (redirectTo) {
    return <Redirect push to={ '/' + redirectTo} />
  }

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
          <button type="button" class="btn btn-outline-primary">Enroll as Student</button>
          <button type="button" class="btn btn-outline-primary">Enroll as Expert</button>  
        </div>
      </div>
    </div> 
  )
}