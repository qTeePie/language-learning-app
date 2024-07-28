import React from "react";
import { Container, Row, Col } from "react-bootstrap";

export const Footer = () => {
  return (
    <footer className="bg-dark text-white py-1">
      <Container>
        <Row>
          <Col className="text-center">
            <p>
              © {new Date().getFullYear()} Positive Learning. All rights
              reserved.
            </p>
            <p>
              <a href="#" className="text-white">
                Privacy Policy
              </a>{" "}
              |{" "}
              <a href="#" className="text-white">
                Terms of Service
              </a>
            </p>
          </Col>
        </Row>
      </Container>
    </footer>
  );
};
