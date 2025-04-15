import React, { useEffect, useState } from "react";
import axios from "axios";
import Layout from "../../components/Layout";
import GigCard from "../../components/GigCard";
import { Gig } from "../../types/gig";
import { Box, Typography } from "@mui/material";

const PublicGigsPage: React.FC = () => {
  const [gigs, setGigs] = useState<Gig[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    axios
      .get("http://localhost:8080/gigs/public")
      .then((res) => setGigs(res.data))
      .catch((err) => {
        console.error("Failed to fetch public gigs:", err);
        setError("Could not load public gigs.");
      });
  }, []);

  return (
    <Layout>
      <Box maxWidth="800px" mx="auto" mt={4}>
        <Typography variant="h4" gutterBottom>
          Public Gigs
        </Typography>

        {error && <Typography color="error">{error}</Typography>}
        {!gigs.length && !error && (
          <Typography color="textSecondary" align="center">
            No gigs currently available.
          </Typography>
        )}

        {gigs.map((gig) => (
          <GigCard key={gig.id} gig={gig} />
        ))}
      </Box>
    </Layout>
  );
};

export default PublicGigsPage;
