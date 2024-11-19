import "./App.css";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { useState, useEffect } from "react";
import { BlogServiceClient } from "./protos/blog.client";

const URL = "http://localhost:4444";

function App() {
  const [time, setTime] = useState("");

  const transport = new GrpcWebFetchTransport({
    baseUrl: URL,
  });
  const client = new BlogServiceClient(transport);

  useEffect(() => {
    const fetchTime = async () => {
      const call = await client.sendTimeOne({
        interval: 23,
      });

      console.log(`### calling method "${call.method.name}"...`);

      const headers = await call.headers;
      console.log("got response headers: ", headers);

      const response = await call.response;
      console.log("got response message: ", response);

      const status = await call.status;
      console.log("got status: ", status);

      const trailers = await call.trailers;
      console.log("got trailers: ", trailers);
      setTime("test");
    };
    fetchTime();
  }, []);

  return <div className="App">Current Time: {time}</div>;
}

export default App;
