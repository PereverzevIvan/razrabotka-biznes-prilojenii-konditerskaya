import {
  createContext,
  useState,
  useContext,
  useCallback,
  ReactNode,
} from "react";
import { TToast, TToastContext } from "..";
import { ToastContainer } from "../../modules/ToastContainer";

// Создаем контекст
const ToastContext = createContext<TToastContext | null>(null);

export const useToastContext = () => useContext(ToastContext) as TToastContext;

// Провайдер для оборачивания всего приложения
export const ToastContextProvider = ({ children }: { children: ReactNode }) => {
  const [toasts, setToasts] = useState<TToast[]>([]);

  const addToast = useCallback(
    (message: string, type: string = "success", seconds: number = 5) => {
      const id = Math.random().toString(36).substr(2, 9); // Генерируем уникальный ID для каждого тоста
      setToasts((prevToasts) => [...prevToasts, { id, message, type }]);

      // Удаляем тост автоматически по истечении времени
      setTimeout(() => {
        removeToast(id);
      }, seconds * 1000);
    },
    []
  );

  const removeToast = useCallback((id: string) => {
    setToasts((prevToasts) => prevToasts.filter((toast) => toast.id !== id));
  }, []);

  return (
    <ToastContext.Provider value={{ addToast, removeToast }}>
      <ToastContainer toasts={toasts} removeToast={removeToast} />
      {children}
    </ToastContext.Provider>
  );
};
