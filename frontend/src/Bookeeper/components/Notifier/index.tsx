import { Close } from "@mui/icons-material";
import { Alert, AlertProps, IconButton, Snackbar, SnackbarProps } from "@mui/material";
import { FC, useEffect, useState } from "react";

export type NotifierMessage = {
  type: "success" | "info" | "warning" | "error";
  message: string;
} | null;

type Props = {
  message: NotifierMessage;
  alertProps?: AlertProps;
  snackbarProps?: SnackbarProps;
};

const Notifier: FC<Props> = ({ message, alertProps, snackbarProps }) => {
  const [isNotifierOpen, setIsNotifierOpen] = useState(false);

  const handleClose = () => setIsNotifierOpen(false);
  const handleOpen = () => setIsNotifierOpen(true);

  useEffect(() => {
    if (message !== null) {
      handleOpen();
    }
  }, [message]);

  return (
    <Snackbar {...snackbarProps} open={isNotifierOpen} autoHideDuration={1500} onClose={handleClose}>
      <Alert
        {...alertProps}
        severity={message?.type}
        action={
          <>
            <IconButton aria-label="close" color="inherit" sx={{ p: 0.5 }} onClick={handleClose}>
              <Close />
            </IconButton>
          </>
        }
      >
        {message?.message}
      </Alert>
    </Snackbar>
  );
};

export default Notifier;
