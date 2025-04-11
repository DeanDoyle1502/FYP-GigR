import { useNavigate } from "react-router-dom";

const HomePage: React.FC = () => {
  const navigate = useNavigate();

  return (
    <div style={{ textAlign: "center", marginTop: "3rem" }}>
      <h1>Welcome to GigR 🎸</h1>
      <p>Connect musicians with acts & bands near you.</p>

      <div style={{ marginTop: "2rem" }}>
        <button
          onClick={() => navigate("/auth/login")}
          style={buttonStyle}
        >
          Log In
        </button>
        <button
          onClick={() => navigate("/auth/register")}
          style={{ ...buttonStyle, marginLeft: "1rem" }}
        >
          Create Profile
        </button>
        <button
          onClick={() => navigate("/gigs/public")}
          style={{ ...buttonStyle, marginLeft: "1rem", backgroundColor: "#6f42c1" }}
        >
          Browse Gigs
        </button>
      </div>
    </div>
  );
};

const buttonStyle: React.CSSProperties = {
  padding: "0.6rem 1.2rem",
  fontSize: "1rem",
  backgroundColor: "#007bff",
  color: "white",
  border: "none",
  borderRadius: "4px",
  cursor: "pointer",
};

export default HomePage;
