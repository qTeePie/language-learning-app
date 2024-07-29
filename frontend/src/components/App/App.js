import React from "react";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  Navigate,
} from "react-router-dom";
import { Home } from "pages";
import { WordDetail } from "pages";
import { MainLayout } from "layouts";

export const App = () => {
  return (
    <Router>
      <MainLayout>
        <Routes>
          <Route path="/" element={<Navigate to="/home" />} />
          <Route path="/home" element={<Home />} />
          <Route path="/word/:word" element={<WordDetail />} />
        </Routes>
      </MainLayout>
    </Router>
  );
};
