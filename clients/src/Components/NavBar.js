import Navbar from "react-bootstrap/Navbar";
import { NavLink } from "react-router-dom";

export default function NavBar() {
  return (
    <Navbar className="nav">
      <Navbar.Brand className="nav-brand">UWL Stranger Cards</Navbar.Brand>
      <Navbar.Toggle />
      <Navbar.Collapse className="justify-content-end">
        <NavLink
          to="/Perception"
          className="nav-link"
          activeClassName="category-text-perception"
        >
          Perception
        </NavLink>
        <NavLink
          to="/Connection"
          className="nav-link"
          activeClassName="category-text-connection"
        >
          Connection
        </NavLink>
        <NavLink
          to="/Reflection"
          className="nav-link"
          activeClassName="category-text-reflection"
        >
          Reflection
        </NavLink>
      </Navbar.Collapse>
    </Navbar>
  );
}
