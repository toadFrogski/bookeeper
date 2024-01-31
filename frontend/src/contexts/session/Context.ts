import { createContext } from "react";

export enum Theme {
  dark = "dark",
  light = "light",
}

type Props = {
  theme: Theme;
  setTheme: (theme: Theme) => void;
};

const initialState: Props = {
  theme: Theme.light,
  setTheme: () => {},
};

export default createContext<Props>(initialState);
