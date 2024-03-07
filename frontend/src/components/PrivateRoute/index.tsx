import { FC, useContext } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { LoginContext } from "../../contexts/login";

const PrivateRoute: FC = () => {
  const { token } = useContext(LoginContext);
  if (token.token === "") return <Navigate to="/sign-in" />;
  return <Outlet />;
};

export default PrivateRoute;
