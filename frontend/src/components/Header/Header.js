import React from "react";
import { Navbar, Nav, Container, Button } from "react-bootstrap";
import styles from "./styles.module.css";

export const Header = () => {
  return (
    <Navbar bg="dark" variant="dark" expand="lg" className={styles.navbar} fixed="top">
      <Container>
        <Navbar.Brand href="#">Positive Learning</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link href="#weekly">Weekly</Nav.Link>
            <Nav.Link href="#language">Language</Nav.Link>
            <Nav.Link href="#vocabulary">Vocabulary</Nav.Link>
            <Nav.Link href="#insights">Cultural Insights</Nav.Link>
          </Nav>
          <Nav>
            <Nav.Link href="#signin">Sign In</Nav.Link>
            <Button variant="outline-success">Explore for Free</Button>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
};
