import React from "react";

import "./App.css";

import { useState, useEffect } from "react";
import { BlogServiceClient } from "./protos/blog.client";
import { GrpcTransport } from "@protobuf-ts/grpc-transport";
import { ChannelCredentials } from "@grpc/grpc-js";

//const URL = "http://localhost:4444/ping"
function App() {
  const [time, setTime] = useState("");

  useEffect(() => {
    let gRPCTransport = new GrpcTransport({
      host: "localhost:5000",
      channelCredentials: ChannelCredentials.createInsecure(),
    });
    let myInterval: any = 3;
    const client = new BlogServiceClient(gRPCTransport);
    const fetchTime = async () => {
      let streamingCall = client.sendCurrentTime({ interval: myInterval });
      for await (let resp of streamingCall.responses) {
        console.log("got another hat! " + resp.currentTime);
        setTime(resp.currentTime)
      }
    };

    fetchTime();
  }, []);

  return <div className="App">Current Time: {time}</div>;
}

export default App;
