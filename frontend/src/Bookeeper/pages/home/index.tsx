import { FC, useContext } from "react";
import { LoginContext } from "../../../contexts/login";
import { Header } from "../../components";

const Home: FC = () => {
  const { logout } = useContext(LoginContext);
  return <Header />
};

export default Home;
