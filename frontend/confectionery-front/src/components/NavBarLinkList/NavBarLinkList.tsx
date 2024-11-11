import "./NavBarLinkList.scss";

type NavBarLinkListProps = {
  links: { to: string; label: string }[];
};

function NavBarLinkList({ links }: NavBarLinkListProps) {
  return (
    <ul className="nav-link-list h-flex">
      {links.map((link, key) => (
        <li key={key} className="nav-link-list__item">
          <a className="nav-link-list__link" href={link.to}>
            {link.label}
          </a>
        </li>
      ))}
    </ul>
  );
}

export default NavBarLinkList;
