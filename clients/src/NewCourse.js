import React, { useState, useEffect } from "react";

export default function NewCourse(props) {
  return (
    <form>
      <div class="form-group">
        <label for="exampleInputEmail1">Department Abbreviation</label>
        <input
          type="text"
          class="form-control"
          id="department"
          placeholder="INFO"
        />
        <small id="emailHelp" class="form-text text-muted">
        </small>
      </div>
      <div class="form-group">
        <label for="exampleInputPassword1">Course Number</label>
        <input
          type="text"
          class="form-control"
          id="courseNumber"
          placeholder="441"
        />
      </div>
      <div class="form-group">
        <label for="exampleInputPassword1">Quarter and Year</label>
        <input
          type="text"
          class="form-control"
          id="quarter"
          placeholder="Spring 2021"
        />
      </div>
      <button type="submit" class="btn btn-primary">
        Submit
      </button>
    </form>
  );
}
