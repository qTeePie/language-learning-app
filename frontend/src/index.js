import React from "react";
import ReactDOM from "react-dom";
import "./index.css"; // Import global styles from styles folder
import { App } from "components"; // Import App component
import "bootstrap/dist/css/bootstrap.min.css"; // Import Bootstrap CSS

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
