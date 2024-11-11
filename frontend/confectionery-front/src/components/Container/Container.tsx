import { ReactNode } from "react";
import "./Container.scss";

type TProps = {
  children: ReactNode;
  addClass?: string;
};

function Container(props: TProps) {
  const classes: string = "container" + (props.addClass ? props.addClass : "");
  return <div className={classes}>{props.children}</div>;
}

export default Container;
