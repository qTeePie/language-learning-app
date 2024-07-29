// src/pages/Home/Home.js
import React from "react";
import { Features } from "components/molecules/Features";
import styles from "./styles.module.css";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faQuoteLeft, faQuoteRight } from "@fortawesome/free-solid-svg-icons";

export const Home = () => {
  return (
    <div className={styles.main}>
      <div className={styles.banner}>
        <div className={styles.icons}>
          <FontAwesomeIcon icon={faQuoteLeft} color="#D8B4FE" size="6x" />
          <FontAwesomeIcon icon={faQuoteRight} color="#D8B4FE" size="6x" />
        </div>
        <h1 className={styles.heading}>Unlock the 🌍 One Word at a Time</h1>
      </div>
      <Features />
    </div>
  );
};
