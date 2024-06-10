import React from "react";
import { useNavigate } from "react-router-dom";

import { Button } from "../../common";

const Header = (): JSX.Element => {
  const navigate = useNavigate();
  return (
    <div className="base-header-container">
      <div className="base-header-content">
        <div>
          <Button
            onClick={() => {
              navigate("/");
            }}
          >
            HOME
          </Button>
        </div>
        <div className="base-header-right-content">
          <span>Screens(dropdown)</span>
          <span>User Icon</span>
        </div>
      </div>
    </div>
  );
};

export default Header;
