import React, { useState } from "react";
import api from "../../api/axios";
import { useNavigate } from "react-router-dom";
import Layout from "../../components/Layout";
import FormInput from "../../components/FormInput";
import Button from "../../components/Button";
import { Typography, Box } from "@mui/material";

const LoginPage: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setMessage(null);
    setError(null);

    try {
      const res = await api.post("http://localhost:8080/auth/login", {
        email,
        password,
      });

      localStorage.setItem("token", res.data.token);
      setMessage("Login successful!");

      setTimeout(() => {
        navigate("/dashboard");
      }, 1000);
    } catch (err: any) {
      const msg = err.response?.data?.error || "Something went wrong";
      setError(msg);
    }
  };

  return (
    <Layout>
      <Box maxWidth="400px" mx="auto" mt={4}>
        <Typography variant="h4" gutterBottom align="center">
          Login
        </Typography>

        <form onSubmit={handleSubmit}>
          <FormInput
            type="email"
            label="Email"
            value={email}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setEmail(e.target.value)}
            required
          />
          <FormInput
            type="password"
            label="Password"
            value={password}
            onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setPassword(e.target.value)}
            required
          />
          <Button type="submit" fullWidth>
            Login
          </Button>
        </form>

        {message && <Typography color="success.main" mt={2}>{message}</Typography>}
        {error && <Typography color="error" mt={2}>{error}</Typography>}
      </Box>
    </Layout>
  );
};

export default LoginPage;
