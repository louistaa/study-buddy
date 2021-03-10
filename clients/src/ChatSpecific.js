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
            <p className="card-text">LOL this class is so hard</p>
          </div>
        </div>
      </div>

      <div className="col-9">
        <div className="card">
          <div className="card-body">
            <h5 className="card-title">{chatWith}</h5>
            <h6 className="card-subtitle mb-2 text-muted">11:47am</h6>
            <p className="card-text">Yaaaaaa</p>
          </div>
        </div>
      </div>
    </div>
  );
}
