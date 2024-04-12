import { Button, List, ListItem, Popover, SxProps } from "@mui/material";
import { FC, useRef, useState } from "react";
import { useTranslation } from "react-i18next";

type Props = {
  bookID: number | string;
  sx?: SxProps;
};

const BookActionsMenu: FC<Props> = ({ bookID, sx }) => {
  const [t] = useTranslation();
  const [open, setOpen] = useState(false);
  const buttonRef = useRef();

  return (
    <>
      <Button ref={buttonRef} sx={sx} color="inherit" variant="outlined" onClick={() => setOpen(true)}>
        test
      </Button>
      <Popover
        anchorEl={buttonRef.current}
        open={open}
        onClose={() => setOpen(false)}
        anchorOrigin={{
          vertical: "top",
          horizontal: "left",
        }}
      >
        <List>
          <ListItem>{t("common.edit")}</ListItem>
        </List>
      </Popover>
    </>
  );
};

export default BookActionsMenu;
