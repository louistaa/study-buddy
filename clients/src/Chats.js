import { useState, useEffect } from "react";
import { Redirect } from 'react-router-dom';

export default function Chats(props) {
  // fetch a list of the people in a course
  const [data, setData] = useState([]);
  const [redirectTo, setRedirectTo] = useState(undefined)

  useEffect(() => {
    fetch("./classSpecificPeople.json")
      .then((res) => res.json())
      .then((data) => {
        setData(data.people); //change the state and re-render
      });
  }, []);

  const handleClick = (name) => {
    setRedirectTo(name)
  }

  if (redirectTo) {
    return <Redirect push to={'/chats/' + redirectTo} />
  }

  let studentNames = data.map((student) => {
    return (
      <div className="col-9">
        <div className="card" onClick={() => handleClick(student.person)}>
          <div className="card-body">
            <h5 className="card-title studentName">{student.person}</h5>
          </div>
        </div>
      </div>
    );
  });

  return (
    <div>
      <div className="students">Your chat history</div>
      {studentNames}
    </div>
  );
}
