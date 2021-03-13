import React from "react";
import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import StudentCard from './StudentCard';


export default function ClassPage(props) {
  const urlParams = useParams();

  // grab the course name from the URL params
  let courseName = urlParams.courseName;

  // fetch a list of the people in a course
  const [data, setData] = useState([]);

  
  // useEffect(() => {
  //   axios({
  //     "method": "GET",
  //     "url": "https://studybuddy-api.kaylalee.me/classes",
  //     "headers": {
  //       "Authorization": props.authToken
  //     }
  //   })
  //   .then((response) => {
  //     setClasses(response.data);
  //   })
  //   .catch((error) => {
  //     console.log(error)
  //   })
  // }, []);

  useEffect(() => {
    fetch("./classSpecificPeople.json")
      .then((res) => res.json())
      .then((data) => {
        setData(data.people); //change the state and re-render
      });
  }, []);


  let students = data.map((student) => {
    return (
      <StudentCard
        person={student.person}
        status={student.status}
        lookingFor={student.lookingFor}
        key={student.id.toString()}
      />
    );
  });

  return (
    <div>
      <div className="students">
        Current and past students of {courseName}
      </div>
      {/* <div className="chatInstructions">
        Click on a student name to chat with them!
      </div> */}
      {students}
      <div className="students">
        Current students avaialable as an Expert  of {courseName}
      </div>
      {/* <div className="chatInstructions">
        Click on a student name to chat with them!
      </div> */}
      {students}
    </div>
  );
}
