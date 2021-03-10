import React from 'react';
import ReactDOM from 'react-dom';
import './Styles/App.css';
import App from './App';

import { BrowserRouter } from "react-router-dom"; // npm install react-router-dom
import 'bootstrap/dist/css/bootstrap.min.css'; // npm install react-bootstrap bootstrap

ReactDOM.render(<BrowserRouter><App /></BrowserRouter>, document.getElementById('root'));
