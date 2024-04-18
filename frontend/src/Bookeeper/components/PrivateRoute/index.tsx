import { FC, useContext } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { LoginContext } from "../../../contexts/login";
import urls from "../../utils/urls";

type Props = {
  redirect?: string;
}

const PrivateRoute: FC<Props> = ({redirect}) => {
  const { token } = useContext(LoginContext);
  if (token.token === "") return <Navigate to={redirect ?? urls.signIn} />;
  return <Outlet />;
};

export default PrivateRoute;
