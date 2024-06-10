import React from "react";
import { Outlet } from "react-router-dom";

import Header from "./Header";
import Footer from "./Footer";

import "./_base.scss";
const Base = (): JSX.Element => {
  return (
    <div className="base-container">
      <Header />
      <Outlet />
      <Footer />
    </div>
  );
};

export default Base;
