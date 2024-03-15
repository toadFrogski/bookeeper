import { FC, useContext } from "react";
import styles from "./styles.module.scss";
import { Button, Container, IconButton, Typography } from "@mui/material";
import { LoginContext } from "../../../contexts/login";
import { Login, Logout, LightMode, DarkMode, AccountBox } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";
import { SessionContext } from "../../../contexts/session";

const Header: FC = () => {
  const { token, logout } = useContext(LoginContext);
  const { theme, switchTheme } = useContext(SessionContext);
  const navigate = useNavigate();

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
        <Typography display={{ xs: "none", md: "block" }} variant="h4">
          Bookeeper
        </Typography>
        <Typography display={{ xs: "block", md: "none" }} variant="h4">
          //.bk
        </Typography>

        <div className={styles.controlPanel}>
          <IconButton className={styles.modeSwitcher} onClick={switchTheme} color="inherit">
            {theme === "light" ? (
              <LightMode className={styles.switchElement} />
            ) : (
              <DarkMode className={styles.switchElement} />
            )}
          </IconButton>

          {token.token && (
            <>
              <IconButton className={styles.cpElementSM}>
                <AccountBox />
              </IconButton>
              <Button className={styles.cpElement} color="inherit" startIcon={<AccountBox />}>
                Profile
              </Button>
            </>
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
            {token.token !== "" ? "Logout" : "Login"}
          </Button>
        </div>
      </Container>
    </div>
  );
};

export default Header;
