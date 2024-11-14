import { useForm } from "react-hook-form";
import Button from "../../../../components/Button/Button";

// Определяем поля для формы
type TLoginFormData = {
  login: string;
  password: string;
};

function LoginForm() {
  // Создаем форму
  const { register, handleSubmit, formState, reset, setValue } =
    useForm<TLoginFormData>({
      mode: "onChange",
      defaultValues: {
        login: "login",
        password: "pass",
      },
    });

  const loginError = formState.errors.login;
  const passwordError = formState.errors.password;

  // Обработчик сабмита
  const onSubmit = (data: any) => {
    console.log(data);
    console.log("Писюн");
  };

  return (
    <form className="form v-flex" onSubmit={handleSubmit(onSubmit)}>
      <fieldset className="form__fieldset v-flex">
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

        {/* Выводим ошибку, если она есть */}
        {loginError && (
          <span className="form__field-error">{loginError.message}</span>
        )}
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
        <Button color="green" type="submit">
          Войти
        </Button>
        <Button color="red" type="button" onClick={() => reset()}>
          Сбросить
        </Button>
      </fieldset>
    </form>
  );
}

export default LoginForm;
