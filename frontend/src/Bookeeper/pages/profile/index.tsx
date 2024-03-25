import { Box, Container, Fab, Menu, Typography } from "@mui/material";
import { FC, useEffect, useRef, useState } from "react";
import { BookCard } from "../../components";
import { Menu as MenuIcon } from "@mui/icons-material";

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
  const menuRef = useRef();
  const [isMenuOpen, setIsMenuOpen] = useState(false);

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
          <Box component="div">
            <BookCard
              key={`book-card-${idx}`}
              title={book.title}
              photo={book.photo}
              author={book.author}
              owner={book.owner}
            />
          </Box>
        ))}
      </Box>
      <Fab
        ref={menuRef}
        color="primary"
        sx={{ position: "absolute", bottom: "24px", right: "24px" }}
        onClick={() => setIsMenuOpen(true)}
      >
        <MenuIcon />
      </Fab>
      <Menu anchorEl={menuRef.current} open={isMenuOpen} onClose={() => setIsMenuOpen(false)}>
        <p>test</p>
        <p>test</p>
        <p>test</p>
      </Menu>
    </Container>
  );
};

export default Profile;
