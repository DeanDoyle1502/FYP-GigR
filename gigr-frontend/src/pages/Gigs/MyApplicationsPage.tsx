import React, { useEffect, useState } from "react";
import axios from "axios";
import Layout from "../../components/Layout";
import { Box, Typography, Paper, Container, Button } from "@mui/material";
import { useNavigate } from "react-router-dom";

interface Gig {
  id: number;
  title: string;
  location: string;
  date: string;
  status: string;
}

interface Application {
  id: number;
  status: string;
  gig: Gig;
}

const MyApplicationsPage: React.FC = () => {
  const [applications, setApplications] = useState<Application[]>([]);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) return;

    axios
      .get("http://localhost:8080/gigs/applications/mine", {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => setApplications(res.data))
      .catch((err) => {
        console.error("Failed to load applications:", err);
        setError("Could not load your applications.");
      });
  }, []);

  return (
    <Layout>
      <Container maxWidth="md" sx={{ mt: 4 }}>
        <Typography variant="h4" gutterBottom>
          My Applications
        </Typography>
        {error && <Typography color="error">{error}</Typography>}
        {applications.map((app) => (
  <Paper
    key={app.id}
    sx={{ p: 2, mb: 2, cursor: 'pointer', '&:hover': { backgroundColor: '#f5f5f5' } }}
    onClick={() => navigate(`/gigs/${app.gig.id}`)}
  >
    <Typography><strong>Gig:</strong> {app.gig.title}</Typography>
    <Typography><strong>Location:</strong> {app.gig.location}</Typography>
    <Typography><strong>Date:</strong> {new Date(app.gig.date).toLocaleString()}</Typography>
    <Typography><strong>Application Status:</strong> {app.status}</Typography>
    <Typography><strong>Gig Status:</strong> {app.gig.status}</Typography>
    <Button
              onClick={() => navigate(`/gigs/${app.gig.id}`)}
              sx={{ mt: 1 }}
              variant="outlined"
            >
              View Gig Details
            </Button>
          </Paper>
        ))}
      </Container>
    </Layout>
  );
};

export default MyApplicationsPage;
