import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import RegisterPage from './pages/Auth/RegisterPage';
import ConfirmPage from './pages/Auth/ConfirmPage';
import LoginPage from './pages/Auth/LoginPage';
import Dashboard from './pages/Dashboard';
import HomePage from './pages/HomePage';
import CreateGigPage from './pages/Gigs/CreateGigPage';
import MyGigsPage from './pages/Gigs/MyGigsPage';
import GigDetailPage from './pages/Gigs/GigDetailsPage';
import EditGigPage from './pages/Gigs/EditGigPage';
import PublicGigsPage from './pages/Gigs/PublicGigPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/confirm" element={<ConfirmPage/>}/>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/gigs/create" element={<CreateGigPage />} />
        <Route path="/gigs/mine" element={<MyGigsPage />} />
        <Route path="/gigs/:id" element={<GigDetailPage />} />
        <Route path="/gigs/:id/edit" element={<EditGigPage />} />
        <Route path="/gigs/public" element={<PublicGigsPage />} />
      </Routes>
    </Router>
  );
}

export default App;
