// src/components/atoms/Button/Button.js
import React from "react";
import styles from "./styles.module.css";

export const Button = ({ children, hoverColor, ...props }) => {
  const style = {
    "--hover-color": hoverColor || "#9b0ccd", // Use provided hover color or default
  };

  return (
    <button className={styles.button} style={style} {...props}>
      {children}
    </button>
  );
};
