import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";

interface Gig {
  id: number;
  title: string;
  description: string;
  location: string;
  date: string;
  instrument: string;
  status: string;
}

const GigDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [gig, setGig] = useState<Gig | null>(null);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      setError("You must be logged in to view this gig.");
      return;
    }

    axios
      .get(`http://localhost:8080/gigs/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => setGig(res.data))
      .catch((err) => {
        console.error("Error fetching gig:", err);
        setError("Could not load gig.");
      });
  }, [id]);

  const handleDelete = async () => {
    const confirmDelete = window.confirm("Are you sure you want to delete this gig?");
    if (!confirmDelete) return;

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

  if (error) return <p style={{ color: "red" }}>{error}</p>;
  if (!gig) return <p>Loading gig...</p>;

  return (
    <div style={{ maxWidth: "800px", margin: "2rem auto" }}>
      <h2>{gig.title}</h2>
      <p><strong>Date:</strong> {new Date(gig.date).toLocaleString()}</p>
      <p><strong>Location:</strong> {gig.location}</p>
      <p><strong>Instrument Needed:</strong> {gig.instrument}</p>
      <p><strong>Status:</strong> {gig.status}</p>
      <p>{gig.description}</p>

      <div style={{ marginTop: "2rem" }}>
        <button onClick={() => navigate("/dashboard")} style={buttonStyle}>Back to Dashboard</button>
        <button onClick={() => navigate("/gigs/mine")} style={{ ...buttonStyle, marginLeft: "1rem", backgroundColor: "#007bff" }}>
          View My Gigs
        </button>
        <button onClick={handleDelete} style={{ ...buttonStyle, marginLeft: "1rem", backgroundColor: "#dc3545" }}>
          Delete Gig
        </button>
        <button onClick={() => navigate(`/gigs/${id}/edit`)} style={{ ...buttonStyle, marginLeft: "1rem", backgroundColor: "#ffc107", color: "#000" }}>
          Edit Gig
        </button>
      </div>
    </div>
  );
};

const buttonStyle: React.CSSProperties = {
  padding: "0.5rem 1rem",
  backgroundColor: "#28a745",
  color: "white",
  border: "none",
  borderRadius: "4px",
  cursor: "pointer",
};

export default GigDetailPage;
