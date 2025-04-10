import React, { useEffect, useState } from "react";
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
        <div key={gig.id} style={{ border: "1px solid #ccc", padding: "1rem", marginBottom: "1rem", borderRadius: "6px" }}>
          <h3>{gig.title}</h3>
          <p><strong>Date:</strong> {new Date(gig.date).toLocaleString()}</p>
          <p><strong>Location:</strong> {gig.location}</p>
          <p><strong>Instrument:</strong> {gig.instrument}</p>
          <p><strong>Status:</strong> {gig.status}</p>
          <p>{gig.description}</p>
        </div>
      ))}
    </div>
  );
};

export default MyGigsPage;
