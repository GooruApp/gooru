import { useEffect, useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

function App() {
  const [backendResponse, setBackendResponse] = useState<string>("Pending...");

  useEffect(() => {
    window.fetch(`http://localhost:8000/`).then(
      (response) => response.json().then((data) => setBackendResponse(data.message))
    );
  }, [])

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
        <span className="font-bold">Backend: </span>{backendResponse}
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
