import { BrowserRouter as Router } from "react-router-dom";
import { Header, MainSection, Features, Footer } from "components";
import styles from "./styles.module.css";

export const App = () => {
  return (
    <Router>
      <div className={styles.main}>
        <Header />
        <main className={`${styles.content}`}>
          <MainSection />
          <Features />
        </main>
        <Footer />
      </div>
    </Router>
  );
};
