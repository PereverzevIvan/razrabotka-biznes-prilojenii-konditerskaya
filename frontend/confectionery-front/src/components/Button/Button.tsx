import "./Button.scss";

type ButtonProps = {
  children: string;
  color?: "blue" | "green" | "red";
  onClick?: () => void;
  type?: "button" | "submit" | "reset";
  disabled?: boolean;
};

function Button({
  children,
  color = "blue",
  onClick,
  type,
  disabled,
}: ButtonProps) {
  return (
    <button
      type={type}
      onClick={onClick}
      className={`button button_color-${color}`}
      disabled={disabled}
    >
      {children}
    </button>
  );
}

export default Button;
