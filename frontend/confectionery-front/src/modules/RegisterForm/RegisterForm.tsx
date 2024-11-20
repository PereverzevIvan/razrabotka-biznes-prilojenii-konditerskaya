import { useForm } from "react-hook-form";
import { Button } from "../../components/Button";
import "./RegisterForm.scss";
import { useState } from "react";
import { TRegisterForm } from ".";

export function RegisterForm() {
  const { register, handleSubmit, reset, formState, watch } =
    useForm<TRegisterForm>({
      mode: "onChange",
    });

  const loginWatch = watch("login");
  const { errors } = formState;

  const [passwordVisible, setPasswordVisible] = useState(false);

  function onSubmit(data: TRegisterForm) {
    console.log(data);
  }

  return (
    <>
      <form className="form register-form" onSubmit={handleSubmit(onSubmit)}>
        <fieldset className="form__fieldset">
          <legend className="form__legend">Регистрация</legend>
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
          {errors.login && (
            <div className="form__field-error">{errors.login.message}</div>
          )}

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
            <Button
              onClick={() => setPasswordVisible(!passwordVisible)}
              type="button"
            >
              Показать
            </Button>
          </div>
          {errors.password && (
            <div className="form__field-error">{errors.password.message}</div>
          )}

          <label className="form__label" htmlFor="firstname-input">
            Имя
          </label>
          <input
            id="firstname-input"
            type="text"
            className="form__input"
            placeholder="Введите имя"
            {...register("first_name")}
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
            {...register("last_name")}
          />
          {errors.last_name && (
            <div className="form__field-error">{errors.last_name.message}</div>
          )}

          <label className="form__label" htmlFor="patronymic-input">
            Отчество
          </label>
          <input
            id="patronymic-input"
            type="text"
            className="form__input"
            placeholder="Введите отчество"
            {...register("patronymic", { required: true })}
          />
          {errors.patronymic && (
            <div className="form__field-error">{errors.patronymic.message}</div>
          )}

          <div className="form__button-box">
            <Button color="green" type="submit">
              Отправить
            </Button>
            <Button onClick={() => reset()} color="red" type="button">
              Сбросить
            </Button>
          </div>
        </fieldset>
      </form>
    </>
  );
}

function validateHasLogin(value: string, loginWatch: string) {
  const isLogin = !!loginWatch;
  const hasLogin = value.includes(loginWatch);
  if (isLogin && hasLogin) {
    return "Пароль не должен содержать логин";
  }
}

// Функция для проверки наличия хотя бы одной заглавной и одной строчной буквы
const validateCase = (value: string) => {
  const hasUpperCase = /[A-Z]/.test(value);
  const hasLowerCase = /[a-z]/.test(value);
  return (
    (hasUpperCase && hasLowerCase) ||
    "Пароль должен содержать как минимум одну заглавную и одну строчную букву"
  );
};
