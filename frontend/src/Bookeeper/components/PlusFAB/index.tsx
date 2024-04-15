import { Add } from "@mui/icons-material";
import { Fab, SxProps } from "@mui/material";
import { FC } from "react";

type Props = {
  sx?: SxProps;
  onClick: () => void;
};

const PlusFAB: FC<Props> = ({ sx, onClick }) => {
  return (
    <Fab color="primary" sx={{ ...sx, position: "fixed", bottom: "24px", right: "24px" }} onClick={() => onClick()}>
      <Add />
    </Fab>
  );
};

export default PlusFAB;
