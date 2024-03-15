import { FC, useContext, useMemo } from "react";
import { SessionContext } from "../contexts/session";
import { CssBaseline, ThemeProvider, createTheme } from "@mui/material";
import { Routes, Route, Navigate } from "react-router-dom";
import { Home, SignIn, SignUp } from "./pages";
import getTheme from "./assets/theme";
import { PrivateRoute } from "./components";
import { ApiProvider } from "./contexts/api";
import urls from "./utils/urls";
import "./index.scss";

const App: FC = () => {
  const { theme } = useContext(SessionContext);
  const mode = useMemo(() => createTheme(getTheme(theme)), [theme]);

  return (
    <ApiProvider>
      <ThemeProvider theme={mode}>
        <Routes>
          <Route element={<PrivateRoute />}></Route>
          <Route path={urls.home} element={<Home />} />
          <Route path={urls.signIn} element={<SignIn />} />
          <Route path={urls.signUp} element={<SignUp />} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
        <CssBaseline />
      </ThemeProvider>
    </ApiProvider>
  );
};

export default App;
