// src/components/NavBar.tsx
import React from "react";
import { AppBar, Toolbar, Typography, Button, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";

const NavBar: React.FC = () => {
  const navigate = useNavigate();
  const isLoggedIn = !!localStorage.getItem("token");

  return (
    <AppBar position="static" color="primary">
      <Toolbar>
        <Typography
          variant="h6"
          sx={{ flexGrow: 1, cursor: "pointer" }}
          onClick={() => navigate("/")}
        >
          GigR
        </Typography>

        {isLoggedIn ? (
          <>
            <Button color="inherit" onClick={() => navigate("/dashboard")}>
              Dashboard
            </Button>
            <Button color="inherit" onClick={() => navigate("/gigs/mine")}>
              My Gigs
            </Button>
            <Button color="inherit" onClick={() => navigate("/gigs/applications")}>
              My Applications
            </Button>
            <Button color="inherit" onClick={() => navigate("/gigs/public")}>
              Public Gigs
            </Button>
            <Button
              color="inherit"
              onClick={() => {
                localStorage.removeItem("token");
                navigate("/");
              }}
            >
              Logout
            </Button>
          </>
        ) : (
          <>
            <Button color="inherit" onClick={() => navigate("/auth/login")}>
              Login
            </Button>
            <Button color="inherit" onClick={() => navigate("/auth/register")}>
              Register
            </Button>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default NavBar;
