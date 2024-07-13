import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';
import List from './components/List';
import ItemDetail from './components/ItemDetail';

function App() {
    return (
        <Router>
            <div className="App">
                <header className="App-header">
                    <h1>My JSON List App</h1>
                </header>
                <Routes>
                    <Route path="/" element={<List />} />
                    <Route path="/item/:id" element={<ItemDetail />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
