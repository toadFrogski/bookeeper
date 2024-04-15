import { Box, BoxProps, Typography } from "@mui/material";
import styles from "./styles.module.scss";
import { FC, ReactNode } from "react";

type Props = BoxProps & {
  title: string;
  author: string;
  photo: string;
  owner: string;
  renderActions: ReactNode;
};

const BookCard: FC<Props> = ({ title, author, photo, renderActions, ...props }) => {
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
          <Typography variant="caption" className={styles.text}>
            {author}
          </Typography>
        </Box>
        {renderActions}
      </div>
    </Box>
  );
};

export default BookCard;
