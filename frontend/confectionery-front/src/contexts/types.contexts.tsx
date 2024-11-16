import { AxiosInstance } from "axios";

export type TApiContext = {
  api: AxiosInstance;
};

export type TToast = {
  id: string;
  message: string;
  type: string;
};

export type TToastContext = {
  addToast: (message: string, type?: string, seconds?: number) => void;
  removeToast: (id: string) => void;
};

export type TCredentialsForLogin = {
  login: string;
  password: string;
};

export type TAuthContext = {
  userID: string | null;
  role: string | null;
  isAuth: string | null;
  login: (credentials: TCredentialsForLogin) => Promise<any>;
  logout: () => Promise<any>;
  refresh: () => Promise<any>;
};
