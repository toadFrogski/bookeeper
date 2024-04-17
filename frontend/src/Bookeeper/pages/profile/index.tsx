import { Box, Container, Typography } from "@mui/material";
import { FC, useContext, useEffect, useState } from "react";
import { BookActionsMenu, BookCard, PlusFAB } from "../../components";
import styles from "./styles.module.scss";
import { useNavigate } from "react-router-dom";
import urls, { parseURL } from "../../utils/urls";
import { ApiContext } from "../../contexts/api";
import { Book } from "../../../services/api";

const Profile: FC = () => {
  const navigate = useNavigate();
  const { bookApi } = useContext(ApiContext);

  const [books, setBooks] = useState<Book[]>([]);

  useEffect(() => {
    bookApi
      .bookUserMeGet()
      .then((res) => {
        if (res.data.data) {
          setBooks(res.data.data);
        }
      })
      .catch((err) => {});
  }, [bookApi]);

  return (
    <Container sx={{ mt: 5, pb: 12 }}>
      <Typography variant="h5">My books</Typography>
      <Box component="section" className={styles.bookContainer} sx={{ mt: 3 }}>
        {books.map((book, idx) => (
          <BookCard
            sx={{ mt: 2 }}
            key={`book-card-${idx}`}
            title={book.name ?? ""}
            photo={book.photo ?? ""}
            author={book.author ?? ""}
            owner={book.user?.email ?? ""}
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
