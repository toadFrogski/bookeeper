import { FC, useContext, useEffect, useState } from "react";
import Forms from "../../components/Forms";
import { AnyResponse, Book } from "../../../services/api";
import { Button, Container } from "@mui/material";
import { useTranslation } from "react-i18next";
import { useParams } from "react-router-dom";
import { ApiContext } from "../../contexts/api";
import { isAxiosError } from "axios";
import { useStateWithError } from "../../../utils/hooks";
import { NotificationContext } from "../../contexts/notification";

const ProfileEditBook: FC = () => {
  const [t] = useTranslation();
  const params = useParams();
  const { bookApi } = useContext(ApiContext);
  const { setNotification } = useContext(NotificationContext);

  const book = useStateWithError<Book>({});
  const [bookImage, setBookImage] = useState<File>();

  useEffect(() => {
    if (book.error !== "") {
      setNotification({ type: "error", message: t(book.error) });
    }
  }, [book.error, setNotification, t]);

  useEffect(() => {
    bookApi
      .bookBookIdGet(parseInt(params.bookID as string))
      .then((res) => {
        if (res.data.data) {
          book.setValue(res.data.data);
        }
      })
      .catch((err) => {
        if (isAxiosError(err)) {
          if (err.response && err.response.data) {
            const data = err.response.data as AnyResponse;
            switch (data.response_message) {
              case "DATA_NOT_FOUND":
                book.setError("error.bookNotFound");
                break;
              default:
                book.setError("error.unexpectedError");
                break;
            }
          } else book.setError("error.unexpectedError");
        } else book.setError("error.unexpectedError");
      });
  }, []);

  const handleSaveBook = () => {
    if (book.value.ID) {
      bookApi
        .bookBookIdPost(book.value.ID, bookImage, book.value.name, book.value.author, book.value.description)
        .then(() => setNotification({ type: "success", message: t("success") }))
        .catch(() => setNotification({ type: "error", message: t("error.unexpectedError") }));
    }
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
    </>
  );
};

export default ProfileEditBook;
