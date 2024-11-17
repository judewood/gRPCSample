import React from "react";

import "./App.css";

import { useState, useEffect } from "react";

const URL = "https://localhost:4444/ping"
function App() {
  const [time, setTime] = useState("")

  useEffect(() => {
    const fetchTime = async () => {
      const result = await fetch(URL);
      result.json().then((json) => {
        console.log(json.datetime);
        setTime(json.datetime);
      });
    };
    fetchTime();
  }, []);

  return <div className="App">Current Time: {time}</div>;
}

export default App;
