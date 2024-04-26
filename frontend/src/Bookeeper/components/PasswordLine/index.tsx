import { CheckCircleOutline, ErrorOutline } from "@mui/icons-material";
import { Box, BoxProps, Typography, useTheme } from "@mui/material";
import { Palette } from "@mui/material/styles";
import { FC } from "react";
import { useTranslation } from "react-i18next";

type Props = Omit<BoxProps, "display" | "alignItems" | "gap" | "margin"> & {
  strength: number;
  limits: number[];
  showStatus?: boolean;
};

const generateColors = (strength: number, limits: number[], palette: Palette) => {
  if (strength >= limits[2]) {
    return [palette.primary.main, palette.primary.main, palette.primary.main];
  } else if (strength >= limits[1]) {
    return [palette.primary.main, palette.primary.main, palette.background.default];
  } else if (strength >= limits[0]) {
    return [palette.warning.main, palette.background.default, palette.background.default];
  } else {
    return [palette.background.default, palette.background.default, palette.background.default];
  }
};

const PasswordLine: FC<Props> = ({ strength, limits, showStatus, ...props }) => {
  const { palette } = useTheme();
  const colors = generateColors(strength, limits, palette);
  const [t] = useTranslation();

  return (
    <Box>
      <Box {...props} display="flex" alignItems="center" justifyContent="center" gap="5px">
        {colors.map((color, index) => (
          <Box key={`password-color-${index}`} flex={1} height="5px" bgcolor={color} borderRadius="4px"></Box>
        ))}
      </Box>
      {showStatus && (
        <Box
          sx={{
            display: "flex",
            mt: 1,
            color: strength < limits[1] ? palette.warning.main : palette.primary.main,
          }}
        >
          {strength >= limits[1] ? <CheckCircleOutline /> : <ErrorOutline />}
          <Typography variant="inherit" sx={{ ml: 1 }}>
            {strength >= limits[2]
              ? t("common.strongPassword")
              : strength >= limits[1]
              ? t("common.mediumPassword")
              : t("common.weakPassword")}
          </Typography>
        </Box>
      )}
    </Box>
  );
};

export default PasswordLine;
