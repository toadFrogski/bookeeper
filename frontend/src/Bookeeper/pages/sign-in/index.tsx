import { Box, Button, Container, FormControl, InputLabel, Link as MuiLink, OutlinedInput, Paper } from "@mui/material";

import { FC, MouseEventHandler, useContext, useState } from "react";
import styles from "./styles.module.scss";
import { Link, useNavigate } from "react-router-dom";
import { ApiContext } from "../../contexts/api";
import { LoginContext } from "../../../contexts/login";
import { Password } from "../../components";

const SignIn: FC = () => {
  const { authApi } = useContext(ApiContext);
  const { setToken } = useContext(LoginContext);
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault();
    authApi
      .loginPost({ email: email, password: password })
      .then((res) => {
        if (res.data.data) {
          setToken(res.data.data);
          navigate("/");
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <Container maxWidth="sm">
      <Box component="section" className={styles.loginForm}>
        <Paper sx={{ padding: 5 }}>
          <FormControl sx={{ width: "100%" }}>
            <InputLabel htmlFor="form-element-email">Email or username</InputLabel>
            <OutlinedInput
              value={email}
              onChange={(e) => {
                setEmail(e.target.value);
              }}
              id="form-element-email"
              label="Email or username"
            />
          </FormControl>
          <Password sx={{ width: "100%", mt: 3 }} password={password} setPassword={setPassword}/>
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

export default SignIn;
