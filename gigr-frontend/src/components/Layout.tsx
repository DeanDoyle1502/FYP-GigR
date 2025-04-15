// src/components/Layout.tsx
import React from "react";
import { Container, Box } from "@mui/material";
import NavBar from "./NavBar";

const Layout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <>
      <NavBar />
      <Container maxWidth="md">
        <Box mt={4}>{children}</Box>
      </Container>
    </>
  );
};

export default Layout;
