import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import { BrowserRouter } from "react-router-dom";
import { ToastContextProvider } from "./contexts/index.tsx";
import { AuthContextProvider } from "./contexts/AuthContext/AuthContext.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

createRoot(document.getElementById("root")!).render(
  <BrowserRouter
    future={{
      v7_relativeSplatPath: true,
      v7_startTransition: true,
    }}
  >
    <QueryClientProvider client={queryClient}>
      <AuthContextProvider>
        <ToastContextProvider>
          <App />
        </ToastContextProvider>
      </AuthContextProvider>
    </QueryClientProvider>
  </BrowserRouter>
);
