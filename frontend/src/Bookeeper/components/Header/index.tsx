import { FC, useContext } from "react";
import styles from "./styles.module.scss";
import { Button, Container, Typography, useTheme } from "@mui/material";
import { LoginContext } from "../../../contexts/login";

const Header: FC = () => {
  const { palette } = useTheme();
  const { token, logout } = useContext(LoginContext);

  return (
    <div className={styles.header}>
      <Container className={styles.container}>
        <Typography variant="h4">Bookeeper</Typography>
        <div className={styles.controlPanel}>
          <Button color="inherit">{token !== "" ? "Logout" : "Login"}</Button>
        </div>
      </Container>
    </div>
  );
};

export default Header;
