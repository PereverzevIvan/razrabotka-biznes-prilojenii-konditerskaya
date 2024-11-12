import { useState } from "react";
import logo from "../../assets/images/logo/Logo_Monochrome.png";
import Button from "../../components/Button/Button";
import Navbar from "../Navbar/Navbar";
import "./Header.scss";

function Header() {
  const [isAuth, setIsAuth] = useState(false);
  const authButton = isAuth ? (
    <Button color="red" onClick={() => setIsAuth(false)}>
      Выйти
    </Button>
  ) : (
    <Button onClick={() => setIsAuth(true)} color="green">
      Войти
    </Button>
  );

  return (
    <>
      <header className="header">
        <div className="container header__container h-flex ">
          <a href="/">
            <img src={logo} alt="AmonicLogo" className="logo" />
          </a>
          <div className="h-flex link-box">
            <Navbar />
            {authButton}
          </div>
        </div>
      </header>
    </>
  );
}

export default Header;
