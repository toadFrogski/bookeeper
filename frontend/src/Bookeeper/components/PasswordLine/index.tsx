import { Box, BoxProps, useTheme } from "@mui/material";
import { Palette } from "@mui/material/styles";
import { FC } from "react";

type Props = Omit<BoxProps, "display" | "alignItems" | "gap" | "margin"> & {
  strength: number;
};

const generateColors = (strength: number, palette: Palette) => {
  if (strength >= 80) {
    return [palette.primary.main, palette.primary.main, palette.primary.main];
  } else if (strength >= 50) {
    return [palette.primary.main, palette.primary.main, palette.background.default];
  } else if (strength >= 25) {
    return [palette.warning.main, palette.background.default, palette.background.default];
  } else {
    return [palette.background.default, palette.background.default, palette.background.default];
  }
};

const PasswordLine: FC<Props> = ({ strength, ...props }) => {
  const theme = useTheme();
  const colors = generateColors(strength, theme.palette);
  return (
    <Box {...props} display="flex" alignItems="center" justifyContent="center" gap="5px">
      {colors.map((color, index) => (
        <Box key={`password-color-${index}`} flex={1} height="5px" bgcolor={color} borderRadius="4px"></Box>
      ))}
    </Box>
  );
};

export default PasswordLine;
