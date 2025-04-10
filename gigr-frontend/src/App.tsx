import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import RegisterPage from './pages/RegisterPage';
import ConfirmPage from './pages/ConfirmPage';
import LoginPage from './pages/LoginPage';
import Dashboard from './pages/Dashboard';
import HomePage from './pages/HomePage';
import CreateGigPage from './pages/CreateGigPage';
import MyGigsPage from './pages/MyGigsPage';
import GigDetailPage from './pages/GigDetailsPage';
import EditGigPage from './pages/EditGigPage';

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
      </Routes>
    </Router>
  );
}

export default App;
