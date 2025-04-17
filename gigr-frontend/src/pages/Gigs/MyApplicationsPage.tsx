import React, { useEffect, useState } from "react";
import axios from "axios";
import Layout from "../../components/Layout";
import { Box, Typography, Paper, Container } from "@mui/material";

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
        <Typography variant="h4" gutterBottom>My Applications</Typography>
        {error && <Typography color="error">{error}</Typography>}
        {applications.map((app) => (
          <Paper key={app.id} sx={{ p: 2, mb: 2 }}>
            <Typography><strong>Gig:</strong> {app.gig.title}</Typography>
            <Typography><strong>Location:</strong> {app.gig.location}</Typography>
            <Typography><strong>Date:</strong> {new Date(app.gig.date).toLocaleString()}</Typography>
            <Typography><strong>Status:</strong> {app.status}</Typography>
          </Paper>
        ))}
      </Container>
    </Layout>
  );
};

export default MyApplicationsPage;
