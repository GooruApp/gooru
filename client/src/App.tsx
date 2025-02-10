import { useEffect, useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import { Greet } from "../wailsjs/go/main/App";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below 👇"
  );
  const [name, setName] = useState("");
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Greet(name).then(updateResultText);
  }
  const [backendResponse, setBackendResponse] = useState<string>("Pending...");

  useEffect(() => {
    window
      .fetch(`http://localhost:8000/`)
      .then((response) =>
        response.json().then((data) => setBackendResponse(data.message))
      );
  }, []);

  return (
    <>
      <div className="flex flex-row items-center justify-center">
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <span className="font-bold">Backend: </span>
        {backendResponse}
      </div>
      <div id="result" className="result">
        {resultText}
      </div>
      <div id="input" className="input-box flex gap-4 p-4">
        <input
          id="name"
          className="input border px-2"
          onChange={updateName}
          autoComplete="off"
          name="input"
          type="text"
        />
        <button className="btn" onClick={greet}>
          Greet
        </button>
      </div>
    </>
  );
}

export default App;
