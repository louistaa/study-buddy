import React, { useState, useEffect } from "react";
import CourseCard from "./CourseCard";

export default function Courses() {
  // fetch a list of the courses
  const [data, setData] = useState([]);

  useEffect(() => {
    fetch("./classes.json")
      .then((res) => res.json())
      .then((data) => {
        setData(data.classes); //change the state and re-render
      });
  }, []);

  let courses = data.map((course) => {
    return (
      <CourseCard
        course={course.course}
        numStudents={course.numStudents}
        description={course.description}
        key={course.id.toString()}
      />
    );
  });

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
