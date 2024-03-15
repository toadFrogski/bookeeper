import { FC, useContext } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { LoginContext } from "../../../contexts/login";
import urls from "../../utils/urls";

const PrivateRoute: FC = () => {
  const { token } = useContext(LoginContext);
  if (token.token === "") return <Navigate to={urls.signIn} />;
  return <Outlet />;
};

export default PrivateRoute;
