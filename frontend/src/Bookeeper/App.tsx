import { FC, useContext, useMemo } from "react";
import { SessionContext } from "../contexts/session";
import { CssBaseline, ThemeProvider, createTheme } from "@mui/material";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Home, Login } from "./pages";
import getTheme from "./assets/theme";
import { LoginProvider } from "../contexts/login";

const App: FC = () => {
  const { theme } = useContext(SessionContext);

  const mode = useMemo(() => createTheme(getTheme(theme)), [theme]);

  return (
    <LoginProvider>
      <ThemeProvider theme={mode}>
        <Router>
          <Routes>
            <Route path="*" element={<Home />} />
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
          </Routes>
        </Router>
        <CssBaseline />
      </ThemeProvider>
    </LoginProvider>
  );
};

export default App;
