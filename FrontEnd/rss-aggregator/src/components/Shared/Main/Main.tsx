import { ChildrenProps } from "../interfaces";

const Main: React.FC<ChildrenProps> = ({ children }) => {
  return <div>{children}</div>;
};

export default Main;