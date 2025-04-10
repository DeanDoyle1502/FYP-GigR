import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";

const EditGigPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [gig, setGig] = useState<any>({
    title: "",
    description: "",
    location: "",
    date: "",
    instrument: "",
  });
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      setError("You must be logged in to edit this gig.");
      return;
    }

    axios
      .get(`http://localhost:8080/gigs/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => {
        const gigData = res.data;
        setGig({
          title: gigData.title,
          description: gigData.description,
          location: gigData.location,
          date: new Date(gigData.date).toISOString().slice(0, 16), // Format for datetime-local input
          instrument: gigData.instrument,
        });
      })
      .catch((err) => {
        console.error("Error fetching gig:", err);
        setError("Could not load gig data for editing.");
      });
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    // Log the data being submitted to ensure it's correct
    console.log("Submitting gig:", gig);

    const token = localStorage.getItem("token");

    // Ensure the date is in the proper ISO format for backend storage
    const formattedDate = new Date(gig.date).toISOString();

    try {
      await axios.put(
        `http://localhost:8080/gigs/${id}`,
        {
          title: gig.title,
          description: gig.description,
          location: gig.location,
          date: formattedDate, // Send the formatted date
          instrument: gig.instrument,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );
      alert("Gig updated successfully!");
      navigate(`/gigs/${id}`);
    } catch (err: any) {
      console.error("Error updating gig:", err);
      setError("Failed to update gig.");
    }
  };

  if (error) return <p style={{ color: "red" }}>{error}</p>;

  return (
    <div style={{ maxWidth: "800px", margin: "2rem auto" }}>
      <h2>Edit Gig</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Title"
          value={gig.title}
          onChange={(e) => setGig({ ...gig, title: e.target.value })}
          required
          style={inputStyle}
        />
        <textarea
          placeholder="Description"
          value={gig.description}
          onChange={(e) => setGig({ ...gig, description: e.target.value })}
          required
          style={{ ...inputStyle, height: "100px" }}
        />
        <input
          type="text"
          placeholder="Location"
          value={gig.location}
          onChange={(e) => setGig({ ...gig, location: e.target.value })}
          required
          style={inputStyle}
        />
        <input
          type="datetime-local"
          value={gig.date}
          onChange={(e) => setGig({ ...gig, date: e.target.value })}
          required
          style={inputStyle}
        />
        <input
          type="text"
          placeholder="Instrument Needed"
          value={gig.instrument}
          onChange={(e) => setGig({ ...gig, instrument: e.target.value })}
          required
          style={inputStyle}
        />
        <button type="submit" style={buttonStyle}>
          Update Gig
        </button>
      </form>
    </div>
  );
};

const inputStyle: React.CSSProperties = {
  width: "100%",
  padding: "0.5rem",
  marginBottom: "1rem",
  borderRadius: "4px",
  border: "1px solid #ccc",
};

const buttonStyle: React.CSSProperties = {
  padding: "0.6rem 1.2rem",
  backgroundColor: "#007bff",
  color: "#fff",
  border: "none",
  borderRadius: "4px",
  cursor: "pointer",
};

export default EditGigPage;
