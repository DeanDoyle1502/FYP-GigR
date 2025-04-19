import { Card, CardContent, Typography, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { Gig } from "../types/gig";

const GigCard = ({ gig, disableNavigation = false }: { gig: Gig; disableNavigation?: boolean }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    if (!disableNavigation) {
      navigate(`/gigs/details/${gig.id}`);
    }
  };

  return (
    <Card
      onClick={handleClick}
      sx={{
        mb: 2,
        cursor: disableNavigation ? "default" : "pointer",
        "&:hover": {
          backgroundColor: disableNavigation ? "inherit" : "#f5f5f5",
        },
      }}
      variant="outlined"
    >
      <CardContent>
        <Typography variant="h6" component="div">
          {gig.title}
        </Typography>

        <Box mt={1}>
          <Typography variant="body2" color="text.secondary">
            <strong>Date:</strong> {new Date(gig.date).toLocaleString()}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            <strong>Location:</strong> {gig.location}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            <strong>Instrument:</strong> {gig.instrument}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            <strong>Status:</strong> {gig.status}
          </Typography>
        </Box>

        <Typography variant="body1" mt={2}>
          {gig.description}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default GigCard;
