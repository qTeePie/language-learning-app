import React from "react";
import { Navbar, Nav, Container } from "react-bootstrap";
import { Button } from "components";
import styles from "./styles.module.css";

export const Header = () => {
  return (
    <Navbar bg="dark" expand="lg" className={styles.navbar} fixed="top">
      <Container>
        <Navbar.Brand href="#">Talk 2 Me</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link href="#weekly">Weekly</Nav.Link>
            <Nav.Link href="#language">Language</Nav.Link>
            <Nav.Link href="#vocabulary">Vocabulary</Nav.Link>
            <Nav.Link href="#insights">Cultural Insights</Nav.Link>
          </Nav>
          <Nav>
            <Button hoverColor="#1a1a1a">Sign In</Button>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
};
