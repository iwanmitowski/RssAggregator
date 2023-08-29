import { ReactNode, createContext } from "react";
import { useLocalStorage } from "../hooks/useLocalStorage";
import { User, UserApiResponse } from "../components/User/interfaces";

export interface UserContextType {
    user: User
    userLogin: (data: UserApiResponse) => void | PromiseLike<void>;
    userLogout: () => void;
}

interface Props {
    children: ReactNode
}

export const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<Props> = ({ children }) => {
  const [user, setUser] = useLocalStorage("auth", null);

  const userLogin = (data: object) => {
    setUser(data); 
  };

  const userLogout = () => {
    setUser(null);
  };

  return (
    <UserContext.Provider
      value={{
        user,
        userLogin,
        userLogout,
      }}
    >
      {children}
    </UserContext.Provider>
  );
};