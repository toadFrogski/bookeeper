import { Box, Chip, Container, Stack, Typography } from "@mui/material";
import { FC } from "react";
import { ActionMenu, BookCard } from "../../components";
import { Edit, Delete } from "@mui/icons-material";
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
];

const Profile: FC = () => {
  const [t] = useTranslation();

  return (
    <Container sx={{ mt: 5 }}>
      <Typography variant="h5">My books</Typography>
      <Box
        component="section"
        sx={{
          display: "grid",
          gap: "20px",
          mt: 3,
          gridTemplateColumns: "repeat(auto-fit, 200px)",
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
              <Stack direction="row" spacing={2} sx={{ mt: 1 }}>
                <Chip
                  label={t("common.Edit")}
                  icon={<Edit />}
                  onClick={() => alert(13)}
                  color="primary"
                  variant="outlined"
                />
                <Chip
                  label={t("common.Delete")}
                  icon={<Delete />}
                  onClick={() => alert(13)}
                  color="error"
                  variant="outlined"
                />
              </Stack>
            }
          />
        ))}
      </Box>
      <ActionMenu />
    </Container>
  );
};

export default Profile;
