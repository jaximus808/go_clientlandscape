import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import { HashRouter as Router, Routes, Route, Link } from "react-router-dom";

import Landing from "./views/Landing"


function App() {
    
    return (
        <div id='app'>
            <Router> 
                <Routes>
                    <Route path="/" element={<Landing />} />
                </Routes>
            </Router>
        </div>
    )
}

export default App
