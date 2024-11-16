import "./Navbar.scss";
import { useLocation } from "react-router-dom";
import { links } from "./routes";

/** Компонент для отображения панели управления */
export function Navbar() {
  const role = "Директор";
  const location = useLocation();

  return (
    <>
      <nav className="nav container">
        <ul className="h-flex nav__list">
          {links.map((link, key) => {
            if (link.allowedRoles.length && !link.allowedRoles.includes(role)) {
              return null;
            }

            const isActive = link.path === location.pathname;
            const linkClass = isActive ? "nav__link active" : "nav__link";

            return (
              <li key={key} className="nav__item">
                <a className={linkClass} href={link.path}>
                  {link.title}
                </a>
              </li>
            );
          })}
        </ul>
      </nav>
    </>
  );
}
