import { FC, useContext, useMemo } from "react";
import { SessionContext } from "../contexts/session";
import { CssBaseline, ThemeProvider, createTheme } from "@mui/material";
import { Routes, Route, Navigate } from "react-router-dom";
import { Home, SignIn, SignUp } from "./pages";
import getTheme from "./assets/theme";
import "./index.scss";
import { ApiProvider } from "./contexts/api";
import { PrivateRoute } from "../components";

const App: FC = () => {
  const { theme } = useContext(SessionContext);
  const mode = useMemo(() => createTheme(getTheme(theme)), [theme]);

  return (
    <ApiProvider>
      <ThemeProvider theme={mode}>
        <Routes>
          <Route element={<PrivateRoute />}>
            <Route path="/" element={<Home />} />
          </Route>
          <Route path="/sign-in" element={<SignIn />} />
          <Route path="/sign-up" element={<SignUp />} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
        <CssBaseline />
      </ThemeProvider>
    </ApiProvider>
  );
};

export default App;
