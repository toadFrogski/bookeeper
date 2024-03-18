import { Box, BoxProps, Button, Typography } from "@mui/material";
import styles from "./styles.module.scss";
import { FC } from "react";

type Props = BoxProps & {
  title: string;
  author: string;
  photo: string;
  owner: string;
};

const BookCard: FC<Props> = ({ title, author, photo, owner, ...props }) => {
  return (
    <Box {...props} className={styles.bookCard} component="div">
      <div className={styles.photoWrap}>
        <img className={styles.photo} src={photo} alt="book-image" />
      </div>
      <div className={styles.infoWrap}>
        <Box component="div" sx={{ flex: 1 }}>
          <Typography className={styles.text} variant="h5">
            {title}
          </Typography>
          <Typography className={styles.text}>{author}</Typography>
          <Typography className={styles.text}>{owner}</Typography>
        </Box>
        <Button variant="outlined" color="inherit" sx={{ border: "2px solid"}}>
          Inquire
        </Button>
      </div>
    </Box>
  );
};

export default BookCard;
