import React, { useEffect, useState } from "react";
import axios from "axios";

interface User {
  id: number;
  name: string;
  email: string;
  cognitoSub: string;
}

const Dashboard: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const [error, setError] = useState<string | null>(null);

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

        // Transform Go response to match our TypeScript interface
        const transformedUser: User = {
          id: raw.ID,
          name: raw.Name,
          email: raw.Email,
          cognitoSub: raw.CognitoSub,
        };

        setUser(transformedUser);
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
      <p>Your user ID: {user.id}</p>
      <p>Cognito Sub: {user.cognitoSub}</p>
    </div>
  );
};

export default Dashboard;
