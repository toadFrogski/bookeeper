import { FC, useContext } from "react";
import { LoginContext } from "../../../contexts/login";

const Home: FC = () => {
  const { logout } = useContext(LoginContext);
  return <span onClick={() => logout()}> Logout</span>;
};

export default Home;
