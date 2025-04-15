import React, { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import Layout from "../components/Layout";
import Button from "../components/Button";
import { Container, Typography, Stack } from "@mui/material";

interface User {
  id: number;
  name: string;
  email: string;
  instrument: string;
  location: string;
  bio: string;
}

const Dashboard: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      setError("You must be logged in.");
      return;
    }

    axios
      .get("http://localhost:8080/auth/me", {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => {
        const raw = res.data;
        const normalizedUser = {
          id: raw.ID,
          name: raw.Name,
          email: raw.Email,
          instrument: raw.Instrument,
          location: raw.Location,
          bio: raw.Bio,
        };
        setUser(normalizedUser);
      })
      .catch((err) => {
        console.error("Failed to fetch user info:", err);
        setError("Failed to load user profile.");
      });
  }, []);

  if (error) {
    return (
      <Layout>
        <Container>
          <Typography color="error">{error}</Typography>
        </Container>
      </Layout>
    );
  }

  if (!user) {
    return (
      <Layout>
        <Container>
          <Typography>Loading...</Typography>
        </Container>
      </Layout>
    );
  }

  return (
    <Layout>
      <Container maxWidth="sm" sx={{ textAlign: "center", mt: 4 }}>
        <Typography variant="h4" fontWeight="bold" gutterBottom>
          Welcome back, {user.name || user.email}!
        </Typography>
        <Typography>Email: {user.email}</Typography>
        <Typography>Instrument: {user.instrument}</Typography>
        <Typography>Location: {user.location}</Typography>
        <Typography>Bio: {user.bio}</Typography>

        <Stack spacing={2} direction="row" justifyContent="center" mt={4} flexWrap="wrap">
          <Button onClick={() => navigate("/gigs/create")}>
            Create a New Gig
          </Button>
          <Button onClick={() => navigate("/gigs/mine")} color="primary">
            View My Gigs
          </Button>
          <Button onClick={() => navigate("/gigs/public")} color="secondary">
            Browse Available Gigs
          </Button>
        </Stack>
      </Container>
    </Layout>
  );
};

export default Dashboard;
