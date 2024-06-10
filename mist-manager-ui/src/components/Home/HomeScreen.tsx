import React from "react";

import { Button } from "../../common";
import { useNavigate } from "react-router-dom";

const HomeScreen = (): JSX.Element => {
  const navigate = useNavigate();
  return (
    <div style={{ margin: "25px" }}>
      <Button
        onClick={() => {
          navigate("/servers");
        }}
      >
        Server Screen
      </Button>
    </div>
  );
};

export default HomeScreen;
