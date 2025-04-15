// src/components/Button.tsx
import React from "react";
import { Button as MUIButton, ButtonProps } from "@mui/material";

const Button: React.FC<ButtonProps> = ({ children, ...props }) => {
  return (
    <MUIButton
      variant="contained"
      color="primary"
      fullWidth
      sx={{ mt: 2, textTransform: "none", fontWeight: 500 }}
      {...props}
    >
      {children}
    </MUIButton>
  );
};

export default Button;
