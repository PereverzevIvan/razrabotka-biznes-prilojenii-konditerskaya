import { roles } from "../../configs";
import { useAuthContext } from "../../contexts";

export function MainPage() {
  const { role } = useAuthContext();

  return (
    <section className="page main-page">
      <h1 className="title">Main Page</h1>
      <p className="common-text">
        Добро пожаловать на сайт нашей кондитерской!
      </p>
      <p className="common-text">Ваша роль: {!role ? roles.guest : role}.</p>
      <p className="common-text">{!role && "Попробуйте войти в систему."}</p>
    </section>
  );
}
