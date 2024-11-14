import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import { BrowserRouter } from "react-router-dom";
import { ApiContextProvider, ToastContextProvider } from "./contexts/index.tsx";
import { AuthContextProvider } from "./contexts/AuthContext/AuthContext.tsx";

createRoot(document.getElementById("root")!).render(
  <BrowserRouter
    future={{
      v7_relativeSplatPath: true,
      v7_startTransition: true,
    }}
  >
    <ApiContextProvider>
      <AuthContextProvider>
        <ToastContextProvider>
          <App />
        </ToastContextProvider>
      </AuthContextProvider>
    </ApiContextProvider>
  </BrowserRouter>
);
