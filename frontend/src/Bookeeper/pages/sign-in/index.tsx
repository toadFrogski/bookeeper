import { Box, Button, Container, FormControl, Paper, TextField, Link as MuiLink } from "@mui/material";

import { FC, MouseEventHandler, useContext, useState } from "react";
import styles from "./styles.module.scss";
import { Link, useNavigate } from "react-router-dom";
import { ApiContext } from "../../contexts/api";
import { LoginContext } from "../../../contexts/login";
import { CommonErrorBox, Password } from "../../components";
import { isAxiosError } from "axios";
import { useStateWithError } from "../../../utils/hooks";
import { AnyResponse } from "../../../services/api";
import urls from "../../utils/urls";
import { useTranslation } from "react-i18next";

const SignIn: FC = () => {
  const { authApi } = useContext(ApiContext);
  const { setToken } = useContext(LoginContext);
  const navigate = useNavigate();
  const [t] = useTranslation();

  const email = useStateWithError("");
  const password = useStateWithError("");
  const [commonError, setCommonError] = useState("");

  const handleSubmit: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault();
    email.setError("");
    password.setError("");
    setCommonError("");
    authApi
      .loginPost({ email: email.value, password: password.value })
      .then((res) => {
        if (res.data.data) {
          setToken(res.data.data);
          navigate("/");
        }
      })
      .catch((err) => {
        if (isAxiosError(err)) {
          if (err.response && err.response.data) {
            const data = err.response.data as AnyResponse;
            switch (data.response_message) {
              case "USER_NOT_FOUND":
                email.setError("error.userNotFound");
                break;
              case "INCORRECT_PASSWORD":
                password.setError("error.incorrectPassword");
                break;
              default:
                setCommonError("error.unexpectedError");
                break;
            }
          } else setCommonError("error.unexpectedError");
        } else setCommonError("error.unexpectedError");
      });
  };

  return (
    <Container maxWidth="sm">
      <Box component="section" className={styles.loginForm}>
        <Paper sx={{ padding: 5 }} elevation={3}>
          {commonError != "" && <CommonErrorBox message={t(commonError)} />}
          <FormControl sx={{ width: "100%" }}>
            <TextField
              error={email.error != ""}
              variant="outlined"
              value={email.value}
              helperText={email.error}
              onChange={(e) => {
                email.setValue(e.target.value);
              }}
              label={t("signIn.email")}
            />
          </FormControl>
          <Password
            sx={{ width: "100%", mt: 3 }}
            password={password.value}
            setPassword={(value) => password.setValue(value)}
            error={password.error}
          />
          <MuiLink aria-disabled sx={{ mt: 2, display: "block" }} component={Link} to={urls.signUp}>
            {t("signIn.signUpLink")}
          </MuiLink>
          <Button
            sx={{ width: "100%", mt: 3, minHeight: "56px" }}
            variant="contained"
            onClick={handleSubmit}
            disabled={email.value.length === 0 || password.value.length == 0}
          >
            {t("signIn.signIn")}
          </Button>
        </Paper>
      </Box>
    </Container>
  );
};

export default SignIn;
