import { createContext } from "react";

type Props = {
  isAuthenticated: boolean;
  logout: () => void;
};

const initialState: Props = {
  isAuthenticated: false,
  logout: () => {},
};

const Context = createContext<Props>(initialState);

export default Context;
