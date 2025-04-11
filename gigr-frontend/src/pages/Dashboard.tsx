import React, { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

interface User {
  id: number;
  name: string;
  email: string;
  instrument: string;
  location: string;
  bio: string;
}

const Dashboard: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      setError("You must be logged in.");
      return;
    }

    axios
      .get("http://localhost:8080/auth/me", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => {
        const raw = res.data;

        const normalizedUser = {
          id: raw.ID,
          name: raw.Name,
          email: raw.Email,
          instrument: raw.Instrument,
          location: raw.Location,
          bio: raw.Bio,
        };

        setUser(normalizedUser);
      })
      .catch((err) => {
        console.error("Failed to fetch user info:", err);
        setError("Failed to load user profile.");
      });
  }, []);

  if (error) return <p style={{ color: "red" }}>{error}</p>;
  if (!user) return <p>Loading...</p>;

  return (
    <div style={{ maxWidth: "600px", margin: "2rem auto", textAlign: "center" }}>
      <h1>Welcome back, {user.name || user.email}!</h1>
      <p>Email: {user.email}</p>
      <p>Instrument: {user.instrument}</p>
      <p>Location: {user.location}</p>
      <p>Bio: {user.bio}</p>

      <button
        style={buttonStyle}
        onClick={() => navigate("/gigs/create")}
      >
        Create a New Gig
      </button>

      <button
        style={{ ...buttonStyle, backgroundColor: "#007bff", marginTop: "1rem" }}
        onClick={() => navigate("/gigs/mine")}
      >
        View My Gigs
      </button>
      <button
        onClick={() => navigate("/gigs/public")}
        style={{ ...buttonStyle, backgroundColor: "#6f42c1", marginLeft: "1rem" }}
      >
        Browse Available Gigs
</button>
    </div>
  );
};

const buttonStyle: React.CSSProperties = {
  marginTop: "2rem",
  padding: "0.6rem 1.2rem",
  fontSize: "1rem",
  cursor: "pointer",
  borderRadius: "4px",
  border: "none",
  backgroundColor: "#28a745",
  color: "white",
};

export default Dashboard;
