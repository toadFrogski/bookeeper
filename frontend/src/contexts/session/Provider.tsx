import { FC, PropsWithChildren, useEffect, useState } from "react";
import Context, { Theme } from "./Context";

const getTheme = () => {
  const theme = `${window?.localStorage?.getItem("theme")}`;
  if (theme === "light" || theme === "dark") return theme as Theme;

  const userMedia = window.matchMedia("(prefers-color-scheme: light)");
  if (userMedia.matches) return "light";

  return "dark";
};

export type Props = PropsWithChildren;

const Provider: FC<Props> = ({ children }: Props) => {
  const [theme, setTheme] = useState<Theme>(getTheme);

  const switchTheme = () => {
    setTheme((prevTheme) => (prevTheme === "dark" ? "light" : "dark"));
  };

  useEffect(() => {
    document.documentElement.dataset.theme = theme;
    localStorage.setItem("theme", theme);
  }, [theme]);

  return <Context.Provider value={{ theme, switchTheme }}>{children}</Context.Provider>;
};

export default Provider;
