import React from "react";
import ReactDOM from "react-dom";
import "./theme.scss";
import { App } from "components"; // Import App component

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
