import React from 'react';
import { useNavigate } from 'react-router-dom';

const HomePage: React.FC = () => {
  const navigate = useNavigate();

  return (
    <div style={styles.container}>
      <h1>Welcome to GigR 🎸</h1>
      <p>Find and create music gigs with ease.</p>

      <div style={styles.buttonContainer}>
        <button onClick={() => navigate('/login')} style={styles.button}>
          Login
        </button>
        <button onClick={() => navigate('/register')} style={{ ...styles.button, backgroundColor: '#28a745' }}>
          Sign Up
        </button>
      </div>
    </div>
  );
};

const styles = {
  container: {
    textAlign: 'center',
    padding: '3rem',
  },
  buttonContainer: {
    display: 'flex',
    justifyContent: 'center',
    gap: '1rem',
    marginTop: '2rem',
  },
  button: {
    padding: '0.75rem 1.5rem',
    fontSize: '1rem',
    borderRadius: '5px',
    border: 'none',
    cursor: 'pointer',
    backgroundColor: '#007bff',
    color: 'white',
  },
} as const;

export default HomePage;
