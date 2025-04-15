import React from "react";
import { useNavigate } from "react-router-dom";
import { Box, Container, Typography, Stack } from "@mui/material";
import Layout from "../components/Layout";
import Button from "../components/Button";

const HomePage: React.FC = () => {
  const navigate = useNavigate();

  return (
    <Layout>
      <Container maxWidth="sm">
        <Box textAlign="center" mt={10}>
          <Typography variant="h3" fontWeight="bold" gutterBottom>
            Welcome to GigR 🎸
          </Typography>
          <Typography variant="subtitle1" color="textSecondary" gutterBottom>
            Connect musicians with acts & bands near you.
          </Typography>

          <Stack direction="row" spacing={2} justifyContent="center" mt={4}>
            <Button onClick={() => navigate("/auth/login")}>Log In</Button>
            <Button
              onClick={() => navigate("/auth/register")}
              style={{ backgroundColor: "#28a745" }}
            >
              Create Profile
            </Button>
            <Button
              onClick={() => navigate("/gigs/public")}
              style={{ backgroundColor: "#6f42c1" }}
            >
              Browse Gigs
            </Button>
          </Stack>
        </Box>
      </Container>
    </Layout>
  );
};

export default HomePage;
