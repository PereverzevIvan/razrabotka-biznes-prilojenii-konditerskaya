import axios from "axios";
import { createContext, ReactNode, useContext } from "react";
import { TApiContext } from "..";

const ApiContext = createContext<TApiContext | null>(null);

export function ApiContextProvider({ children }: { children: ReactNode }) {
  const api = axios.create({
    baseURL: "http://localhost:3000/api",
    withCredentials: true,
  });

  return <ApiContext.Provider value={{ api }}>{children}</ApiContext.Provider>;
}

export function useApiContext() {
  return useContext(ApiContext);
}
