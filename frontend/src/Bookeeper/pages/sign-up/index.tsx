import { Box, Button, Container, FormControl, Paper, TextField, Typography, useTheme } from "@mui/material";
import { FC, MouseEventHandler, useContext, useState } from "react";
import styles from "./styles.module.scss";
import { ApiContext } from "../../contexts/api";
import { LoginContext } from "../../../contexts/login";
import { Password, PasswordLine } from "../../components";
import { useNavigate } from "react-router-dom";
import { useStateWithError } from "../../../utils/hooks";
import { isAxiosError } from "axios";
import { NamedValidationErrors } from "../../../services/api";
import { ErrorOutline } from "@mui/icons-material";

const atLeastMinimumLength = (password: string) => new RegExp(/(?=.{8,})/).test(password);
const atLeastOneUppercaseLetter = (password: string) => new RegExp(/(?=.*?[A-Z])/).test(password);
const atLeastOneLowercaseLetter = (password: string) => new RegExp(/(?=.*?[a-z])/).test(password);
const atLeastOneNumber = (password: string) => new RegExp(/(?=.*?[0-9])/).test(password);
const atLeastOneSpecialChar = (password: string) => new RegExp(/(?=.*?[#?!@$ %^&*-])/).test(password);

const SignUp: FC = () => {
  const { authApi } = useContext(ApiContext);
  const { setToken } = useContext(LoginContext);
  const { palette } = useTheme();
  const navigate = useNavigate();

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
                  if (error.type === "invalid_email") email.setError("Error.invalidEmail");
                  if (error.type === "existed_email") email.setError("Error.existedEmail");
                });
                break;
              case "username":
                data.errors?.forEach((error) => {
                  if (error.type === "existed_username") email.setError("Error.existedUsername");
                });
                break;
              default:
                setCommonError("Error.unexpectedError");
                break;
            }
          } else setCommonError("Error.unexpectedError");
        } else setCommonError("Error.unexpectedError");
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
        <Paper sx={{ padding: 5 }}>
          {commonError != "" && (
            <Box sx={{ mb: 5, display: "flex", color: palette.error.main }}>
              <ErrorOutline />
              <Typography sx={{ ml: 1 }} variant="inherit">
                {commonError}
              </Typography>
            </Box>
          )}
          <FormControl sx={{ width: "100%" }}>
            <TextField
              error={email.error != ""}
              variant="outlined"
              value={email.value}
              helperText={email.error}
              onChange={(e) => {
                email.setValue(e.target.value);
              }}
              label="Email"
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
              label="Username"
            />
          </FormControl>
          <Password
            sx={{ width: "100%", mt: 2 }}
            password={password.value}
            setPassword={(password) => {
              validatePassword(password);
              password.setPassword(password);
            }}
            error={password.error}
          />
          <PasswordLine
            sx={{ mt: 1 }}
            strength={passwordStrength}
            limits={[10, 50, 60]}
            showStatus={password.value.length > 2}
          />
          <Button
            sx={{ width: "100%", mt: 3, minHeight: "56px" }}
            variant="contained"
            onClick={handleSubmit}
            disabled={email.value.length === 0 || username.value.length === 0 || passwordStrength < 50}
          >
            Submit
          </Button>
        </Paper>
      </Box>
    </Container>
  );
};

export default SignUp;
