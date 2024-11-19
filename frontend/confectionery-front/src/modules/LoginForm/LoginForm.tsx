import { useEffect, useState } from "react";
import {
  TCredentialsForLogin,
  useAuthContext,
  useToastContext,
} from "../../contexts";
import { Timer } from "./components/Timer/Timer";
import { Button } from "../../components/Button";
import { useForm } from "react-hook-form";
import "./LoginForm.scss";

const MAX_ATTEMPTS = 3;
const TIMER_SECONDS = 5;

export function LoginForm() {
  const { login } = useAuthContext();
  const { addToast } = useToastContext();

  const [message, setMessage] = useState("");
  const [attempts, setAttempts] = useState(0);

  // Создаем форму
  const { register, handleSubmit, formState, reset } =
    useForm<TCredentialsForLogin>({
      mode: "onChange",
      defaultValues: {
        login: "",
        password: "",
      },
    });

  // Получаем ошибки
  const loginError = formState.errors.login;
  const passwordError = formState.errors.password;

  // Обработчик сабмита
  const onSubmit = (data: TCredentialsForLogin) => {
    login(data)
      .then(() => {
        reset();
        addToast("Вы успешно вошли в систему", "success");
      })
      .catch((error) => {
        console.error(error);
        setAttempts(attempts + 1);
      });
  };

  // Функция для старта таймера и сброса попыток входа
  function resetFailedTries() {
    setMessage(
      `Достигнуто максимальное количество попыток входа. Повторите попытку через ${TIMER_SECONDS} секунд.`
    );

    setTimeout(() => {
      setAttempts(0);
      setMessage("");
      console.log("Вы можете снова попытаться войти в систему");
    }, TIMER_SECONDS * 1000);
  }

  useEffect(() => {
    if (attempts == MAX_ATTEMPTS) resetFailedTries();
  }, [attempts]);

  return (
    <form className="form login-form" onSubmit={handleSubmit(onSubmit)}>
      {/* Сообщение для вывода ошибок от сервера */}
      {message && <div className="message message_error">{message}</div>}

      {/* Таймер для отчета времени до следующей попытки входа */}
      <Timer seconds={5} isVisible={attempts === MAX_ATTEMPTS} />

      {/* Основная секция */}
      <fieldset className="form__fieldset">
        <legend className="form__legend">Login form</legend>

        {/* Поле для ввода логина */}
        <label className="form__label" htmlFor="login">
          Login:
        </label>
        <input
          className="form__input"
          type="text"
          id="login"
          placeholder="Enter the Login"
          /* Добавляем поле в форму */
          {...register("login", {
            required: "Login is required", // Добавляем сообщение об ошибке и условие обязательности поля
          })}
        />
        {loginError && ( // Выводим ошибку, если она есть
          <span className="form__field-error">{loginError.message}</span>
        )}

        {/* Поле для ввода пароля */}
        <label className="form__label" htmlFor="password">
          Password:
        </label>
        <input
          className="form__input"
          type="password"
          id="password"
          placeholder="Enter the Password"
          {...register("password", { required: "Password is required" })}
        />
        {passwordError && (
          <span className="form__field-error">{passwordError.message}</span>
        )}

        {/* Кнопки для отправки и сброса формы */}
        <div className="form__button-box">
          <Button
            disabled={!formState.isValid || attempts === MAX_ATTEMPTS}
            color="green"
            type="submit"
          >
            Войти
          </Button>
          <Button
            disabled={attempts === MAX_ATTEMPTS}
            color="red"
            type="button"
            onClick={() => reset()}
          >
            Сбросить
          </Button>
        </div>
      </fieldset>
    </form>
  );
}
