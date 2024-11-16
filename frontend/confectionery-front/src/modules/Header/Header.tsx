import logo from "../../assets/images/logo/Logo_Monochrome.png";
import "./Header.scss";
import { LoginButton } from "./components/LoginButton";
import { LogoutButton } from "./components/LogoutButton";
import { useAuthContext } from "../../contexts";
import { Navbar } from "../Navbar";

export function Header() {
  const { isAuth } = useAuthContext();

  return (
    <>
      <header className="header v-flex">
        <div className="container header__container h-flex">
          <a href="/" className="header__logo-container h-flex">
            <img
              src={logo}
              alt="Логотип кондитерской"
              className="header__logo"
            />
            <h1 className="header__title">Кондитерская</h1>
          </a>

          <div className="h-flex link-box">
            {!isAuth || isAuth === "false" ? <LoginButton /> : <LogoutButton />}
          </div>
        </div>
        <Navbar />
      </header>
    </>
  );
}
