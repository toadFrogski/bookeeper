import { FC, useContext, useMemo } from "react";
import { SessionContext } from "../contexts/session";
import { CssBaseline, ThemeProvider, createTheme } from "@mui/material";
import { Routes, Route, Navigate } from "react-router-dom";
import { Home, Profile, SignIn, SignUp } from "./pages";
import getTheme from "./assets/theme";
import { Header, PrivateRoute } from "./components";
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
          <Route path={urls.home} element={<Header />} />
          <Route path={urls.profile} element={<Header />} />
          <Route path={urls.signIn} element={<Header disableSign/>} />
          <Route path={urls.signUp} element={<Header disableSign/>} />
        </Routes>
        <Routes>
          <Route element={<PrivateRoute />}>
            <Route path={urls.profile} element={<Profile />} />
          </Route>
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
