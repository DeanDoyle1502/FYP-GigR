import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const RegisterPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [name, setName] = useState('');
  const [instrument, setInstrument] = useState('');
  const [location, setLocation] = useState('');
  const [bio, setBio] = useState('');
  const [message, setMessage] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setMessage(null);
    setError(null);

    try {
      const res = await axios.post('http://localhost:8080/auth/register', {
        email,
        password,
        name,
        instrument,
        location,
        bio,
      });

      setMessage(res.data.message);
      setTimeout(() => {
        navigate('/confirm', { state: { email } });
      }, 2000);
    } catch (err: any) {
      const msg = err.response?.data?.error || 'Something went wrong';
      setError(msg);
    }
  };

  return (
    <div style={{ maxWidth: '500px', margin: 'auto', padding: '2rem' }}>
      <h2>Register</h2>
      <form onSubmit={handleSubmit}>
        <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} required style={inputStyle} />
        <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} required style={inputStyle} />
        <input type="text" placeholder="Full Name" value={name} onChange={(e) => setName(e.target.value)} required style={inputStyle} />
        <input type="text" placeholder="Instrument" value={instrument} onChange={(e) => setInstrument(e.target.value)} required style={inputStyle} />
        <input type="text" placeholder="Location" value={location} onChange={(e) => setLocation(e.target.value)} required style={inputStyle} />
        <textarea placeholder="Bio" value={bio} onChange={(e) => setBio(e.target.value)} required style={{ ...inputStyle, height: '80px' }} />

        <button type="submit" style={buttonStyle}>Register</button>
      </form>

      {message && <p style={{ color: 'green', marginTop: '1rem' }}>{message}</p>}
      {error && <p style={{ color: 'red', marginTop: '1rem' }}>{error}</p>}
    </div>
  );
};

const inputStyle: React.CSSProperties = {
  width: '100%',
  padding: '0.5rem',
  marginBottom: '1rem',
  borderRadius: '4px',
  border: '1px solid #ccc',
};

const buttonStyle: React.CSSProperties = {
  width: '100%',
  padding: '0.6rem',
  backgroundColor: '#007bff',
  color: 'white',
  border: 'none',
  borderRadius: '4px',
  cursor: 'pointer',
};

export default RegisterPage;
