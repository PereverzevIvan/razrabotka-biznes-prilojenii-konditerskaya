import { useForm } from "react-hook-form";
import { Button } from "../../components/Button";
import "./RegisterForm.scss";
import { useState } from "react";
import { TRegisterForm } from ".";
import { useMutation } from "@tanstack/react-query";
import { postRegister } from "./api/registerApi";
import { useToastContext } from "../../contexts";

export function RegisterForm() {
  const { addToast } = useToastContext();

  const { register, handleSubmit, reset, formState, watch } =
    useForm<TRegisterForm>({
      mode: "onChange",
    });

  const {
    mutate,
    isPending,
    isError,
    error: registerError,
  } = useMutation({
    mutationFn: (data: TRegisterForm) => postRegister(data),
    onSuccess: () => {
      reset();
      addToast("Регистрация прошла успешно");
    },
    onError: (error: any) => {
      console.error(error);
      if (error?.response?.data) {
        addToast(error?.response?.data, "error");
      } else {
        addToast("Произошла ошибка при регистрации", "error");
      }
    },
  });

  const loginWatch = watch("login");
  const { errors } = formState;

  const [passwordVisible, setPasswordVisible] = useState(false);

  function onSubmit(data: TRegisterForm) {
    data = prepareData(data);
    mutate(data);
  }

  return (
    <>
      <form className="form register-form" onSubmit={handleSubmit(onSubmit)}>
        {/* Отображение ошибок */}
        {isError &&
          (typeof registerError?.response?.data === "string" &&
          registerError?.response?.data != "" ? (
            <div className="message message_error">
              {registerError?.response?.data}
            </div>
          ) : (
            <div className="message message_error">{registerError.message}</div>
          ))}
        <fieldset className="form__fieldset">
          <legend className="form__legend">Регистрация</legend>

          {/* Поле для ввода логина */}
          <label className="form__label" htmlFor="login-input">
            Логин
          </label>
          <input
            id="login-input"
            type="text"
            className="form__input"
            placeholder="Введите логин"
            {...register("login", {
              required: "Логин обязателен",
              minLength: {
                value: 5,
                message: "Логин должен содержать не менее 5 символов",
              },
            })}
          />

          {/* Отображение ошибок для логина */}
          {errors.login && (
            <div className="form__field-error">{errors.login.message}</div>
          )}

          {/* Поле для ввода пароля */}
          <label className="form__label" htmlFor="password-input">
            Пароль
          </label>
          <div className="h-flex password-input">
            <input
              id="password-input"
              type={passwordVisible ? "text" : "password"}
              className="form__input"
              placeholder="Введите пароль"
              {...register("password", {
                required: "Пароль обязателен",
                minLength: {
                  value: 5,
                  message: "Пароль должен содержать не менее 5 символов",
                },
                maxLength: {
                  value: 20,
                  message: "Пароль должен содержать не более 20 символов",
                },
                validate: {
                  containsLogin: (value: string) =>
                    validateHasLogin(value, loginWatch),
                  validateCase: validateCase,
                },
              })}
            />
            {/* Кнопка для отображения пароля */}
            <Button
              onClick={() => setPasswordVisible(!passwordVisible)}
              type="button"
            >
              {passwordVisible ? "Скрыть" : "Показать"}
            </Button>
          </div>
          {errors.password && (
            <div className="form__field-error">{errors.password.message}</div>
          )}

          {/* Поля для ввода имени, фамилии и отчества */}
          <label className="form__label" htmlFor="firstname-input">
            Имя
          </label>
          <input
            id="firstname-input"
            type="text"
            className="form__input"
            placeholder="Введите имя"
            {...register("first_name", { required: "Имя обязательно" })}
          />
          {errors.first_name && (
            <div className="form__field-error">{errors.first_name.message}</div>
          )}

          <label className="form__label" htmlFor="lastname-input">
            Фамилия
          </label>
          <input
            id="lastname-input"
            type="text"
            className="form__input"
            placeholder="Введите фамилию"
            {...register("last_name", { required: "Фамилия обязательна" })}
          />
          {errors.last_name && (
            <div className="form__field-error">{errors.last_name.message}</div>
          )}

          <label className="form__label" htmlFor="patronymic-input">
            Отчество (если есть)
          </label>
          <input
            id="patronymic-input"
            type="text"
            className="form__input"
            placeholder="Введите отчество (если есть)"
            {...register("patronymic")}
          />
          {errors.patronymic && (
            <div className="form__field-error">{errors.patronymic.message}</div>
          )}

          {/* Кнопки отправки и сброса формы */}
          <div className="form__button-box">
            <Button
              disabled={!formState.isValid || isPending}
              color="green"
              type="submit"
            >
              Отправить
            </Button>
            <Button
              disabled={!formState.isValid || isPending}
              onClick={() => reset()}
              color="red"
              type="button"
            >
              Сбросить
            </Button>
          </div>
        </fieldset>
      </form>
    </>
  );
}

/** Функция для подготовки данных перед отправкой. В пароле лишние пробелы на удаляются, так как они могут быть его частью. */
function prepareData(data: TRegisterForm) {
  return {
    ...data,
    login: data.login.trim(),
    first_name: data.first_name.trim(),
    last_name: data.last_name.trim(),
    patronymic: data.patronymic?.trim(),
  };
}

/** Функция для проверки наличия логина в пароле */
function validateHasLogin(value: string, loginWatch: string) {
  const isLogin = !!loginWatch;
  const hasLogin = value.includes(loginWatch);
  if (isLogin && hasLogin) {
    return "Пароль не должен содержать логин";
  }
}

/** Функция для проверки наличия хотя бы одной заглавной и одной строчной буквы */
const validateCase = (value: string) => {
  const hasUpperCase = /[A-Z]/.test(value);
  const hasLowerCase = /[a-z]/.test(value);
  return (
    (hasUpperCase && hasLowerCase) ||
    "Пароль должен содержать как минимум одну заглавную и одну строчную букву"
  );
};
