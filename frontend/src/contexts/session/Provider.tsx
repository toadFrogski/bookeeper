import { FC, PropsWithChildren, useEffect, useState } from "react";
import Context, { Theme } from "./Context";

const getTheme = () => {
  const theme = `${window?.localStorage?.getItem("theme")}`;
  if (theme in Theme) return theme as Theme;

  const userMedia = window.matchMedia("(prefers-color-scheme: light)");
  if (userMedia.matches) return Theme.light;

  return Theme.dark;
};

export type Props = PropsWithChildren;

const Provider: FC<Props> = ({ children }: Props) => {
  const [theme, setTheme] = useState<Theme>(getTheme);

  useEffect(() => {
    document.documentElement.dataset.theme = theme;
    localStorage.setItem("theme", theme);
  }, [theme]);

  return <Context.Provider value={{ theme, setTheme }}>{children}</Context.Provider>;
};

export default Provider;
