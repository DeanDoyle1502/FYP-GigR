import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import api from "../../api/axios"
import Layout from "../../components/Layout";
import FormInput from "../../components/FormInput";
import Button from "../../components/Button";
import {
  Box,
  Typography,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  SelectChangeEvent,
} from "@mui/material";

const EditGigPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const [gig, setGig] = useState({
    title: "",
    description: "",
    location: "",
    date: "",
    instrument: "",
    status: "Available",
  });

  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      setError("You must be logged in to edit this gig.");
      return;
    }

    api.get(`http://localhost:8080/gigs/details/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => {
        const gigData = res.data;
        setGig({
          title: gigData.title,
          description: gigData.description,
          location: gigData.location,
          date: new Date(gigData.date).toISOString().slice(0, 16),
          instrument: gigData.instrument,
          status: gigData.status || "Available",
        });
      })
      .catch((err) => {
        console.error("Error fetching gig:", err);
        setError("Could not load gig data.");
      });
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const token = localStorage.getItem("token");
    const formattedDate = new Date(gig.date).toISOString();

    try {
      await api.put(
        `http://localhost:8080/gigs/details/${id}`,
        { ...gig, date: formattedDate },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      alert("Gig updated successfully!");
      navigate(`/gigs/details/${id}`);
    } catch (err: any) {
      console.error("Error updating gig:", err);
      setError("Failed to update gig.");
    }
  };

  return (
    <Layout>
      <Box maxWidth="600px" mx="auto" mt={4}>
        <Typography variant="h4" gutterBottom>
          Edit Gig
        </Typography>

        <form onSubmit={handleSubmit}>
          <FormInput
            label="Title"
            value={gig.title}
            onChange={(e: { target: { value: any; }; }) => setGig({ ...gig, title: e.target.value })}
            required
          />
          <FormInput
            label="Description"
            value={gig.description}
            onChange={(e: { target: { value: any; }; }) => setGig({ ...gig, description: e.target.value })}
            multiline
            rows={3}
            required
          />
          <FormInput
            label="Location"
            value={gig.location}
            onChange={(e: { target: { value: any; }; }) => setGig({ ...gig, location: e.target.value })}
            required
          />
          <FormInput
            type="datetime-local"
            label="Date"
            InputLabelProps={{ shrink: true }}
            value={gig.date}
            onChange={(e: { target: { value: any; }; }) => setGig({ ...gig, date: e.target.value })}
            required
          />
          <FormInput
            label="Instrument"
            value={gig.instrument}
            onChange={(e: { target: { value: any; }; }) => setGig({ ...gig, instrument: e.target.value })}
            required
          />
          <FormControl fullWidth margin="normal">
            <InputLabel>Status</InputLabel>
            <Select
              label="Status"
              value={gig.status}
              onChange={(e: SelectChangeEvent) =>
                setGig({ ...gig, status: e.target.value })
              }
            >
              <MenuItem value="Available">Available</MenuItem>
              <MenuItem value="Pending">Pending</MenuItem>
              <MenuItem value="Covered">Covered</MenuItem>
            </Select>
          </FormControl>

          <Button type="submit" fullWidth>
            Update Gig
          </Button>
        </form>

        {error && <Typography color="error" mt={2}>{error}</Typography>}
      </Box>
    </Layout>
  );
};

export default EditGigPage;
