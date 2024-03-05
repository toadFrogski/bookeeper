import { FC, PropsWithChildren, useEffect, useRef, useState } from "react";
import Context from "./Context";

type Props = PropsWithChildren;

const Provider: FC<Props> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  const logout = () => {
    setIsAuthenticated(false);
  };

  const timer = useRef<number>();

  useEffect(() => {
    timer.current = setTimeout(() => {
      setIsAuthenticated(true);
      console.log(123);
    }, 2000);

    return () => clearTimeout(timer.current);
  }, [isAuthenticated]);

  return (
    <Context.Provider
      value={{
        isAuthenticated: isAuthenticated,
        logout: logout,
      }}
    >
      {children}
    </Context.Provider>
  );
};

export default Provider;
