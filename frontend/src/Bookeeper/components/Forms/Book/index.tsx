import { Box, Container, ContainerProps, FormControl, Paper, Skeleton, TextField } from "@mui/material";
import { FC, useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { Book } from "../../../../services/api";
import styles from "./styles.module.scss";

const mockBook: Book = {
  id: 2,
  name: "Little Prince",
  description:
    "The Little Prince (French: Le Petit Prince, pronounced [lə p(ə)ti pʁɛ̃s]) is a novella written and illustrated by French writer, and military pilot, Antoine de Saint-Exupéry. It was first published in English and French in the United States by Reynal & Hitchcock in April 1943 and was published posthumously in France following liberation; Saint-Exupéry's works had been banned by the Vichy Regime. The story follows a young prince who visits various planets, including Earth, and addresses themes of loneliness, friendship, love, and loss. Despite its style as a children's book, The Little Prince makes observations about life, adults, and human nature",
  photo: "https://m.media-amazon.com/images/I/71QKDKxL-jL._AC_UF1000,1000_QL80_.jpg",
  author: "antuan de saint exupery",
  user: {
    email: "cutegr49@gmail.com",
  },
};

type Props = ContainerProps;

const BookForm: FC<Props> = (props) => {
  const [t] = useTranslation();
  const [book, setBook] = useState<Book>({});

  // Visual effects
  const [imageLoaded, setImageLoaded] = useState(false);

  useEffect(() => {
    setTimeout(() => setBook(mockBook), 1000);
  }, []);

  return (
    <Container {...props} sx={{ p: 2, mt: 10 }} maxWidth="md">
      <Paper elevation={3} sx={{p: 1}}>
        <Box component="section" className={styles.formContainer}>
          <div className={styles.imageContainer}>
            <img
              src={book.photo}
              onLoad={() => setImageLoaded(true)}
              style={{ display: imageLoaded ? "block" : "none" }}
            />
            {!imageLoaded && <Skeleton height={600} variant="rounded" />}
          </div>
          <div className={styles.infoContainer}>
            {book.name ? (
              <FormControl variant="standard">
                <TextField
                  variant="standard"
                  placeholder="Book title"
                  value={book.name}
                  onChange={(e) =>
                    setBook((book) => {
                      return { ...book, name: e.target.value };
                    })
                  }
                  InputProps={{
                    style: { fontSize: 24 },
                    disableUnderline: true,
                  }}
                />
              </FormControl>
            ) : (
              <Skeleton variant="rounded" height={40} />
            )}
            {book.description ? (
              <FormControl variant="standard" sx={{ width: "100%" }}>
                <TextField
                  variant="standard"
                  placeholder="Book description"
                  multiline
                  value={book.description}
                  onChange={(e) =>
                    setBook((book) => {
                      return { ...book, description: e.target.value };
                    })
                  }
                  InputProps={{
                    disableUnderline: true,
                  }}
                />
              </FormControl>
            ) : (
              <Skeleton variant="rounded" height={200} sx={{ mt: 1 }} />
            )}
          </div>
        </Box>
      </Paper>
    </Container>
  );
};

export default BookForm;
