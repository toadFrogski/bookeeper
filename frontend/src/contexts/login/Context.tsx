import { createContext } from "react";
import { Auth } from "../../services/api";

type Props = {
  token: Auth;
  setToken: (token: Auth) => void;
  logout: () => void;
};

const initialState: Props = {
  token: {token: ""},
  setToken: () => {},
  logout: () => {},
};

const Context = createContext<Props>(initialState);

export default Context;
