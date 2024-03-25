import { Autocomplete, InputAdornment, TextField, TextFieldProps } from "@mui/material";
import { Search } from '@mui/icons-material';
import { FC } from "react";

type Props = TextFieldProps;

const SearchBar: FC<Props> = ({ ...props }) => {
  return (
    <Autocomplete
      freeSolo
      disableClearable
      renderInput={(params) => (
        <TextField
          {...params}
          {...props}
          variant="outlined"
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <Search />
              </InputAdornment>
            ),
          }}
        />
      )}
      options={[]}
    />
  );
};

export default SearchBar;
