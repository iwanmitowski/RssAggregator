import './App.css';
import { Route, Routes } from 'react-router';

import Home  from './components/Shared/Home';
import Main from './components/Shared/Main/Main';
import Register from './components/Auth/Register/Register';
import NoAuthGuard from './components/Shared/Guards/NoAuthGuard';
import Header from './components/Shared/Header/Header';
import FeedForm from './components/Feed/FeedForm';
import AuthGuard from './components/Shared/Guards/AuthGuard';

function App() {
  return (
    <div className="App">
      <Header />
      <Main>
        <Routes>
          <Route path="/" element={<Home/>} /> 
          <Route path="/register" element={
            <NoAuthGuard>
              <Register />
            </NoAuthGuard>
            } 
          />
          <Route path="/feed" element={
            <AuthGuard>
              <FeedForm />
            </AuthGuard>
            } 
          />
        </Routes>
      </Main>
    </div>
  );
}

export default App;
