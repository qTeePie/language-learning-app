import React from "react";
import { Header } from "components/organisms/Header/Header";
import { Footer } from "components/organisms/Footer/Footer";
import styles from "./styles.module.css";

export const MainLayout = ({ children }) => {
  return (
    <div className={styles.main}>
      <Header className={styles.header} />
      <main className={styles.content}>{children}</main>
      <Footer className={styles.footer} />
    </div>
  );
};
