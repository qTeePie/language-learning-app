import React from "react";
import styles from "./styles.module.css";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faQuoteLeft, faQuoteRight } from "@fortawesome/free-solid-svg-icons";

export const Slideshow = ({ children }) => {
  return (
    <div className={styles.main}>
      <FontAwesomeIcon icon={faQuoteLeft} color="#D6A3FF" size="6x" />
      <div class="space"></div>
      <FontAwesomeIcon icon={faQuoteRight} color="#D6A3FF" size="6x" />

      <h1 className={styles.heading}>Unlock the 🌍 One Word at a Time</h1>
    </div>
  );
};
