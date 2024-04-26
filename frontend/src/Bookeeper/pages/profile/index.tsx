import { Box, Button, Container, Modal, Paper, Typography } from "@mui/material";
import { FC, useContext, useEffect, useState } from "react";
import { BookActionsMenu, BookCard, PlusFAB } from "../../components";
import styles from "./styles.module.scss";
import { useNavigate } from "react-router-dom";
import urls, { parseURL } from "../../utils/urls";
import { ApiContext } from "../../contexts/api";
import { Book } from "../../../services/api";
import { t } from "i18next";
import { NotificationContext } from "../../contexts/notification";
import { isAxiosError } from "axios";
import Boo from "../../assets/svg/boo.svg?react";

const Profile: FC = () => {
  const navigate = useNavigate();
  const { bookApi } = useContext(ApiContext);
  const { setNotification } = useContext(NotificationContext);

  const [books, setBooks] = useState<Book[]>([]);
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState(false);
  const [selectedBook, setSelectedBook] = useState(0);

  const deleteBook = (bookID: number) => {
    bookApi
      .bookBookIdDelete(bookID)
      .then(() => {
        setBooks(books.filter((book) => book.ID != bookID));
        setNotification({ type: "info", message: "Book was deleted :(" });
      })
      .catch(() => {
        setNotification({ type: "error", message: t("error.unexpectedError") });
      });
  };

  useEffect(() => {
    bookApi
      .bookUserMeGet()
      .then((res) => {
        if (res.data.data) {
          setBooks(res.data.data);
        }
      })
      .catch((err) => {
        if (isAxiosError(err)) {
          /** @TODO Handle errors  */
        } else setNotification({ type: "error", message: t("error.unexpectedError") });
      });
  }, []);

  return (
    <Container sx={{ mt: 5, pb: 12 }}>
      <Typography variant="h5">My books</Typography>
      {books.length !== 0 ? (
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
                  onDelete={() => {
                    setIsDeleteModalOpen(true);
                    setSelectedBook(book.ID ?? 0);
                  }}
                  onEdit={() => {
                    const path = parseURL(urls.profileEditBook);
                    navigate(`${path[0]}${path[1]}/${book.ID}`);
                  }}
                />
              }
            />
          ))}
        </Box>
      ) : (
        <Box>
          <Boo
            style={{ maxWidth: "200px", maxHeight: "200px", filter: "invert(1)", margin: "0 auto", display: "block" }}
          />
          <p style={{ textAlign: "center" }}>Oh... I guess you don't have any books yet</p>
        </Box>
      )}
      <PlusFAB onClick={() => navigate(urls.profileAddBook)} />
      <Modal
        open={isDeleteModalOpen}
        onClose={() => setIsDeleteModalOpen(false)}
        sx={{ display: "flex", alignItems: "center" }}
      >
        <Container
          maxWidth="sm"
          component={Paper}
          elevation={3}
          sx={{ height: "150px", display: "flex", flexDirection: "column" }}
        >
          <Typography variant="h5" sx={{ mt: 3, flex: 1 }}>
            Are you sure?
          </Typography>
          <Box component="div" sx={{ direction: "rtl" }}>
            <Button
              sx={{ mb: 3 }}
              variant="outlined"
              color="error"
              onClick={() => {
                setIsDeleteModalOpen(false);
                deleteBook(selectedBook);
                setSelectedBook(0);
              }}
            >
              {t("common.delete")}
            </Button>
          </Box>
        </Container>
      </Modal>
    </Container>
  );
};

export default Profile;
