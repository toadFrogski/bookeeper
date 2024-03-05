import { ThemeOptions } from "@mui/material";

const darkThemeOptions: ThemeOptions = {
  palette: {
    mode: "dark",
    primary: {
      main: "#43a047",
    },
    secondary: {
      main: "#7cb342",
    },
  },
  typography: {
    fontFamily: "Roboto",
  },
};

const lightThemeOptions: ThemeOptions = {
  palette: {
    mode: "light",
    primary: {
      main: "#43a047",
    },
    secondary: {
      main: "#7cb342",
    },
  },
  typography: {
    fontFamily: "Roboto",
  },
};

const getTheme = (theme: string) => {
  switch (theme) {
    case "light":
      return lightThemeOptions;
    case "dark":
      return darkThemeOptions;
  }
};

export default getTheme;
