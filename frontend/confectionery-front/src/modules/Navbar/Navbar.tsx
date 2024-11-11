import NavBarLinkList from "../../components/NavBarLinkList/NavBarLinkList";
import "./Navbar.sass";

const links = [
  { to: "/", label: "Главная" },
  { to: "/about", label: "О нас" },
];

function Navbar() {
  return (
    <>
      <nav className="navbar">
        <NavBarLinkList links={links} />
      </nav>
    </>
  );
}

export default Navbar;
