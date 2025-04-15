// src/theme.ts
import { createTheme } from "@mui/material/styles";

// Customize your design here
const theme = createTheme({
  palette: {
    primary: {
      main: "#007bff", // Blue
    },
    secondary: {
      main: "#6f42c1", // Purple
    },
    success: {
      main: "#28a745",
    },
    error: {
      main: "#dc3545",
    },
    background: {
      default: "#f5f5f5",
    },
  },
  typography: {
    fontFamily: `"Roboto", "Helvetica", "Arial", sans-serif`,
    h1: {
      fontSize: "2rem",
      fontWeight: 700,
    },
    button: {
      textTransform: "none", 
      fontWeight: 600,
    },
  },
  shape: {
    borderRadius: 8,
  },
});

export default theme;
