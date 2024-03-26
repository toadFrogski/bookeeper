import { FC } from "react";
import { BookCard, SearchBar } from "../../components";
import { Box, Button, Container } from "@mui/material";
import { useTranslation } from "react-i18next";

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

const Home: FC = () => {
  const [t] = useTranslation();

  return (
    <Container>
      <SearchBar sx={{ mt: 3 }} />
      <Box
        component="section"
        sx={{
          display: "grid",
          gap: "20px",
          mt: 3,
          gridTemplateColumns: "repeat(auto-fill, 200px)",
          justifyContent: "space-between",
        }}
      >
        {mock.map((book, idx) => (
          <BookCard
            key={`book-card-${idx}`}
            title={book.title}
            photo={book.photo}
            author={book.author}
            owner={book.owner}
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
