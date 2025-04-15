import React from "react";
import { TextField, TextFieldProps } from "@mui/material";

const FormInput: React.FC<TextFieldProps> = (props) => {
  return (
    <TextField
      fullWidth
      variant="outlined"
      {...props} // Spread props correctly here
    />
  );
};

export default FormInput;
