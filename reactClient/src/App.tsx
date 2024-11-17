import "./App.css";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { useState, useEffect } from "react";
import { BlogServiceClient, IBlogServiceClient } from "./protos/blog.client";

const URL = "https://localhost:4444";

async function callServerStream(client: IBlogServiceClient) {
  const call = client.sendCurrentTime({
    interval: 3,
  });

  console.log(`### calling method "${call.method.name}"...`);

  const headers = await call.headers;
  console.log("got response headers: ", headers);

  for await (let response of call.responses) {
    console.log("got response message: ", response);
  }

  const status = await call.status;
  console.log("got status: ", status);

  const trailers = await call.trailers;
  console.log("got trailers: ", trailers);

  console.log();
}

function App() {
  const [time, setTime] = useState("");

  const transport = new GrpcWebFetchTransport({
    baseUrl: URL,
    
  });
  const client = new BlogServiceClient(transport);

  useEffect(() => {
    const fetchTime = async () => {
      await callServerStream(client);
      setTime("jude");
    };
    // const fetchTime = async () => {
    //   const result = await fetch(URL);
    //   result.json().then((json) => {
    //     console.log(json.datetime);
    //     setTime(json.datetime);
    //   });
    // };
    fetchTime();
  }, []);

  return <div className="App">Current Time: {time}</div>;
}

export default App;
