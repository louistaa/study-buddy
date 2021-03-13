import React, { useState, useEffect } from "react";
import CourseCard from "./CourseCard";
import axios from 'axios';

export default function Courses(props) {
  // fetch a list of the courses
  const [classes, setClasses] = useState([]);



  useEffect(() => {
    // axios.get("https://studybuddy-api.")
    // fetch("./classes.json")
    //   .then((res) => res.json())
    //   .then((data) => {
    //     setData(data.classes); //change the state and re-render
    //   });

    axios({
      "method": "GET",
      "url": "https://studybuddy-api.kaylalee.me/classes",
      "headers": {
        "Authorization": props.authToken
      }
    })
    .then((response) => {
      setClasses(response.data);
    })
    .catch((error) => {
      console.log(error)
    })
  }, []);

  let courses;
  // console.log(classes)
  if (classes != null) {
     courses = classes.map((course) => {
      return (
        <CourseCard
          course={course.name}
          department={course.departmentName}
          professor={course.professorName}
          quarter={course.quarterName}
          courseID={course.id}
          key={course.id.toString()}
        />
      );
    });
  }

  
  

  return (
    <div>
      <div className="row">
        <div className="col select">
          Select a class below to find classmates to study with:
        </div>
      </div>
      <div className="row justify-content-center">{courses}</div>
    </div>
  );
}
