import { FC, lazy } from "react";
import { LoginProvider } from "./contexts/login";
import { SessionProvider } from "./contexts/session";
import { BrowserRouter as Router } from "react-router-dom";

const Bookeeper = lazy(() => import("./Bookeeper/App"));

const App: FC = () => {
  return (
    <Router>
      <SessionProvider>
        <LoginProvider>
          <Bookeeper />
        </LoginProvider>
      </SessionProvider>
    </Router>
  );
};

export default App;
