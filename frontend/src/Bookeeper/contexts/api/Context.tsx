import { createContext } from "react";
import { AuthApi, AuthApiFactory, BookApi, BookApiFactory, UserApi, UserApiFactory } from "../../../services/api";

type Props = {
  authApi: AuthApi;
  userApi: UserApi;
  bookApi: BookApi;
};

const initialState: Props = {
  authApi: AuthApiFactory() as AuthApi,
  userApi: UserApiFactory() as UserApi,
  bookApi: BookApiFactory() as BookApi,
};

const Context = createContext<Props>(initialState);

export default Context;
