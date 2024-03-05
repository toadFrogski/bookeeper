import {
  Box,
  Button,
  Container,
  FormControl,
  IconButton,
  InputAdornment,
  InputLabel,
  OutlinedInput,
} from "@mui/material";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import { FC, useState } from "react";
import styles from "./styles.module.scss";

const Login: FC = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  return (
    <Container maxWidth="sm">
      <Box component="section" className={styles.loginForm}>
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
        <Button sx={{ width: "100%", mt: 3, minHeight: "56px" }} variant="outlined">
          Submit
        </Button>
      </Box>
    </Container>
  );
};

export default Login;
