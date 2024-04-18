import { Close } from "@mui/icons-material";
import { Alert, AlertProps, IconButton, Snackbar, SnackbarProps } from "@mui/material";
import { FC, useContext, useEffect, useState } from "react";
import { NotificationContext } from "../../contexts/notification";

export type NotifierMessage = {
  type: "success" | "info" | "warning" | "error";
  message: string;
} | null;

type Props = {
  alertProps?: AlertProps;
  snackbarProps?: SnackbarProps;
};

const Notifier: FC<Props> = ({ alertProps, snackbarProps }) => {
  const [isNotifierOpen, setIsNotifierOpen] = useState(false);
  const { notification } = useContext(NotificationContext);

  const handleClose = () => setIsNotifierOpen(false);
  const handleOpen = () => setIsNotifierOpen(true);

  useEffect(() => {
    if (notification !== null) {
      handleOpen();
    }
  }, [notification]);

  return (
    <Snackbar {...snackbarProps} open={isNotifierOpen} autoHideDuration={1500} onClose={handleClose}>
      <Alert
        {...alertProps}
        severity={notification?.type}
        action={
          <>
            <IconButton aria-label="close" color="inherit" sx={{ p: 0.5 }} onClick={handleClose}>
              <Close />
            </IconButton>
          </>
        }
      >
        {notification?.message}
      </Alert>
    </Snackbar>
  );
};

export default Notifier;
