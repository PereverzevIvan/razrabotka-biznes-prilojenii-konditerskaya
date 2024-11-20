import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { routePaths } from "../../configs";
import { useAuthContext } from "../../contexts";
import { LoginForm } from "../../modules/LoginForm";

export function LoginPage() {
  const { isAuth } = useAuthContext();
  const navigate = useNavigate();

  useEffect(() => {
    console.log(isAuth === "false", isAuth);
    if (isAuth === "true") navigate(routePaths.main.path);
  }, [isAuth]);

  return (
    <>
      <section className="page login-page">
        <h1 className="title">Login Page</h1>
        <LoginForm />
      </section>
    </>
  );
}
