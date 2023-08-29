import { createContext } from "react";
import { useLocalStorage } from "../hooks/useLocalStorage";
import { User, UserApiResponse } from "../components/User/interfaces";
import { ChildrenProps } from "../components/Shared/interfaces";

export interface UserContextType {
    user: User
    userLogin: (data: UserApiResponse) => void | PromiseLike<void>
    userLogout: () => void
    isLogged: boolean
}

export const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<ChildrenProps> = ({ children }) => {
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
        isLogged: user && !!user.api_key,
      }}
    >
      {children}
    </UserContext.Provider>
  );
};