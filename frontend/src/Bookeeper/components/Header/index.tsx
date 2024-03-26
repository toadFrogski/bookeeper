import { FC, useContext } from "react";
import styles from "./styles.module.scss";
import { Button, Container, IconButton, Typography, Link as MuiLink, Stack } from "@mui/material";
import { LoginContext } from "../../../contexts/login";
import { Login, Logout, LightMode, DarkMode, AccountBox } from "@mui/icons-material";
import { Link, useNavigate } from "react-router-dom";
import { SessionContext } from "../../../contexts/session";
import urls from "../../utils/urls";
import { useTranslation } from "react-i18next";

type Props = {
  disableSign?: boolean;
};

const Header: FC<Props> = ({ disableSign }) => {
  const { token, logout } = useContext(LoginContext);
  const { theme, switchTheme } = useContext(SessionContext);
  const navigate = useNavigate();
  const [t] = useTranslation();

  const handleLogin = () => {
    if (token.token !== "") {
      logout();
    } else {
      navigate("/sign-in");
    }
  };

  return (
    <div className={styles.header}>
      <Container className={styles.container}>
        <MuiLink component={Link} to={urls.home} sx={{ textDecoration: "none", color: "inherit" }}>
          <Typography display={{ xs: "none", md: "block" }} variant="h4">
            Bookeeper
          </Typography>
          <Typography display={{ xs: "block", md: "none" }} variant="h4">
            //.bk
          </Typography>
        </MuiLink>

        <div className={styles.controlPanel}>
          <Stack>
            <IconButton className={styles.modeSwitcher} onClick={switchTheme} color="inherit">
              {theme === "light" ? (
                <LightMode className={styles.switchElement} />
              ) : (
                <DarkMode className={styles.switchElement} />
              )}
            </IconButton>
          </Stack>

          {!disableSign && (
            <>
              {token.token && (
                <MuiLink component={Link} to={urls.profile} sx={{ textDecoration: "none", color: "inherit" }}>
                  <IconButton className={styles.cpElementSM}>
                    <AccountBox />
                  </IconButton>
                  <Button className={styles.cpElement} color="inherit" startIcon={<AccountBox />}>
                    {t("common.profile")}
                  </Button>
                </MuiLink>
              )}
              <IconButton className={styles.cpElementSM} onClick={handleLogin}>
                {token.token !== "" ? <Logout /> : <Login />}
              </IconButton>
              <Button
                className={styles.cpElement}
                onClick={handleLogin}
                color="inherit"
                startIcon={token.token !== "" ? <Logout /> : <Login />}
              >
                {token.token !== "" ? t("common.logout") : t("common.login")}
              </Button>
            </>
          )}
        </div>
      </Container>
    </div>
  );
};

export default Header;
