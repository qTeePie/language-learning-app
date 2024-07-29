import React from "react";
import styles from "./styles.module.css";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faQuoteLeft, faQuoteRight } from '@fortawesome/free-solid-svg-icons';

export const MainSection = () => {
  return (
    <div className={styles.main}>
      <div className={styles.quoteIcon}>
        <FontAwesomeIcon icon={faQuoteLeft} size="6x" />
        <div class="space"></div>
        <FontAwesomeIcon icon={faQuoteRight} size="6x" />
      </div>
       <h1 className={styles.heading}>
        Unlock the 🌍 One Word at a Time
        </h1>
    </div>
  );
};
