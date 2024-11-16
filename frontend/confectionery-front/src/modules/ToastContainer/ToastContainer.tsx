import "./ToastContainer.scss"; // Создадим стили для тостов
import { TToast } from "../../contexts";

type ToastContainerProps = {
  toasts: TToast[];
  removeToast: (id: string) => void;
};

export const ToastContainer = ({
  toasts,
  removeToast,
}: ToastContainerProps) => {
  return (
    <div className="toast-container">
      {toasts.map((toast) => (
        <div key={toast.id} className={`toast toast-${toast.type}`}>
          <span>{toast.message}</span>
          <button className="close-btn" onClick={() => removeToast(toast.id)}>
            &times;
          </button>
        </div>
      ))}
    </div>
  );
};
