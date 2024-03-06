import {
  Box,
  Button,
  Container,
  FormControl,
  IconButton,
  InputAdornment,
  InputLabel,
  Link as MuiLink,
  OutlinedInput,
  Paper,
} from "@mui/material";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import { FC, MouseEventHandler, useContext, useReducer, useState } from "react";
import styles from "./styles.module.scss";
import { Link } from "react-router-dom";
import { ApiContext } from "../../contexts/api";
import { PasswordLine } from "../../components";

const atLeastMinimumLength = (password: string) => new RegExp(/(?=.{8,})/).test(password);
const atLeastOneUppercaseLetter = (password: string) => new RegExp(/(?=.*?[A-Z])/).test(password);
const atLeastOneLowercaseLetter = (password: string) => new RegExp(/(?=.*?[a-z])/).test(password);
const atLeastOneNumber = (password: string) => new RegExp(/(?=.*?[0-9])/).test(password);
const atLeastOneSpecialChar = (password: string) => new RegExp(/(?=.*?[#?!@$ %^&*-])/).test(password);

type PasswordStatus = {
  minimumLength: boolean;
  oneUppercaseLetter: boolean;
  oneLowercaseLetter: boolean;
  oneNumber: boolean;
  oneSpecialChar: boolean;
};

const initialStates = {
  passwordStatus: {
    minimumLength: false,
    oneUppercaseLetter: false,
    oneLowercaseLetter: false,
    oneNumber: false,
    oneSpecialChar: false,
  },
};

const Login: FC = () => {
  const { authApi } = useContext(ApiContext);

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  // Styled states
  const [showPassword, setShowPassword] = useState(false);
  const [passwordStrength, setPasswordStrength] = useState(0);
  const [passwordStatus, setPasswordStatus] = useState<PasswordStatus>(initialStates.passwordStatus);

  const handleSubmit: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault();
    authApi
      .loginPost({ email: email, password: password })
      .then((res) => {
        console.log(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const validatePassword = (password: string) => {
    let strength: number = 0;
    if (atLeastMinimumLength(password)) {
      strength += 16;
      setPasswordStatus((status) => ({ ...status, minimumLength: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, minimumLength: true }));
    }
    if (atLeastOneLowercaseLetter(password)) {
      strength += 16;
      setPasswordStatus((status) => ({ ...status, oneLowercaseLetter: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, oneLowercaseLetter: true }));
    }
    if (atLeastOneUppercaseLetter(password)) {
      strength += 16;
      setPasswordStatus((status) => ({ ...status, oneUppercaseLetter: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, oneUppercaseLetter: true }));
    }
    if (atLeastOneNumber(password)) {
      strength += 16;
      setPasswordStatus((status) => ({ ...status, oneNumber: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, oneNumber: true }));
    }
    if (atLeastOneSpecialChar(password)) {
      strength += 16;
      setPasswordStatus((status) => ({ ...status, oneSpecialChar: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, oneSpecialChar: true }));
    }

    setPasswordStrength(strength);
  };

  return (
    <Container maxWidth="sm">
      <Box component="section" className={styles.loginForm}>
        <Paper sx={{ padding: 5 }}>
          <FormControl sx={{ width: "100%" }}>
            <InputLabel htmlFor="login-email">Email or username</InputLabel>
            <OutlinedInput
              value={email}
              onChange={(e) => {
                setEmail(e.target.value);
              }}
              id="login-email"
              label="Email or username"
            />
          </FormControl>
          <FormControl sx={{ width: "100%", mt: 3 }} variant="outlined">
            <InputLabel htmlFor="login-password">Password</InputLabel>
            <OutlinedInput
              value={password}
              onChange={(e) => {
                setPassword(e.target.value);
                validatePassword(e.target.value);
              }}
              id="login-password"
              type={showPassword ? "text" : "password"}
              endAdornment={
                <InputAdornment position="end">
                  <IconButton
                    aria-label="toggle password visibility"
                    onMouseDown={() => setShowPassword(true)}
                    onMouseUp={() => setShowPassword(false)}
                    edge="end"
                  >
                    {showPassword ? <VisibilityOff /> : <Visibility />}
                  </IconButton>
                </InputAdornment>
              }
              label="Password"
            />
          </FormControl>
          {password.length !== 0 && <PasswordLine sx={{ mt: 2 }} strength={passwordStrength} />}
          <MuiLink sx={{ mt: 2, display: "block" }} component={Link} to="/">
            Forgot password
          </MuiLink>
          <Button
            sx={{ width: "100%", mt: 3, minHeight: "56px" }}
            variant="contained"
            onClick={handleSubmit}
            disabled={email.length === 0 || password.length == 0}
          >
            Submit
          </Button>
        </Paper>
      </Box>
    </Container>
  );
};

export default Login;
