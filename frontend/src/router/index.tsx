import { createBrowserRouter } from "react-router-dom";
import MainLayout from "../layouts/MainLayout";
import HomePage from "../pages/HomePage";
import PlacePage from "../pages/PlacePage";
// import CreatePlacePage from "../pages/CreatePlacePage";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <MainLayout />,
    children: [
      { path: "/", element: <HomePage /> },
      { path: "/places", element: <PlacePage /> },
      // { path: "/places/new", element: <CreatePlacePage /> },
    ],
  },
]);