import { useState } from 'react';
import logo from '../assets/images/palantir-logo.png';
import '../App.css';
import { Greet } from "../../wailsjs/go/main/App";

function Landing() {
    
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);

    return (
        <div className='flex flex-col items-center justify-center min-h-screen  text-center'>
            <div className='w-24 h-24 mb-8'>
                <img src={logo} alt="logo" />
            </div>
            <div className='text-2xl font-light'>Welcome to 
                <strong className='font-extrabold'> Client Landscape</strong>
            </div>

            <div id="result" className="result mb-2 text-lg">Enter a transcript to get started</div>
            <div id="input" className=" flex gap-2"> 
                <input
                    id="name"
                    className="input px-2 py-1 border rounded bg-white text-black"
                    onChange={updateName}
                    autoComplete="off"
                    name="input"
                    type="text"
                />
                <button className="btn px-4 py-1 bg-blue-500 text-white rounded hover:bg-blue-600" >
                    Upload Transcript
                </button>
            </div>
        </div>
    );
}

export default Landing;
