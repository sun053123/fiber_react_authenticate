import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { Suspense, lazy,  } from 'react';

import Navbar from './components/navbar/navbar';
import Login from './layouts/LoginLayout/LoginLayout';
import Register from './layouts/RegisterLayout/RegisterLayout';
import Home from './layouts/HomeLayout/HomeLayout';
// const Login = lazy(()  => import('./layouts/LoginLayout/LoginLayout'))
// const Home = lazy(() => import('./layouts/HomeLayout/HomeLayout'))
// const Register = lazy(() => import('./layouts/RegisterLayout/RegisterLayout'))

function App() {

  return (
    <Router >
      <Suspense fallback={<div>Loading...</div>}>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </Suspense>
    </Router>

  );
}

export default App;
