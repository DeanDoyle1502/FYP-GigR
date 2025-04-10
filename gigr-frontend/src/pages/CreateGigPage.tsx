import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const CreateGigPage: React.FC = () => {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [location, setLocation] = useState("");
  const [date, setDate] = useState("");
  const [instrument, setInstrument] = useState("");
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
      const res = await axios.post(
        "http://localhost:8080/gigs/",
        {
          title,
          description,
          location,
          date: formattedDate, // fix: rename to "date"
          instrument,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
  
      const createdGigId = res.data.gig.id;
  
      setSuccess("Gig created successfully!");
      setTimeout(() => navigate(`/gigs/${createdGigId}`), 1500);
    } catch (err: any) {
      console.error("Gig creation failed:", err);
      setError(err.response?.data?.error || "Something went wrong.");
    }
  };
  

  return (
    <div style={{ maxWidth: "600px", margin: "2rem auto" }}>
      <h2>Create a New Gig</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" placeholder="Title" value={title} onChange={(e) => setTitle(e.target.value)} required style={inputStyle} />
        <textarea placeholder="Description" value={description} onChange={(e) => setDescription(e.target.value)} required style={{ ...inputStyle, height: "100px" }} />
        <input type="text" placeholder="Location" value={location} onChange={(e) => setLocation(e.target.value)} required style={inputStyle} />
        <input type="datetime-local" value={date} onChange={(e) => setDate(e.target.value)} required style={inputStyle} />
        <input type="text" placeholder="Instrument Needed" value={instrument} onChange={(e) => setInstrument(e.target.value)} required style={inputStyle} />
        <button type="submit" style={buttonStyle}>Create Gig</button>
      </form>

      {success && <p style={{ color: "green" }}>{success}</p>}
      {error && <p style={{ color: "red" }}>{error}</p>}
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

export default CreateGigPage;
