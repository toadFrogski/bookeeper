import { FC, PropsWithChildren, ReactNode, useEffect, useState } from "react";
import Context from "./Context";
import { Auth } from "../../services/api";

type Props = PropsWithChildren & {
  fallback?: ReactNode;
};

const Provider: FC<Props> = ({ children, fallback }) => {
  const [token, setToken] = useState<Auth>({ token: "" });
  const [isLoading, setIsLoading] = useState(true);

  const logout = () => {
    setToken({ token: "" });
    localStorage.removeItem("accessToken");
  };

  const setSessionToken = (token: Auth) => {
    setToken(token);
    localStorage.setItem("accessToken", JSON.stringify(token));
  };

  useEffect(() => {
    (async () => {
      const token = localStorage.getItem("accessToken");
      if (token !== null) {
        const tokenObj = JSON.parse(token) as Auth;
        setToken(tokenObj);
        // if (Math.round(Date.now() / 1000) < tokenObj.expires) {
        //   AuthApiFactory().refresh(tokenObj.token)
        // }
      } else {
        logout();
      }
      setIsLoading(false);
    })();
  }, []);

  return (
    <Context.Provider
      value={{
        token: token,
        setToken: setSessionToken,
        logout: logout,
      }}
    >
      {isLoading ? fallback : children}
    </Context.Provider>
  );
};

export default Provider;
