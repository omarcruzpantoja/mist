import { RouterProvider, createBrowserRouter } from "react-router-dom";
import "./App.scss";
import { Base } from "./shared";
import { HomeScreen, ServerScreen, ChannelScreen } from "./components";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Base />,
    children: [
      {
        path: "",
        element: <HomeScreen />,
      },
      {
        path: "/servers",
        element: <ServerScreen />,
      },
      {
        path: "/channels",
        element: <ChannelScreen />,
      },
    ],
  },
]);

function App() {
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
