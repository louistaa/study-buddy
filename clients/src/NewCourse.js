import React, { Component } from "react";
import PropTypes from "prop-types";
import api from "./Constants/APIEndpoints/APIEndpoints";
import Errors from "./Components/Errors/Errors"
/**
 * @class
 * @classdesc SignUp handles the sign up component
 */
class CourseForm extends Component {
  static propTypes = {
    setPage: PropTypes.func,
    setAuthToken: PropTypes.func,
  };

  constructor(props) {
    super(props);

    this.state = {
     name: "",
     professorName: "",
     departmentName: "",
     quarterName: "",
      error: "",
    };

    this.fields = [
      {
        name: "Name",
        key: "name",
      },
      {
        name: "ProfessorName",
        key: "professorName",
      },
      {
        name: "DepartmentName",
        key: "departmentName",
      },
      {
        name: "QuarterName",
        key: "quarterName",
      }
    ];
  }

  /**
   * @description setField will set the field for the provided argument
   */
  setField = (e) => {
    this.setState({ [e.target.name]: e.target.value });
  };

  /**
   * @description setError sets the error message
   */
  setError = (error) => {
    this.setState({ error });
  };

  /**
   * @description submitForm handles the form submission
   */
  submitForm = async (e) => {
      console.log("Hello I have been submitted")
    e.preventDefault();
    const {
      name,
      professorName,
      departmentName,
      quarterName
    } = this.state;
    const sendData = {
      name,
      professorName,
      departmentName,
      quarterName
    };
   
    const response = await fetch(api.base + api.handlers.classes, {
      method: "POST",
      body: JSON.stringify(sendData),
      headers: new Headers({
        "Content-Type": "application/json",
        "Authorization": this.props.authToken
      }),
    });
    if (response.status >= 300) {
      const error = await response.text();
      this.setError(error);
      return;
    }
  };

  render() {
    const values = this.state;
    const { error } = this.state;
    return (
      <>
        <div className="row justify-content-center">
          <Errors error={error} setError={this.setError} />
        </div>

        <div className="row justify-content-center">
        <form onSubmit={this.submitForm}>
            {this.fields.map(d => {
                const { key, name } = d;
                return <div key={key}>
                    <span>{name}: </span>
                    <input
                        value={values[key]}
                        name={key}
                        onChange={ this.setField}
                        
                    />
                </div>
            })}
            <input type="submit" value="Submit" />
        </form>
        </div>
      </>
    );
  }
}

CourseForm.propTypes = {
    setField: PropTypes.func.isRequired,
    submitForm: PropTypes.func.isRequired,
    values: PropTypes.shape({
        name: PropTypes.string.isRequired,
        professorName: PropTypes.string.isRequired,
        departmentName: PropTypes.string.isRequired,
        quarterName: PropTypes.string.isRequired,
    }),
    fields: PropTypes.arrayOf(PropTypes.shape({
        key: PropTypes.string,
        name: PropTypes.string
    }))
}


export default CourseForm;
