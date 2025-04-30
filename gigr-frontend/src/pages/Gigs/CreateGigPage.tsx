import React, { useState } from "react";
import api from "../../api/axios";
import { useNavigate } from "react-router-dom";
import Layout from "../../components/Layout";
import FormInput from "../../components/FormInput";
import Button from "../../components/Button";
import {
  Box,
  MenuItem,
  Typography,
  Select,
  SelectChangeEvent,
  FormControl,
  InputLabel,
} from "@mui/material";

const CreateGigPage: React.FC = () => {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [location, setLocation] = useState("");
  const [date, setDate] = useState("");
  const [instrument, setInstrument] = useState("");
  const [status, setStatus] = useState("Available");
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);

    const token = localStorage.getItem("token");
    const formattedDate = new Date(date).toISOString();

    try {
      await api.post(
        "http://localhost:8080/gigs/",
        { title, description, location, date: formattedDate, instrument, status },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      setSuccess("Gig created successfully!");
      setTimeout(() => navigate("/gigs/mine"), 1500);
    } catch (err: any) {
      console.error("Gig creation failed:", err);
      setError(err.response?.data?.error || "Something went wrong.");
    }
  };

  return (
    <Layout>
      <Box maxWidth="600px" mx="auto" mt={4}>
        <Typography variant="h4" gutterBottom>
          Create a New Gig
        </Typography>

        <form onSubmit={handleSubmit}>
          <FormInput
            label="Title"
            value={title}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setTitle(e.target.value)}
            required
          />
          <FormInput
            label="Description"
            value={description}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setDescription(e.target.value)}
            multiline
            rows={3}
            required
          />
          <FormInput
            label="Location"
            value={location}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setLocation(e.target.value)}
            required
          />
          <FormInput
            type="datetime-local"
            label="Date"
            InputLabelProps={{ shrink: true }}
            value={date}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setDate(e.target.value)}
            required
          />
          <FormInput
            label="Instrument Needed"
            value={instrument}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setInstrument(e.target.value)}
            required
          />
          <FormControl fullWidth margin="normal">
            <InputLabel>Status</InputLabel>
            <Select
              label="Status"
              value={status}
              onChange={(e: SelectChangeEvent) => setStatus(e.target.value)}
            >
              <MenuItem value="Available">Available</MenuItem>
              <MenuItem value="Pending">Pending</MenuItem>
              <MenuItem value="Covered">Covered</MenuItem>
            </Select>
          </FormControl>

          <Button type="submit" fullWidth>
            Create Gig
          </Button>
        </form>

        {success && <Typography color="success.main" mt={2}>{success}</Typography>}
        {error && <Typography color="error" mt={2}>{error}</Typography>}
      </Box>
    </Layout>
  );
};

export default CreateGigPage;
