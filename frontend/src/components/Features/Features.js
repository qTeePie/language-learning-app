import React from "react";
import { Container, Row, Col } from "react-bootstrap";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faSmile,
  faGraduationCap,
  faLanguage,
} from "@fortawesome/free-solid-svg-icons";

const features = [
  {
    icon: faSmile,
    title: "Happy Learning",
    description:
      "Improve your language skills and enhance a new way of communication.",
  },
  {
    icon: faGraduationCap,
    title: "Improved Fluency",
    description:
      "Manage language fluency and communicate efficiently in any situation.",
  },
  {
    icon: faLanguage,
    title: "Reduced Language Barrier",
    description: "Break language barriers and connect with the world.",
  },
];

export const Features = () => {
  return (
    <div className="bg-primary text-white py-5">
      <Container>
        <Row>
          {features.map((feature, index) => (
            <Col key={index} md={4} className="text-center">
              <FontAwesomeIcon icon={feature.icon} size="3x" />
              <h4 className="my-3">{feature.title}</h4>
              <p>{feature.description}</p>
            </Col>
          ))}
        </Row>
      </Container>
    </div>
  );
};
