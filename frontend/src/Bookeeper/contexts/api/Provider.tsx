import { FC, PropsWithChildren, useContext, useEffect, useRef } from "react";
import Context from "./Context";
import axios from "axios";
import { AuthApi, BookApi, UserApi } from "../../../services/api";
import { LoginContext } from "../../../contexts/login";

type Props = PropsWithChildren;

const Provider: FC<Props> = ({ children }) => {
  const { token, setToken, logout } = useContext(LoginContext);

  const client = useRef(axios.create());

  const authApi = useRef(new AuthApi(undefined, undefined, client.current));
  const userApi = useRef(new UserApi(undefined, undefined, client.current));
  const bookApi = useRef(new BookApi(undefined, undefined, client.current));

  useEffect(() => {
    client.current.interceptors.request.use((request) => {
      request.headers.Authorization = `Bearer ${token.token}`;

      return request;
    });
    client.current.interceptors.response.use(
      async (response) => response,
      (error) => {
        if (error.response && error.response.status === 401) {
          logout();
        }
        if (error.response && error.response.status === 403) {
          if (token.expires && token.token && new Date(token.expires) < new Date()) {
            authApi.current
              .refreshPost()
              .then((res) => {
                if (res.data.data) {
                  setToken(res.data.data);
                }
              })
              .catch(() => {
                logout();
              });
          }
        }

        throw error;
      }
    );
  }, []);

  return (
    <Context.Provider value={{ authApi: authApi.current, userApi: userApi.current, bookApi: bookApi.current }}>
      {children}
    </Context.Provider>
  );
};

export default Provider;
