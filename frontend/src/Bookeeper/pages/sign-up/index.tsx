import { Box, Button, Container, FormControl, Paper, TextField, Link as MuiLink } from "@mui/material";
import { FC, MouseEventHandler, useContext, useState } from "react";
import styles from "./styles.module.scss";
import { ApiContext } from "../../contexts/api";
import { LoginContext } from "../../../contexts/login";
import { CommonErrorBox, Password, PasswordLine } from "../../components";
import { Link, useNavigate } from "react-router-dom";
import { useStateWithError } from "../../../utils/hooks";
import { isAxiosError } from "axios";
import { NamedValidationErrors } from "../../../services/api";
import urls from "../../utils/urls";
import { useTranslation } from "react-i18next";

const atLeastMinimumLength = (password: string) => new RegExp(/(?=.{8,})/).test(password);
const atLeastOneUppercaseLetter = (password: string) => new RegExp(/(?=.*?[A-Z])/).test(password);
const atLeastOneLowercaseLetter = (password: string) => new RegExp(/(?=.*?[a-z])/).test(password);
const atLeastOneNumber = (password: string) => new RegExp(/(?=.*?[0-9])/).test(password);
const atLeastOneSpecialChar = (password: string) => new RegExp(/(?=.*?[#?!@$ %^&*-])/).test(password);

const SignUp: FC = () => {
  const { authApi } = useContext(ApiContext);
  const { setToken } = useContext(LoginContext);
  const navigate = useNavigate();
  const [t] = useTranslation();

  const email = useStateWithError("");
  const username = useStateWithError("");
  const password = useStateWithError("");
  const [commonError, setCommonError] = useState("");

  const [passwordStrength, setPasswordStrength] = useState(0);

  const handleSubmit: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault();
    authApi
      .registerPost({ email: email.value, password: password.value, username: username.value })
      .then((res) => {
        if (res.data.data) {
          setToken(res.data.data);
          navigate("/");
        }
      })
      .catch((err) => {
        if (isAxiosError(err)) {
          console.log(err);
          if (err.response && err.response.data.response_message === "INVALID_REQUEST") {
            const data: NamedValidationErrors = err.response.data.data;
            switch (data.name) {
              case "email":
                data.errors?.forEach((error) => {
                  if (error.type === "invalid_email") email.setError("error.invalidEmail");
                  if (error.type === "existed_email") email.setError("error.existedEmail");
                });
                break;
              case "username":
                data.errors?.forEach((error) => {
                  if (error.type === "existed_username") email.setError("error.existedUsername");
                });
                break;
              default:
                setCommonError("error.unexpectedError");
                break;
            }
          } else setCommonError("error.unexpectedError");
        } else setCommonError("error.unexpectedError");
      });
  };

  const validatePassword = (password: string) => {
    let strength: number = 0;

    if (atLeastMinimumLength(password)) strength += 40;
    if (atLeastOneLowercaseLetter(password)) strength += 5;
    if (atLeastOneUppercaseLetter(password)) strength += 5;
    if (atLeastOneNumber(password)) strength += 5;
    if (atLeastOneSpecialChar(password)) strength += 5;

    setPasswordStrength(strength);
  };

  return (
    <Container maxWidth="sm">
      <Box component="section" className={styles.registerForm}>
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
              label={t("common.email")}
            />
          </FormControl>
          <FormControl sx={{ width: "100%", mt: 2 }}>
            <TextField
              error={username.error != ""}
              variant="outlined"
              value={username.value}
              helperText={username.error}
              onChange={(e) => {
                username.setValue(e.target.value);
              }}
              label={t("common.username")}
            />
          </FormControl>
          <Password
            sx={{ width: "100%", mt: 2 }}
            password={password.value}
            setPassword={(value) => {
              validatePassword(value);
              password.setValue(value);
            }}
            error={password.error}
          />
          <PasswordLine
            sx={{ mt: 1 }}
            strength={passwordStrength}
            limits={[10, 50, 60]}
            showStatus={password.value.length > 2}
          />
          <MuiLink aria-disabled sx={{ mt: 2, display: "block" }} component={Link} to={urls.signIn}>
            {t("signUp.signInLink")}
          </MuiLink>
          <Button
            sx={{ width: "100%", mt: 3, minHeight: "56px" }}
            variant="contained"
            onClick={handleSubmit}
            disabled={email.value.length === 0 || username.value.length === 0 || passwordStrength < 50}
          >
            {t("signUp.signUp")}
          </Button>
        </Paper>
      </Box>
    </Container>
  );
};

export default SignUp;
