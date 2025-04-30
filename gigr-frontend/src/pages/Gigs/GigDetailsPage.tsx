import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import api from "../../api/axios";
import Layout from "../../components/Layout";
import Button from "../../components/Button";
import { Box, Typography, Paper, Stack } from "@mui/material";
import { Gig } from "../../types/gig";
import ChatView from "../../components/ChatView";

interface User {
  id: number;
  name: string;
  email: string;
}

interface Application {
  id: number;
  musician_id: number;
  status: string;
  musician?: User;
}

const GigDetailsPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [gig, setGig] = useState<Gig | null>(null);
  const [user, setUser] = useState<User | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isOwner, setIsOwner] = useState<boolean>(false);
  const [hasApplied, setHasApplied] = useState<boolean>(false);
  const [applications, setApplications] = useState<Application[]>([]);
  const [selectedChatUserId, setSelectedChatUserId] = useState<number | null>(null);
  const [chatExists, setChatExists] = useState<boolean>(false);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setError("You must be logged in to view this gig.");
      return;
    }

    const fetchGig = api.get(`http://localhost:8080/gigs/details/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    const fetchUser = api.get(`http://localhost:8080/auth/me`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    Promise.all([fetchGig, fetchUser])
      .then(async ([gigRes, userRes]) => {
        const userData = {
          id: userRes.data.id || userRes.data.ID,
          name: userRes.data.name || userRes.data.Name,
          email: userRes.data.email || userRes.data.Email,
        };

        const gigData = gigRes.data;
        setUser(userData);
        setGig(gigData);
        const owner = userData.id === gigData.user?.id;
        setIsOwner(owner);

        if (owner) {
          const appsRes = await api.get(`http://localhost:8080/gigs/details/${gigData.id}/applications`, {
            headers: { Authorization: `Bearer ${token}` },
          });

          const appsWithProfiles = await Promise.all(
            appsRes.data.map(async (app: Application) => {
              const userRes = await api.get(`http://localhost:8080/users/${app.musician_id}`, {
                headers: { Authorization: `Bearer ${token}` },
              });
              return {
                ...app,
                musician: {
                  id: userRes.data.id,
                  name: userRes.data.name,
                  email: userRes.data.email,
                },
              };
            })
          );

          setApplications(appsWithProfiles);
        } else {
          const applied = await api.get(`http://localhost:8080/gigs/details/${gigData.id}/applications`, {
            headers: { Authorization: `Bearer ${token}` },
          });

          const has = applied.data.some((app: Application) => app.musician_id === userData.id);
          setHasApplied(has);

          // ✅ Check if poster has sent any messages
          if (has && gigData.user?.id) {
            try {
              const threadRes = await api.get(`http://localhost:8080/gigs/${gigData.id}/thread/${gigData.user.id}`, {
                headers: { Authorization: `Bearer ${token}` },
              });

              if (threadRes.data && threadRes.data.length > 0) {
                setChatExists(true);
              }
            } catch (err) {
              console.log("No messages from poster yet.");
            }
          }
        }
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
      await api.delete(`http://localhost:8080/gigs/details/${id}`, {
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
      await api.post(
        `http://localhost:8080/gigs/${id}/apply`,
        {},
        { headers: { Authorization: `Bearer ${token}` } }
      );
      alert("Application submitted!");
      setHasApplied(true);
    } catch (err: any) {
      console.error("Failed to apply for gig:", err);
      alert(err.response?.data?.error || "Failed to apply. Try again later.");
    }
  };

  const handleAccept = async (musicianId: number) => {
    const token = localStorage.getItem("token");
    try {
      await api.post(
        `http://localhost:8080/gigs/${id}/accept/${musicianId}`,
        {},
        { headers: { Authorization: `Bearer ${token}` } }
      );
      alert("Musician accepted!");
      setGig((prev) => prev ? { ...prev, status: "Covered" } : prev);
      setApplications((prev) =>
        prev.map((app) =>
          app.musician_id === musicianId
            ? { ...app, status: "accepted" }
            : { ...app, status: "rejected" }
        )
      );
    } catch (err) {
      console.error("Failed to accept musician:", err);
      alert("Could not accept musician.");
    }
  };

  const handleOpenChat = (musicianId: number) => {
    setSelectedChatUserId((prev) => (prev === musicianId ? null : musicianId));
  };

  return (
    <Layout>
      <Box maxWidth="700px" mx="auto" mt={5}>
        {error && <Typography color="error">{error}</Typography>}
        {!gig ? (
          <Typography>Loading gig...</Typography>
        ) : (
          <Paper elevation={3} sx={{ p: 4, borderRadius: 3 }}>
            <Typography variant="h4" gutterBottom>{gig.title}</Typography>
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
                  <Button onClick={handleDelete} style={{ backgroundColor: "#dc3545" }}>Delete Gig</Button>
                  <Button onClick={() => navigate(`/gigs/details/${id}/edit`)} style={{ backgroundColor: "#ffc107", color: "#000" }}>Edit Gig</Button>
                </>
              )}

              {!isOwner && gig.status === "Available" && !hasApplied && (
                <Button onClick={handleApply} style={{ backgroundColor: "#28a745", color: "#fff" }}>
                  Apply for Gig
                </Button>
              )}

              {!isOwner && hasApplied && (
                <Typography mt={2} color="primary">
                  You've already applied to this gig.
                </Typography>
              )}
            </Box>

            {!isOwner && hasApplied && chatExists && gig?.user?.id && (
              <Box mt={4}>
                <ChatView gigID={gig.id} otherUserID={gig.user.id} />
              </Box>
            )}

            {isOwner && applications.length > 0 && (
              <Box mt={5}>
                <Typography variant="h6" gutterBottom>Applications</Typography>
                {applications.map((app) => (
                  <Paper key={app.id} sx={{ p: 2, mb: 2, backgroundColor: "#f9f9f9" }}>
                    <Typography><strong>Name:</strong> {app.musician?.name}</Typography>
                    <Typography><strong>Email:</strong> {app.musician?.email}</Typography>
                    <Typography><strong>Status:</strong> {app.status}</Typography>
                    <Stack direction="row" spacing={2} mt={1}>
                      <Button onClick={() => handleAccept(app.musician_id)} style={{ backgroundColor: "#28a745" }}>Accept</Button>
                      <Button onClick={() => navigate(`/users/${app.musician_id}`)} style={{ backgroundColor: "#6c757d" }}>View Profile</Button>
                      <Button onClick={() => handleOpenChat(app.musician_id)} style={{ backgroundColor: "#0d6efd", color: "#fff" }}>💬 Chat</Button>
                    </Stack>
                    {selectedChatUserId === app.musician_id && (
                      <Box mt={2}>
                        <ChatView gigID={gig.id} otherUserID={app.musician_id} />
                      </Box>
                    )}
                  </Paper>
                ))}
              </Box>
            )}
          </Paper>
        )}
      </Box>
    </Layout>
  );
};

export default GigDetailsPage;
