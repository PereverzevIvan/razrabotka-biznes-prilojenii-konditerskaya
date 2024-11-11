import { useState } from "react";
import logo from "../../assets/images/logo/Logo_Monochrome.png";
import Button from "../../components/Button/Button";
import Navbar from "../Navbar/Navbar";
import "./Header.scss";

function Header() {
  const [isAuth, setIsAuth] = useState(false);

  return (
    <>
      <header className="header">
        <div className="container header__container">
          <a href="/">
            <img src={logo} alt="AmonicLogo" className="logo" />
          </a>
          <Navbar />
          {isAuth ? (
            <Button color="red" onClick={() => setIsAuth(false)}>
              Выйти
            </Button>
          ) : (
            <Button onClick={() => setIsAuth(true)} color="green">
              Войти
            </Button>
          )}
        </div>
      </header>
    </>
  );
}

export default Header;
