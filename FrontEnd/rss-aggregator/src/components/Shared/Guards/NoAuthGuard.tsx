import { Navigate, Outlet } from "react-router-dom";
import { useUserContext } from "../../../hooks/useUserContext";
import { ChildrenProps } from "../interfaces";
import { UserContextType } from "../../../contexts/UserContext";

const NoAuthGuard: React.FC<ChildrenProps> = ({ children }) => {
  const { isLogged } = useUserContext() as UserContextType;

  if (isLogged) {
    return <Navigate to="/" replace />;
  }

  return children ? children : <Outlet />;
}

export default NoAuthGuard;