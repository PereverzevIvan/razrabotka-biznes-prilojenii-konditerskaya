import { useState } from "react";
import logo from "../../assets/images/logo/Logo_Monochrome.png";
import Navbar from "../Navbar/Navbar";
import "./Header.scss";
import { LoginButton } from "./components/LoginButton";
import { LogoutButton } from "./components/LogoutButton";

function Header() {
  const [isAuth, setIsAuth] = useState(false);

  return (
    <>
      <header className="header">
        <div className="container header__container h-flex ">
          <a href="/">
            <img src={logo} alt="AmonicLogo" className="logo" />
          </a>
          <div className="h-flex link-box">
            <Navbar />
            {isAuth === false ? <LoginButton /> : <LogoutButton />}
          </div>
        </div>
      </header>
    </>
  );
}

export default Header;
