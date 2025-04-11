import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom"; // Import Link for navigation

interface Gig {
  id: number;
  title: string;
  description: string;
  location: string;
  date: string;
  instrument: string;
  status: string;
}

const MyGigsPage: React.FC = () => {
  const [gigs, setGigs] = useState<Gig[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      setError("You must be logged in.");
      return;
    }

    axios
      .get("http://localhost:8080/gigs/mine", {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => {
        setGigs(res.data);
      })
      .catch((err) => {
        console.error("Failed to fetch gigs:", err);
        setError("Failed to load your gigs.");
      });
  }, []);

  if (error) return <p style={{ color: "red" }}>{error}</p>;
  if (!gigs.length) return <p style={{ textAlign: "center" }}>No gigs yet.</p>;

  return (
    <div style={{ maxWidth: "800px", margin: "2rem auto" }}>
      <h2>My Gigs</h2>
      {gigs.map((gig) => (
        <Link
          key={gig.id}
          to={`/gigs/${gig.id}`} // Navigates to the gig details page
          style={{
            textDecoration: "none",
            color: "inherit", // Inherit text color from parent
          }}
        >
          <div
            style={{
              border: "1px solid #ccc",
              padding: "1rem",
              marginBottom: "1rem",
              borderRadius: "6px",
              cursor: "pointer", // Show pointer cursor to indicate clickability
              transition: "background-color 0.3s", // For smooth hover transition
            }}
            onMouseEnter={(e) => (e.currentTarget.style.backgroundColor = "#f0f0f0")}
            onMouseLeave={(e) => (e.currentTarget.style.backgroundColor = "white")}
          >
            <h3>{gig.title}</h3>
            <p><strong>Date:</strong> {new Date(gig.date).toLocaleString()}</p>
            <p><strong>Location:</strong> {gig.location}</p>
            <p><strong>Instrument:</strong> {gig.instrument}</p>
            <p><strong>Status:</strong> {gig.status}</p>
            <p>{gig.description}</p>
          </div>
        </Link>
      ))}
    </div>
  );
};

export default MyGigsPage;
