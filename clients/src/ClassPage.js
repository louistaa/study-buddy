import React from "react";
import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import StudentCard from './StudentCard';
import axios from 'axios';


export default function ClassPage(props) {
  let { classID } = useParams();

  // fetch a list of the people in a course
  const [student, setStudents] = useState([]);
  const [expert, setExperts] = useState([]);
  const [course, setCourse] = useState([]);

  
  useEffect(() => {
    axios({
      "method": "GET",
      "url": "https://studybuddy-api.kaylalee.me/classes/" + classID,
      "headers": {
        "Authorization": props.authToken
      }
    })
    .then((response) => {
      setCourse(response.data);
      console.log(course)
    })
    .catch((error) => {
      console.log(error)
    })
  }, []);

  useEffect(() => {
    axios({
      "method": "GET",
      "url": "https://studybuddy-api.kaylalee.me/classes/" + classID + "/people",
      "headers": {
        "Authorization": props.authToken
      }
    })
    .then((response) => {
      setStudents(response.data);
    })
    .catch((error) => {
      console.log(error)
    })
  }, []);

  useEffect(() => {
    axios({
      "method": "GET",
      "url": "https://studybuddy-api.kaylalee.me/classes/" + classID + "/experts",
      "headers": {
        "Authorization": props.authToken
      }
    })
    .then((response) => {
      setExperts(response.data);
    })
    .catch((error) => {
      console.log(error)
    })
  }, []);


  let students = student.map((student) => {
    return (
      <StudentCard
        person={student.firstName + " " + student.lastName}
        username={student.userName}
        major={student.major}
        phonenumber={student.phoneNumber}
        email={student.email}
        id={student.id}
        key={student.id.toString()}
      />
    );
  });

  let experts = expert.map((expert) => {
    return (
      <StudentCard
        person={expert.firstName + " " + expert.lastName}
        username={expert.userName}
        major={expert.major}
        phonenumber={expert.phoneNumber}
        email={expert.email}
        id={student.id}
        key={expert.id.toString()}
      />
    );    
  });

  return (
    <div>
      <div className="courseInfo">
        <h3>Course Name: {course.name}</h3>
        <h4>Department Name: {course.departmentName}</h4>
        <h4>Professor Name: {course.professorName}</h4>
        <h4>Quarter Name: {course.quarterName}</h4>
      </div>
      <div className="students">
        Current and past students of {course.name}
      </div>
      {students}
      <div className="students">
        Students available as an Expert of {course.name}
      </div>
      {experts}
    </div>
  );
}
