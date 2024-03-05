import { createContext } from "react";

export type Theme = "dark" | "light";

type Props = {
  theme: Theme;
  switchTheme: () => void;
};

const initialState: Props = {
  theme: "light",
  switchTheme: () => {},
};

export default createContext<Props>(initialState);
