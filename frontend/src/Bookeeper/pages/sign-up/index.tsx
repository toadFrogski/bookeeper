import { Box, Button, Container, FormControl, InputLabel, OutlinedInput, Paper } from "@mui/material";
import { FC, MouseEventHandler, useContext, useReducer, useState } from "react";
import styles from "./styles.module.scss";
import { ApiContext } from "../../contexts/api";
import { LoginContext } from "../../../contexts/login";
import { Password, PasswordLine } from "../../components";

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

const SignUp: FC = () => {
  const { authApi } = useContext(ApiContext);
  const { setToken } = useContext(LoginContext);

  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const [passwordStrength, setPasswordStrength] = useState(0);
  const [passwordStatus, setPasswordStatus] = useState<PasswordStatus>(initialStates.passwordStatus);

  const handleSubmit: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault();
    authApi
      .registerPost({ email: email, password: password, username: username })
      .then((res) => {
        if (res.data.data) {
          setToken(res.data.data);
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const validatePassword = (password: string) => {
    let strength: number = 0;
    if (atLeastMinimumLength(password)) {
      strength += 35;
      setPasswordStatus((status) => ({ ...status, minimumLength: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, minimumLength: true }));
    }
    if (atLeastOneLowercaseLetter(password)) {
      strength += 5;
      setPasswordStatus((status) => ({ ...status, oneLowercaseLetter: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, oneLowercaseLetter: true }));
    }
    if (atLeastOneUppercaseLetter(password)) {
      strength += 5;
      setPasswordStatus((status) => ({ ...status, oneUppercaseLetter: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, oneUppercaseLetter: true }));
    }
    if (atLeastOneNumber(password)) {
      strength += 15;
      setPasswordStatus((status) => ({ ...status, oneNumber: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, oneNumber: true }));
    }
    if (atLeastOneSpecialChar(password)) {
      strength += 20;
      setPasswordStatus((status) => ({ ...status, oneSpecialChar: false }));
    } else {
      setPasswordStatus((status) => ({ ...status, oneSpecialChar: true }));
    }

    setPasswordStrength(strength);
  };

  return (
    <Container maxWidth="sm">
      <Box component="section" className={styles.registerForm}>
        <Paper sx={{ padding: 5 }}>
          <FormControl sx={{ width: "100%" }}>
            <InputLabel htmlFor="form-element-email">Email</InputLabel>
            <OutlinedInput
              value={email}
              onChange={(e) => {
                setEmail(e.target.value);
              }}
              id="form-element-email"
              label="Email"
            />
          </FormControl>
          <FormControl sx={{ width: "100%", mt: 2 }}>
            <InputLabel htmlFor="form-element-username">Username</InputLabel>
            <OutlinedInput
              value={username}
              onChange={(e) => {
                setUsername(e.target.value);
              }}
              id="form-element-username"
              label="Username"
            />
          </FormControl>
          <Password
            sx={{ width: "100%", mt: 2 }}
            password={password}
            setPassword={(password) => {
              setPassword(password);
              validatePassword(password)
            }}
          />
          <PasswordLine sx={{ mt: 1 }} strength={passwordStrength} />
          <Button
            sx={{ width: "100%", mt: 3, minHeight: "56px" }}
            variant="contained"
            onClick={handleSubmit}
            disabled={email.length === 0 || passwordStrength < 50}
          >
            Submit
          </Button>
        </Paper>
      </Box>
    </Container>
  );
};

export default SignUp;
