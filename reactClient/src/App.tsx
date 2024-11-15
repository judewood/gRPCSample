import React from "react";
import logo from "./logo.svg";
import "./App.css";

import { useState, useEffect } from "react";

const URL = "http://worldtimeapi.org/api/timezone/Europe/London";
function App() {
  let currentTime = "nowby";
  const [time, setTime] = useState("")

  useEffect(() => {
    const fetchTime = async () => {
      const result = await fetch(URL);
      result.json().then((json) => {
        console.log();
        setTime(json.datetime);
      });
    };
    fetchTime();
  }, []);

  return <div className="App">Current Time: {time}</div>;
}

export default App;
