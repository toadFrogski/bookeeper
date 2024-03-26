import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import { FormControl, FormControlProps, IconButton, InputAdornment, TextField } from "@mui/material";
import { FC, useState } from "react";
import { useTranslation } from "react-i18next";

type Props = Omit<FormControlProps, "error"> & {
  password: string;
  error: string;
  setPassword: (password: string) => void;
};

const Password: FC<Props> = ({ password, setPassword, error, ...props }: Props) => {
  const [showPassword, setShowPassword] = useState(false);
  const [t] = useTranslation();

  return (
    <FormControl {...props}>
      <TextField
        error={error !== ""}
        variant="outlined"
        value={password}
        type={showPassword ? "text" : "password"}
        helperText={error}
        onChange={(e) => setPassword(e.target.value)}
        label={t("common.password")}
        InputProps={{
          autoComplete: "new-password",
          endAdornment: (
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
          ),
        }}
      />
    </FormControl>
  );
};

export default Password;
