import { createBrowserRouter } from "react-router-dom";
import { ConfirmationPage } from "./components/confirmation";
import App from "./App";
import LoginPage from "./app/auth/login-page";
import RegisterPage from "./app/auth/register-page";
import {
  confirmationEndpoint,
  forgotPasswordEndpoint,
  loginEndpoint,
  registerEndpoint,
} from "./utils/endpoints";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "*",
    element: <div>404 Not Found</div>,
  },
  {
    path: loginEndpoint,
    element: <LoginPage />,
  },
  {
    path: registerEndpoint,
    element: <RegisterPage />,
  },
  {
    path: confirmationEndpoint,
    element: <ConfirmationPage />,
  },
  {
    path: forgotPasswordEndpoint,
    element: <div>Forgot Password Page</div>,
  },
]);
