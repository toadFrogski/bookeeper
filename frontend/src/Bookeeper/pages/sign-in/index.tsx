import {
  Box,
  Button,
  Container,
  FormControl,
  Paper,
  TextField,
  Typography,
  useTheme,
  Link as MuiLink,
} from "@mui/material";

import { FC, MouseEventHandler, useContext, useState } from "react";
import styles from "./styles.module.scss";
import { Link, useNavigate } from "react-router-dom";
import { ApiContext } from "../../contexts/api";
import { LoginContext } from "../../../contexts/login";
import { Password } from "../../components";
import { isAxiosError } from "axios";
import { useStateWithError } from "../../../utils/hooks";
import { ErrorOutline } from "@mui/icons-material";
import { AnyResponse } from "../../../services/api";

const SignIn: FC = () => {
  const { authApi } = useContext(ApiContext);
  const { setToken } = useContext(LoginContext);
  const { palette } = useTheme();
  const navigate = useNavigate();

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
                email.setError("Error.userNotFound");
                break;
              case "INCORRECT_PASSWORD":
                password.setError("Error.incorrectPassword");
                break;
              default:
                setCommonError("Error.unexpectedError");
                break;
            }
          } else setCommonError("Error.unexpectedError");
        } else setCommonError("Error.unexpectedError");
      });
  };

  return (
    <Container maxWidth="sm">
      <Box component="section" className={styles.loginForm} >
        <Paper sx={{ padding: 5 }} elevation={1}>
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
                label="Email or username"
              />
            </FormControl>
            <Password
              sx={{ width: "100%", mt: 3 }}
              password={password.value}
              setPassword={(value) => password.setValue(value)}
              error={password.error}
            />
            <MuiLink aria-disabled sx={{ mt: 2, display: "block" }} component={Link} to="/">
              Forgot password
            </MuiLink>
            <Button
              sx={{ width: "100%", mt: 3, minHeight: "56px" }}
              variant="contained"
              onClick={handleSubmit}
              disabled={email.value.length === 0 || password.value.length == 0}
            >
              Submit
            </Button>
        </Paper>
      </Box>
    </Container>
  );
};

export default SignIn;
