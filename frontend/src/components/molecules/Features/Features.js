import React from "react";
import { Container, Row, Col } from "react-bootstrap";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faSmile,
  faGraduationCap,
  faLanguage,
} from "@fortawesome/free-solid-svg-icons";
import styles from "./styles.module.css";

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
    <div className={styles.features}>
      <Container>
        <Row>
          {features.map((feature, index) => (
            <Col key={index} md={4} className="text-center">
              <FontAwesomeIcon
                icon={feature.icon}
                size="3x"
                className={styles["feature-icon"]}
              />
              <h4 className={styles["feature-title"]}>{feature.title}</h4>
              <p className={styles["feature-description"]}>
                {feature.description}
              </p>
            </Col>
          ))}
        </Row>
      </Container>
    </div>
  );
};
