import { useContext } from "react";
import { UserContext } from "../contexts/UserContext";

export const useUserContext = () => {
    const context = useContext(UserContext);

    return context;
};