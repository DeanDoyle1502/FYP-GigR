import React, { useEffect, useState } from "react";
import api from "../../api/axios";
import Layout from "../../components/Layout";
import GigCard from "../../components/GigCard";
import { Gig } from "../../types/gig";

const MyGigsPage: React.FC = () => {
  const [gigs, setGigs] = useState<Gig[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      setError("You must be logged in.");
      return;
    }

    api
      .get("/gigs/mine", {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => setGigs(res.data))
      .catch((err) => {
        console.error("Failed to fetch gigs:", err);
        setError("Failed to load your gigs.");
      });
  }, []);

  return (
    <Layout>
      <div className="max-w-3xl mx-auto mt-8">
        <h2 className="text-2xl font-bold mb-6">My Gigs</h2>

        {error && <p className="text-red-500">{error}</p>}
        {!gigs.length && !error && (
          <p className="text-center text-gray-500">No gigs yet.</p>
        )}

        {gigs.map((gig) => (
          <GigCard key={gig.id} gig={gig} />
        ))}
      </div>
    </Layout>
  );
};

export default MyGigsPage;
