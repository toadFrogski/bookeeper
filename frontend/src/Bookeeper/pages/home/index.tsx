import { FC, useContext, useEffect, useState } from "react";
import { BookCard, SearchBar } from "../../components";
import { Box, Button, Container } from "@mui/material";
import { useTranslation } from "react-i18next";
import styles from "./styles.module.scss";
import { ApiContext } from "../../contexts/api";
import { BookPaginator } from "../../../services/api";

const Home: FC = () => {
  const { bookApi } = useContext(ApiContext);
  const [t] = useTranslation();

  const [books, setBooks] = useState<BookPaginator>();

  useEffect(() => {
    bookApi.bookGet().then((res) => {
      if (res.data.data) {
        setBooks(res.data.data);
      }
    });
  }, []);

  return (
    <Container>
      <SearchBar sx={{ mt: 3 }} />
      <Box
        component="section"
        className={styles.bookContainer}
        sx={{
          mt: 3,
        }}
      >
        {books?.content?.map((book, idx) => (
          <BookCard
            sx={{
              mb: 2,
            }}
            key={`book-card-${idx}`}
            title={book.name ?? ""}
            photo={book.photo ?? ""}
            author={book.author ?? ""}
            owner={book.user?.email ?? ""}
            renderActions={
              <Button variant="outlined" color="inherit" sx={{ mt: 1, border: "2px solid" }}>
                {t("home.inquire")}
              </Button>
            }
          />
        ))}
      </Box>
    </Container>
  );
};

export default Home;
