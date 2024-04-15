import { Box, Container, Typography } from "@mui/material";
import { FC } from "react";
import { BookActionsMenu, BookCard, PlusFAB } from "../../components";
import styles from "./styles.module.scss";
import { useNavigate } from "react-router-dom";
import urls, { parseURL } from "../../utils/urls";

type Book = {
  id: number;
  title: string;
  photo: string;
  author: string;
  owner: string;
};

const mock: Book[] = [
  {
    id: 1,
    title: "Little Prince",
    photo: "https://m.media-amazon.com/images/I/71QKDKxL-jL._AC_UF1000,1000_QL80_.jpg",
    author: "antuan de saint exupery",
    owner: "cutegr49@gmail.com",
  },
  {
    id: 2,
    title: "Little Prince",
    photo: "https://m.media-amazon.com/images/I/71QKDKxL-jL._AC_UF1000,1000_QL80_.jpg",
    author: "antuan de saint exupery",
    owner: "cutegr49@gmail.com",
  },
  {
    id: 3,
    title: "Little Prince",
    photo: "https://m.media-amazon.com/images/I/71QKDKxL-jL._AC_UF1000,1000_QL80_.jpg",
    author: "antuan de saint exupery",
    owner: "cutegr49@gmail.com",
  },
];

const Profile: FC = () => {
  const navigate = useNavigate();

  return (
    <Container sx={{ mt: 5, pb: 12 }}>
      <Typography variant="h5">My books</Typography>
      <Box component="section" className={styles.bookContainer} sx={{ mt: 3 }}>
        {mock.map((book, idx) => (
          <BookCard
            sx={{ mt: 2 }}
            key={`book-card-${idx}`}
            title={book.title}
            photo={book.photo}
            author={book.author}
            owner={book.owner}
            renderActions={
              <BookActionsMenu
                sx={{ mt: 1 }}
                onDelete={() => console.log(123)}
                onEdit={() => {
                  const path = parseURL(urls.profileEditBook);
                  navigate(`${path[0]}${path[1]}/${book.id}`);
                }}
              />
            }
          />
        ))}
      </Box>
      <PlusFAB onClick={() => navigate(urls.profileAddBook)} />
    </Container>
  );
};

export default Profile;
