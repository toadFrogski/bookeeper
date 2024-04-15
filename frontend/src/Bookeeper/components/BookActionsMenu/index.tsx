import { Brush, Delete } from "@mui/icons-material";
import { Box, Fab, SxProps } from "@mui/material";
import { FC } from "react";

type Props = {
  onDelete: () => void;
  onEdit: () => void;
  sx?: SxProps;
};

const BookActionsMenu: FC<Props> = ({ onDelete, onEdit, sx }) => {
  return (
    <Box component="div" sx={{ ...sx, justifyContent: "space-around", display: "flex", userSelect: "none" }}>
      <Fab size="small" onClick={() => onEdit()}>
        <Brush />
      </Fab>
      <Fab size="small" sx={{ display: "inline-flex" }} onClick={() => onDelete()}>
        <Delete />
      </Fab>
    </Box>
  );
};

export default BookActionsMenu;
