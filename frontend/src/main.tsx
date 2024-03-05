import React, { lazy } from "react";
import ReactDOM from "react-dom/client";
import "./index.scss";
import { SessionProvider } from "./contexts/session";

const Bookeeper = lazy(() => import("./Bookeeper/App"));

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <SessionProvider>
      <Bookeeper />
    </SessionProvider>
  </React.StrictMode>
);
