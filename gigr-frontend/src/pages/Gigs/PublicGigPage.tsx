import React, { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

interface Gig {
  id: number;
  title: string;
  description: string;
  location: string;
  date: string;
  instrument: string;
  status: string;
}

const PublicGigsPage: React.FC = () => {
  const [gigs, setGigs] = useState<Gig[]>([]);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    axios
      .get("http://localhost:8080/gigs/public")
      .then((res) => setGigs(res.data))
      .catch((err) => {
        console.error("Failed to fetch public gigs:", err);
        setError("Could not load public gigs.");
      });
  }, []);

  if (error) return <p style={{ color: "red" }}>{error}</p>;
  if (!gigs.length) return <p style={{ textAlign: "center" }}>No gigs currently available.</p>;

  return (
    <div style={{ maxWidth: "800px", margin: "2rem auto" }}>
      <h2>Public Gigs</h2>
      {gigs.map((gig) => (
        <div
          key={gig.id}
          onClick={() => navigate(`/gigs/${gig.id}`)}
          style={{
            border: "1px solid #ccc",
            padding: "1rem",
            marginBottom: "1rem",
            borderRadius: "6px",
            cursor: "pointer",
            transition: "background 0.2s",
          }}
          onMouseOver={(e) => (e.currentTarget.style.backgroundColor = "#f8f9fa")}
          onMouseOut={(e) => (e.currentTarget.style.backgroundColor = "white")}
        >
          <h3>{gig.title}</h3>
          <p><strong>Date:</strong> {new Date(gig.date).toLocaleString()}</p>
          <p><strong>Location:</strong> {gig.location}</p>
          <p><strong>Instrument:</strong> {gig.instrument}</p>
          <p>{gig.description}</p>
        </div>
      ))}
    </div>
  );
};

export default PublicGigsPage;
