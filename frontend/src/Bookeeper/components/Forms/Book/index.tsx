import { Box, Container, ContainerProps, FormControl, Paper, Skeleton, TextField } from "@mui/material";
import { FC, useEffect, useState } from "react";
import { Book } from "../../../../services/api";
import styles from "./styles.module.scss";
import { useDropzone } from "react-dropzone";
import { Image } from "@mui/icons-material";

type Props = ContainerProps & {
  book: Book;
  onTitleChange: (value: string) => void;
  onDescriptionChange: (value: string) => void;
  onPhotoChange: (file: File) => void;
};

const BookForm: FC<Props> = ({ book, onTitleChange, onDescriptionChange, onPhotoChange, ...props }) => {
  const { acceptedFiles, getRootProps, getInputProps } = useDropzone({
    accept: {
      "image/png": [".png"],
      "image/jpeg": [".jpeg"],
    },
  });

  // Visual effects
  const [imageLoaded, setImageLoaded] = useState(false);
  const [preview, setPreview] = useState<string>();
  const [isDragOver, setIsDragOver] = useState(false);

  useEffect(() => {
    acceptedFiles.map((file) => {
      setPreview(URL.createObjectURL(file));
      onPhotoChange(file);
    });
  }, [acceptedFiles, onPhotoChange]);

  return (
    <Container {...props} sx={{ p: 2, mt: 10 }} maxWidth="md">
      <Paper elevation={3} sx={{ p: 1 }}>
        <Box component="section" className={styles.formContainer}>
          <div
            {...getRootProps({ className: `${styles.imageContainer} ${isDragOver ? styles.activeDropzone : ""}` })}
            onLoad={() => setImageLoaded(true)}
          >
            {book.photo || preview ? (
              <>
                <img
                  src={preview ? preview : book.photo}
                  style={{ display: imageLoaded ? "block" : "none" }}
                  onDragLeave={() => setIsDragOver(false)}
                  onDrop={() => setIsDragOver(false)}
                  onDragEnter={() => setIsDragOver(true)}
                />
                <input {...getInputProps()} />
                {!imageLoaded && <Skeleton height={600} variant="rounded" />}
              </>
            ) : (
              <Box
                component="div"
                sx={{ display: "flex", height: "100%", alignItems: "center", alignSelf: "center" }}
                onDragLeave={() => setIsDragOver(false)}
                onDrop={() => setIsDragOver(false)}
                onDragEnter={() => setIsDragOver(true)}
              >
                <Image />
                <input {...getInputProps()} />
              </Box>
            )}
          </div>
          <div className={styles.infoContainer}>
            <FormControl variant="standard">
              <TextField
                variant="standard"
                placeholder="Book title"
                value={book.name}
                onChange={(e) => onTitleChange(e.target.value)}
                InputProps={{
                  style: { fontSize: 24 },
                  disableUnderline: true,
                }}
              />
            </FormControl>
            <FormControl variant="standard" sx={{ width: "100%" }}>
              <TextField
                variant="standard"
                placeholder="Book description"
                multiline
                value={book.description}
                onChange={(e) => onDescriptionChange(e.target.value)}
                InputProps={{
                  disableUnderline: true,
                }}
              />
            </FormControl>
          </div>
        </Box>
      </Paper>
    </Container>
  );
};

export default BookForm;
