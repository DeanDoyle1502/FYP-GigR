import React, { useEffect, useState } from "react";
import api from "../api/axios";
import { useNavigate } from "react-router-dom";
import Layout from "../components/Layout";
import Button from "../components/Button";
import {
  Container,
  Typography,
  Card,
  CardContent,
  Divider,
  Stack,
  Box,
} from "@mui/material";

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

    api
      .get("/auth/me", {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => {
        setUser({
          id: res.data.id,
          name: res.data.name,
          email: res.data.email,
          instrument: res.data.instrument,
          location: res.data.location,
          bio: res.data.bio,
        });
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
      <Container maxWidth="md" sx={{ mt: 4 }}>
        <Typography variant="h4" fontWeight="bold" gutterBottom>
          Welcome back, {user.name || user.email}!
        </Typography>

        <Card elevation={3} sx={{ p: 2, borderRadius: 3, mt: 2 }}>
          <CardContent>
            <Typography variant="h6" gutterBottom>
              Your Profile
            </Typography>
            <Divider sx={{ mb: 2 }} />
            <Box mb={1}><strong>Email:</strong> {user.email}</Box>
            <Box mb={1}><strong>Instrument:</strong> {user.instrument}</Box>
            <Box mb={1}><strong>Location:</strong> {user.location}</Box>
            <Box><strong>Bio:</strong> {user.bio}</Box>
          </CardContent>
        </Card>

        <Stack spacing={2} direction="column" justifyContent="center" mt={4} flexWrap="wrap">
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
