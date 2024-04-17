import { FC, useContext, useEffect, useState } from "react";
import Forms from "../../components/Forms";
import { AnyResponse, Book } from "../../../services/api";
import { Button, Container } from "@mui/material";
import { useTranslation } from "react-i18next";
import { ApiContext } from "../../contexts/api";
import { useStateWithError } from "../../../utils/hooks";
import { Notifier, NotifierMessage } from "../../components";
import { isAxiosError } from "axios";

const ProfileAddBook: FC = () => {
  const [t] = useTranslation();
  const { bookApi } = useContext(ApiContext);

  const book = useStateWithError<Book>({});
  const [bookImage, setBookImage] = useState<File>();
  const [notifyMessage, setNotifyMessage] = useState<NotifierMessage>(null);

  useEffect(() => {
    if (book.error !== "") {
      setNotifyMessage({ type: "error", message: t(book.error) });
      book.setError("");
    }
  }, [book, t]);

  const handleSaveBook = () => {
    if (
      bookImage === undefined ||
      book.value.name === undefined ||
      book.value.author === undefined ||
      book.value.description === undefined
    ) {
      return;
    }

    bookApi
      .bookSavePost(bookImage, book.value.name, book.value.author, book.value.description)
      .then(() => {
        setNotifyMessage({ type: "success", message: "profile.bookSaved" });
      })
      .catch((err) => {
        if (isAxiosError(err)) {
          if (err.response && err.response.data) {
            const data = err.response.data as AnyResponse;
            switch (data.response_message) {
              default:
                book.setError("error.unexpectedError");
                break;
            }
          } else book.setError("error.unexpectedError");
        } else book.setError("error.unexpectedError");
      });
  };

  return (
    <>
      <Forms.BookForm
        book={book.value}
        onPhotoChange={(file) => setBookImage(file)}
        onTitleChange={(value) =>
          book.setValue((book) => {
            return { ...book, name: value };
          })
        }
        onAuthorChange={(value) =>
          book.setValue((book) => {
            return { ...book, author: value };
          })
        }
        onDescriptionChange={(value) =>
          book.setValue((book) => {
            return { ...book, description: value };
          })
        }
      />
      <Container maxWidth="md">
        <Button fullWidth variant="contained" onClick={handleSaveBook}>
          {t("common.save")}
        </Button>
      </Container>
      <Notifier message={notifyMessage} alertProps={{ variant: "filled" }} />
    </>
  );
};

export default ProfileAddBook;
