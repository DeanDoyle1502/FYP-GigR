import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import Layout from "../../components/Layout";
import Button from "../../components/Button";
import { Box, Typography, Paper } from "@mui/material";
import { Gig } from "../../types/gig";

interface User {
  id: number;
  name: string;
  email: string;
}

const GigDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [gig, setGig] = useState<Gig | null>(null);
  const [user, setUser] = useState<User | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isOwner, setIsOwner] = useState<boolean>(false);
  const [hasApplied, setHasApplied] = useState<boolean>(false);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setError("You must be logged in to view this gig.");
      return;
    }

    const fetchGig = axios.get(`http://localhost:8080/gigs/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    const fetchUser = axios.get(`http://localhost:8080/auth/me`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    Promise.all([fetchGig, fetchUser])
      .then(([gigRes, userRes]) => {
        const userData = {
          id: userRes.data.id || userRes.data.ID,
          name: userRes.data.name || userRes.data.Name,
          email: userRes.data.email || userRes.data.Email,
        };

        const gigData = gigRes.data;

        setUser(userData);
        setGig(gigData);
        setIsOwner(userData.id === gigData.user?.id);
      })
      .catch((err) => {
        console.error("Error loading gig or user:", err);
        setError("Could not load gig or user.");
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
    } catch (err) {
      console.error("Failed to delete gig:", err);
      alert("Failed to delete gig.");
    }
  };

  const handleApply = async () => {
    const token = localStorage.getItem("token");
    try {
      await axios.post(
        `http://localhost:8080/gigs/${id}/apply`,
        {},
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );
      alert("Application submitted!");
      setHasApplied(true);
    } catch (err: any) {
      console.error("Failed to apply for gig:", err);
      alert(
        err.response?.data?.error || "Failed to apply. Try again or check console."
      );
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
            <Typography><strong>Posted by:</strong> {gig.user?.name || gig.user?.email}</Typography>
            <Typography mt={2}>{gig.description}</Typography>

            <Box display="flex" gap={2} mt={4} flexWrap="wrap">
              <Button onClick={() => navigate("/dashboard")}>Dashboard</Button>
              <Button onClick={() => navigate("/gigs/mine")} style={{ backgroundColor: "#007bff" }}>
                My Gigs
              </Button>

              {isOwner && (
                <>
                  <Button onClick={handleDelete} style={{ backgroundColor: "#dc3545" }}>
                    Delete Gig
                  </Button>
                  <Button onClick={() => navigate(`/gigs/${id}/edit`)} style={{ backgroundColor: "#ffc107", color: "#000" }}>
                    Edit Gig
                  </Button>
                </>
              )}

              {!isOwner && gig.status === "Available" && !hasApplied && (
                <Button
                  onClick={handleApply}
                  style={{ backgroundColor: "#28a745", color: "#fff" }}
                >
                  Apply for Gig
                </Button>
              )}

              {!isOwner && hasApplied && (
                <Typography mt={2} color="primary">
                  You've already applied to this gig.
                </Typography>
              )}
            </Box>
          </Paper>
        )}
      </Box>
    </Layout>
  );
};

export default GigDetailPage;
