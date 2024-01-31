import React from "react";
import ReactDOM from "react-dom/client";
import { SessionProvider } from "./contexts/session/index.ts";
import App from "./App.tsx";
import "./index.scss";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <SessionProvider>
      <App />
    </SessionProvider>
  </React.StrictMode>
);
