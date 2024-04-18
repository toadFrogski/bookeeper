import { FC, useContext, useMemo } from "react";
import { SessionContext } from "../contexts/session";
import { CssBaseline, ThemeProvider, createTheme } from "@mui/material";
import { Routes, Route, Navigate } from "react-router-dom";
import { Home, Profile, ProfileAddBook, ProfileEditBook, SignIn, SignUp } from "./pages";
import getTheme from "./assets/theme";
import { Header, Notifier, PrivateRoute } from "./components";
import { ApiProvider } from "./contexts/api";
import urls from "./utils/urls";
import "./index.scss";
import "./i18n";
import { useTranslation } from "react-i18next";
import { I18Namespace } from "./i18n";
import { NotificationProvider } from "./contexts/notification";

const App: FC = () => {
  const { theme } = useContext(SessionContext);
  const mode = useMemo(() => createTheme(getTheme(theme)), [theme]);
  const [, i18n] = useTranslation();
  i18n.setDefaultNamespace(I18Namespace.Bookeeper);

  return (
    <ApiProvider>
      <NotificationProvider>
        <ThemeProvider theme={mode}>
          <Routes>
            <Route path={urls.signIn} element={<Header disableSign />} />
            <Route path={urls.signUp} element={<Header disableSign />} />
            <Route path="*" element={<Header />} />
          </Routes>
          <Routes>
            <Route element={<PrivateRoute />}>
              <Route path={urls.profile} element={<Profile />} />
              <Route path={urls.profileAddBook} element={<ProfileAddBook />} />
              <Route path={urls.profileEditBook} element={<ProfileEditBook />} />
            </Route>
            <Route path={urls.home} element={<Home />} />
            <Route path={urls.signIn} element={<SignIn />} />
            <Route path={urls.signUp} element={<SignUp />} />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
          <Notifier />
          <CssBaseline />
        </ThemeProvider>
      </NotificationProvider>
    </ApiProvider>
  );
};

export default App;
