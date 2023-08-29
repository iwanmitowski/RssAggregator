import './App.css';
import { Route, Routes } from 'react-router';

import Home  from './components/Shared/Home';
import Main from './components/Shared/Main/Main';
import Register from './components/Auth/Register/Register';
import NoAuthGuard from './components/Shared/Guards/NoAuthGuard';

function App() {
  return (
    <div className="App">
      <Main>
        <Routes>
          <Route path="/" element={<Home/>} /> 
          <Route path="/register" element={
            <NoAuthGuard>
              <Register />
            </NoAuthGuard>
            } 
          />
        </Routes>
      </Main>
    </div>
  );
}

export default App;
