import { ErrorOutline } from "@mui/icons-material";
import { Box, Typography, useTheme } from "@mui/material";
import { FC } from "react";

type Props = {
  message: string;
};

const CommonErrorBox: FC<Props> = ({ message }) => {
  const { palette } = useTheme();

  return (
    <Box sx={{ mb: 5, display: "flex", color: palette.error.main }}>
      <ErrorOutline />
      <Typography sx={{ ml: 1 }} variant="inherit">
        {message}
      </Typography>
    </Box>
  );
};

export default CommonErrorBox;
