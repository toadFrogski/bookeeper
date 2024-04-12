import { Box, Container, Popover, Typography } from "@mui/material";
import { FC } from "react";
import { ActionMenu, BookActionsMenu, BookCard } from "../../components";
import { Edit, Delete } from "@mui/icons-material";
import { useTranslation } from "react-i18next";
import styles from "./styles.module.scss";

type Book = {
  title: string;
  photo: string;
  author: string;
  owner: string;
};

const mock: Book[] = [
  {
    title: "Little Prince",
    photo: "https://m.media-amazon.com/images/I/71QKDKxL-jL._AC_UF1000,1000_QL80_.jpg",
    author: "antuan de saint exupery",
    owner: "cutegr49@gmail.com",
  },
  {
    title: "Little Prince",
    photo: "https://m.media-amazon.com/images/I/71QKDKxL-jL._AC_UF1000,1000_QL80_.jpg",
    author: "antuan de saint exupery",
    owner: "cutegr49@gmail.com",
  },
  {
    title: "Little Prince",
    photo: "https://m.media-amazon.com/images/I/71QKDKxL-jL._AC_UF1000,1000_QL80_.jpg",
    author: "antuan de saint exupery",
    owner: "cutegr49@gmail.com",
  },
];

const Profile: FC = () => {
  const [t] = useTranslation();


  return (
    <Container sx={{ mt: 5 }}>
      <Typography variant="h5">My books</Typography>
      <Box component="section" className={styles.bookContainer} sx={{ mt: 3 }}>
        {mock.map((book, idx) => (
          <BookCard
            key={`book-card-${idx}`}
            title={book.title}
            photo={book.photo}
            author={book.author}
            owner={book.owner}
            renderActions={
              <BookActionsMenu bookID={""} />
            }
          />
        ))}
      </Box>
      <ActionMenu />
    </Container>
  );
};

export default Profile;
