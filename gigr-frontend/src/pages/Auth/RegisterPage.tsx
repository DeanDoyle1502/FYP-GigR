import React, { useState } from "react";
import api from "../../api/axios";
import { useNavigate } from "react-router-dom";
import Layout from "../../components/Layout";
import FormInput from "../../components/FormInput";
import Button from "../../components/Button";
import MuiAlert from "@mui/material/Alert"; 
import { Container, Typography, Box } from "@mui/material";

const RegisterPage: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [instrument, setInstrument] = useState("");
  const [location, setLocation] = useState("");
  const [bio, setBio] = useState("");
  const [message, setMessage] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setMessage(null);
    setError(null);

    try {
      const res = await api.post("http://localhost:8080/auth/register", {
        email,
        password,
        name,
        instrument,
        location,
        bio,
      });

      setMessage(res.data.message);
      setTimeout(() => {
        navigate("/auth/confirm", { state: { email } });
      }, 2000);
    } catch (err: any) {
      const msg = err.response?.data?.error || "Something went wrong";
      setError(msg);
    }
  };

  return (
    <Layout>
      <Container maxWidth="sm">
        <Typography variant="h4" fontWeight="bold" textAlign="center" mb={3}>
          Register
        </Typography>
        <Box component="form" onSubmit={handleSubmit}>
          <FormInput
            label="Email"
            type="email"
            value={email}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setEmail(e.target.value)}
            required
          />
          <FormInput
            label="Password"
            type="password"
            value={password}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setPassword(e.target.value)}
            required
          />
          <FormInput
            label="Full Name"
            value={name}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setName(e.target.value)}
            required
          />
          <FormInput
            label="Instrument"
            value={instrument}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setInstrument(e.target.value)}
            required
          />
          <FormInput
            label="Location"
            value={location}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setLocation(e.target.value)}
            required
          />
          <FormInput
            label="Bio"
            multiline
            rows={3}
            value={bio}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setBio(e.target.value)}
            required
          />

          <Button type="submit" fullWidth sx={{ mt: 2 }}>
            Register
          </Button>
        </Box>

        {message && (
          <MuiAlert severity="success" sx={{ mt: 2 }}>
            {message}
          </MuiAlert>
        )}
        {error && (
          <MuiAlert severity="error" sx={{ mt: 2 }}>
            {error}
          </MuiAlert>
        )}
      </Container>
    </Layout>
  );
};

export default RegisterPage;
