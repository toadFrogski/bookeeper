import { FC, PropsWithChildren, useCallback, useEffect, useState } from "react";
import Context from "./Context";
import { useNavigate } from "react-router-dom";
import { Auth } from "../../services/api";

type Props = PropsWithChildren;

const Provider: FC<Props> = ({ children }) => {
  const [token, setToken] = useState<Auth>({token: ""});

  const navigate = useNavigate();

  const logout = useCallback(() => {
    setToken({token: ""});
    sessionStorage.removeItem("refreshToken");
    navigate("/");
  }, [navigate]);

  const setSessionToken = (token: string) => {
    setToken({token: token});
    sessionStorage.setItem("refreshToken", JSON.stringify({token: token}));
  };

  useEffect(() => {
    const token = sessionStorage.getItem("refreshToken");
    if (token !== null) {
      // const tokenObj = JSON.parse(token) as Auth;
      // if (Math.round(Date.now() / 1000) < tokenObj.expires) {
      //   AuthApiFactory().refresh(tokenObj.token)
      // }
     } else {
      logout();
    }
  }, [logout]);

  return (
    <Context.Provider
      value={{
        token: token,
        setToken: setSessionToken,
        logout: logout,
      }}
    >
      {children}
    </Context.Provider>
  );
};

export default Provider;
