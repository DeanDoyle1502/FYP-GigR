import React, { useState } from "react";
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";
import Layout from "../../components/Layout";
import FormInput from "../../components/FormInput";
import Button from "../../components/Button";
import { Typography, Box } from "@mui/material";

const ConfirmPage: React.FC = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const emailFromState = (location.state as { email: string })?.email || "";

  const [email, setEmail] = useState(emailFromState);
  const [code, setCode] = useState("");
  const [message, setMessage] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setMessage(null);
    setError(null);

    try {
      await axios.post("http://localhost:8080/auth/confirm", { email, code });
      setMessage("Account confirmed! You can now log in.");
      setTimeout(() => navigate("/auth/login"), 1500);
    } catch (err: any) {
      const msg = err.response?.data?.error || "Something went wrong";
      setError(msg);
    }
  };

  return (
    <Layout>
      <Box maxWidth="400px" mx="auto" mt={4}>
        <Typography variant="h4" gutterBottom align="center">
          Confirm Account
        </Typography>

        <form onSubmit={handleSubmit}>
          <FormInput
            label="Email"
            type="email"
            value={email}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setEmail(e.target.value)}
            required
          />
          <FormInput
            label="Confirmation Code"
            value={code}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setCode(e.target.value)}
            required
          />
          <Button type="submit" fullWidth>
            Confirm
          </Button>
        </form>

        {message && <Typography color="success.main" mt={2}>{message}</Typography>}
        {error && <Typography color="error" mt={2}>{error}</Typography>}
      </Box>
    </Layout>
  );
};

export default ConfirmPage;
