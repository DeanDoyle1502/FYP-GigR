import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import Layout from "../../components/Layout";
import Button from "../../components/Button";
import { Box, Typography, Paper } from "@mui/material";
import { Gig } from "../../types/gig";

const GigDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [gig, setGig] = useState<Gig | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      setError("You must be logged in to view this gig.");
      return;
    }

    axios
      .get(`http://localhost:8080/gigs/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => setGig(res.data))
      .catch((err) => {
        console.error("Error fetching gig:", err);
        setError("Could not load gig.");
      });
  }, [id]);

  const handleDelete = async () => {
    if (!window.confirm("Are you sure you want to delete this gig?")) return;

    const token = localStorage.getItem("token");
    try {
      await axios.delete(`http://localhost:8080/gigs/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      alert("Gig deleted.");
      navigate("/gigs/mine");
    } catch (err: any) {
      console.error("Failed to delete gig:", err);
      alert("Failed to delete gig.");
    }
  };

  return (
    <Layout>
      <Box maxWidth="700px" mx="auto" mt={5}>
        {error && <Typography color="error">{error}</Typography>}
        {!gig ? (
          <Typography>Loading gig...</Typography>
        ) : (
          <Paper elevation={3} sx={{ p: 4, borderRadius: 3 }}>
            <Typography variant="h4" gutterBottom>
              {gig.title}
            </Typography>
            <Typography><strong>Date:</strong> {new Date(gig.date).toLocaleString()}</Typography>
            <Typography><strong>Location:</strong> {gig.location}</Typography>
            <Typography><strong>Instrument:</strong> {gig.instrument}</Typography>
            <Typography><strong>Status:</strong> {gig.status}</Typography>
            <Typography mt={2}>{gig.description}</Typography>

            <Box display="flex" gap={2} mt={4} flexWrap="wrap">
              <Button onClick={() => navigate("/dashboard")}>Dashboard</Button>
              <Button onClick={() => navigate("/gigs/mine")} style={{ backgroundColor: "#007bff" }}>
                My Gigs
              </Button>
              <Button onClick={handleDelete} style={{ backgroundColor: "#dc3545" }}>
                Delete Gig
              </Button>
              <Button onClick={() => navigate(`/gigs/${id}/edit`)} style={{ backgroundColor: "#ffc107", color: "#000" }}>
                Edit Gig
              </Button>
            </Box>
          </Paper>
        )}
      </Box>
    </Layout>
  );
};

export default GigDetailPage;
