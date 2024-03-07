import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import { FormControl, FormControlProps, IconButton, InputAdornment, InputLabel, OutlinedInput } from "@mui/material";
import { FC, useState } from "react";

type Props = FormControlProps & {
  password: string;
  setPassword: (password: string) => void;
};

const Password: FC<Props> = ({ password, setPassword, ...props }) => {
  const [showPassword, setShowPassword] = useState(false);

  return (
    <FormControl {...props}>
      <InputLabel htmlFor="login-password">Password</InputLabel>
      <OutlinedInput
        value={password}
        onChange={(e) => {
          setPassword(e.target.value);
        }}
        id="form-element-password"
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
  );
};

export default Password;
