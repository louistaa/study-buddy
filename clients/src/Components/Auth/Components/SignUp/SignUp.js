import React, { Component } from "react";
import PropTypes from "prop-types";
import SignForm from "../SignForm/SignForm";
import api from "../../../../Constants/APIEndpoints/APIEndpoints";
import Errors from "../../../Errors/Errors";
import PageTypes from "../../../../Constants/PageTypes/PageTypes";

/**
 * @class
 * @classdesc SignUp handles the sign up component
 */
class SignUp extends Component {
  static propTypes = {
    setPage: PropTypes.func,
    setAuthToken: PropTypes.func,
  };

  constructor(props) {
    super(props);

    this.state = {
      email: "",
      userName: "",
      firstName: "",
      lastName: "",
      major: "",
      password: "",
      passwordConf: "",
      error: "",
    };

    this.fields = [
      {
        name: "Email",
        key: "email",
      },
      {
        name: "Phone Number",
        key: "phoneNumber",
      },
      {
        name: "Username",
        key: "userName",
      },
      {
        name: "First name",
        key: "firstName",
      },
      {
        name: "Last name",
        key: "lastName",
      },
      {
        name: "Major",
        key: "major",
      },
      {
        name: "Password",
        key: "password",
      },
      {
        name: "Password Confirmation",
        key: "passwordConf",
      },
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
    e.preventDefault();
    const {
      email,
      phoneNumber,
      userName,
      firstName,
      lastName,
      major,
      password,
      passwordConf,
    } = this.state;
    const sendData = {
      email,
      phoneNumber,
      userName,
      firstName,
      lastName,
      major,
      password,
      passwordConf,
    };
    const response = await fetch(api.base + api.handlers.users, {
      method: "POST",
      body: JSON.stringify(sendData),
      headers: new Headers({
        "Content-Type": "application/json",
      }),
    });
    if (response.status >= 300) {
      const error = await response.text();
      this.setError(error);
      return;
    }
    const authToken = response.headers.get("Authorization");
    localStorage.setItem("Authorization", authToken);
    this.setError("");
    this.props.setAuthToken(authToken);
    const user = await response.json();
    this.props.setUser(user);
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
          <SignForm
            setField={this.setField}
            submitForm={this.submitForm}
            values={values}
            fields={this.fields}
          />
        </div>
        <div className="row justify-content-center">
          <button onClick={(e) => this.props.setPage(e, PageTypes.signIn)}>
            Sign in instead
          </button>
        </div>
      </>
    );
  }
}

export default SignUp;
