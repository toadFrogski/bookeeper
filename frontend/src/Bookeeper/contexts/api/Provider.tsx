import { FC, PropsWithChildren, useContext, useEffect, useRef } from "react";
import Context from "./Context";
import axios from "axios";
import { AuthApi, BookApi, UserApi } from "../../../services/api";
import { LoginContext } from "../../../contexts/login";

type Props = PropsWithChildren;

const Provider: FC<Props> = ({ children }) => {
  const client = useRef(axios.create());
  const { logout } = useContext(LoginContext);

  const authApi = useRef(new AuthApi(undefined, undefined, client.current));
  const userApi = useRef(new UserApi(undefined, undefined, client.current));
  const bookApi = useRef(new BookApi(undefined, undefined, client.current));

  useEffect(() => {
    client.current.interceptors.response.use(
      async (response) => response,
      (error) => {
        if (error.response && error.response.status === 401) {
          logout();
        }
        throw error;
      }
    );
  }, [logout]);

  return (
    <Context.Provider value={{ authApi: authApi.current, userApi: userApi.current, bookApi: bookApi.current }}>
      {children}
    </Context.Provider>
  );
};

export default Provider;
