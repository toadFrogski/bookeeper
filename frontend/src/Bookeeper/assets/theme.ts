import { ThemeOptions, alpha } from "@mui/material";

const light = {
  primaryBase: "#43a047",
  secondaryBase: "#7cb342",
};

const dark = {
  primaryBase: "#43a047",
  secondaryBase: "#7cb342",
};

const darkThemeOptions: ThemeOptions = {
  palette: {
    mode: "dark",
    primary: {
      main: alpha(light.primaryBase, 0.7),
      light: alpha(light.primaryBase, 0.9),
      dark: alpha(light.primaryBase, 0.5),
    },
    secondary: {
      main: alpha(light.secondaryBase, 0.7),
      light: alpha(light.secondaryBase, 0.9),
      dark: alpha(light.secondaryBase, 0.5),
    },
  },
  typography: {
    fontFamily: "Roboto",
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: `
      input:-webkit-autofill,
      input:-webkit-autofill:hover,
      input:-webkit-autofill:focus,
      input:-webkit-autofill:active {
        -webkit-box-shadow: none !important;
      }`,
    },
  },
};

const lightThemeOptions: ThemeOptions = {
  palette: {
    mode: "light",
    primary: {
      main: alpha(dark.primaryBase, 0.7),
      light: alpha(dark.primaryBase, 0.9),
      dark: alpha(dark.primaryBase, 0.5),
    },
    secondary: {
      main: alpha(dark.secondaryBase, 0.7),
      light: alpha(dark.secondaryBase, 0.9),
      dark: alpha(dark.secondaryBase, 0.5),
    },
  },
  typography: {
    fontFamily: "Roboto",
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: `
      input:-webkit-autofill,
      input:-webkit-autofill:hover,
      input:-webkit-autofill:focus,
      input:-webkit-autofill:active {
        -webkit-box-shadow: none !important;
      }`,
    },
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
