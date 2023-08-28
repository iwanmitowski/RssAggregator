import { ReactNode } from "react";

interface Props {
  children: ReactNode;
}

const Main: React.FC<Props> = ({ children }) => {
  return <div>{children}</div>;
};

export default Main;