import React from 'react'
import { useParams } from "react-router-dom";

export default function ChatSpecific(props) {
  const urlParams = useParams();

  // grab the chat history from URL Params
  let chatWith = urlParams.person;

  // future GET request

  return (
    <div>
      <div className="students">Your chat with {chatWith}</div>
      <div className="col-9">
        <div className="card">
          <div className="card-body">
            <h5 className="card-title">Louis Ta</h5>
            <h6 className="card-subtitle mb-2 text-muted">11:45am</h6>
            <p className="card-text">Hey! I am looking for classmates to do homeworks with. Can we work together?</p>
          </div>
        </div>
      </div>

      <div className="col-9">
        <div className="card">
          <div className="card-body">
            <h5 className="card-title">{chatWith}</h5>
            <h6 className="card-subtitle mb-2 text-muted">11:47am</h6>
            <p className="card-text">Hi Louis! Sure!</p>
          </div>
        </div>
      </div>

      <input class="form-control" type="text" placeholder="enter message"/>
    </div>
  );
}
