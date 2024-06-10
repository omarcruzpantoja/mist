import React, { useEffect } from "react";

import BaseService from "../../api/BaseService";
import { Server } from "../../types/common";
import { Accordion, Button } from "../../common";
import { useNavigate } from "react-router-dom";

interface ServerLinkProps {
  server: Server;
}

const ServerLink = ({ server }: ServerLinkProps): JSX.Element => {
  const navigate = useNavigate();

  return (
    <Button
      onClick={() => {
        navigate(`/servers/${server.id}/channels`);
      }}
    >
      Server: {server.name}
    </Button>
  );
};

const ServerScreen = (): JSX.Element => {
  const [servers, setServers] = React.useState<Server[]>([]);

  useEffect(() => {
    const load = async () => {
      const response = await new BaseService("http://localhost:3010").get<
        Server[]
      >("/servers");
      setServers(response);
    };

    load();
  }, []);

  return (
    <div style={{ margin: "25px" }}>
      {servers.map((server) => {
        return <ServerLink server={server} />;
      })}
      <Accordion header={<div>test</div>}>hi</Accordion>
    </div>
  );
};

export default ServerScreen;
