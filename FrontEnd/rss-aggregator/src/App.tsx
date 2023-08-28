import './App.css';
import { Route, Routes } from 'react-router';

import Home  from './components/Home';
import Main from './components/Shared/Main/Main';

function App() {
  return (
    <div className="App">
      <Main>
        <Routes>
          <Route path="/" element={<Home/>} /> 
        </Routes>
      </Main>
    </div>
  );
}

export default App;
