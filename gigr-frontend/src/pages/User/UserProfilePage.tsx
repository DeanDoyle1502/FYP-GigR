// src/pages/User/UserProfilePage.tsx
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Layout from "../../components/Layout";
import { Container, Typography, Paper, Box } from "@mui/material";
import axios from "axios";

interface UserProfile {
  id: number;
  name: string;
  email: string;
  instrument: string;
  location: string;
  bio: string;
}

const UserProfilePage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [user, setUser] = useState<UserProfile | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      setError("Unauthorized access.");
      return;
    }

    axios
      .get(`http://localhost:8080/users/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => setUser(res.data))
      .catch((err) => {
        console.error("Error fetching user profile:", err);
        setError("Failed to load user profile.");
      });
  }, [id]);

  return (
    <Layout>
      <Container maxWidth="sm" sx={{ mt: 4 }}>
        {error ? (
          <Typography color="error">{error}</Typography>
        ) : !user ? (
          <Typography>Loading profile...</Typography>
        ) : (
          <Paper elevation={3} sx={{ p: 4, borderRadius: 3 }}>
            <Typography variant="h4" fontWeight="bold" gutterBottom>
              {user.name}
            </Typography>
            <Box mb={2}><strong>Email:</strong> {user.email}</Box>
            <Box mb={2}><strong>Instrument:</strong> {user.instrument}</Box>
            <Box mb={2}><strong>Location:</strong> {user.location}</Box>
            <Box><strong>Bio:</strong> {user.bio}</Box>
          </Paper>
        )}
      </Container>
    </Layout>
  );
};

export default UserProfilePage;
