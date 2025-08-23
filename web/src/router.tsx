import { createBrowserRouter } from "react-router-dom";
import { ConfirmationPage } from "./components/confirmation";
import App from "./App";
import LoginPage from "./pages/auth/login-page";
import RegisterPage from "./pages/auth/register-page";

const confirmationEndpoint = "/confirm/:token";
const loginEndpoint = "/login";
const registerEndpoint = "/register";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: confirmationEndpoint,
    element: <ConfirmationPage />,
  },
  {
    path: loginEndpoint,
    element: <LoginPage />,
  },
  {
    path: registerEndpoint,
    element: <RegisterPage />,
  }
]);