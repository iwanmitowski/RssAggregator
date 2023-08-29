import React, { Fragment, useEffect, useState } from "react";
import UserForm from "../../User/UserForm";
import { Link, useNavigate } from "react-router-dom";
import { User } from "../../User/interfaces";
import { CatchError } from "../../Shared/interfaces";
    
import * as userService from "../../../services/userService";
import { useUserContext } from "../../../hooks/useUserContext";
import { UserContextType } from "../../../contexts/UserContext";

const Register: React.FC = () => {
  const { userLogin } = useUserContext() as UserContextType
  const navigate = useNavigate();

  const [user, setUser] = useState<User>({
    name: ""
  });

  const [, setError] =  useState<string>("");

  const onInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUser((prevState: User) => {
      let currentName = e.target.name;
      let currentValue = e.target.value;

      return {
        ...prevState,
        [currentName]: currentValue,
      };
    });

    setError("");
  };

  const onFormSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    userService
      .register(user)
      .then((res) => {
        userLogin(res)
        navigate(`/`);
      })
      .catch((error: CatchError) => {
        setError(error.message);
      });
  };

  return (
      <Fragment>
        <UserForm
          user={user}
          onFormSubmit={onFormSubmit}
          onInputChange={onInputChange}
        />
        <div className="mt-3">
          <p>
            Already have account ?
            <Link className="nav-link" to="/login">
              Login
            </Link>
          </p>
        </div>
      </Fragment>
    );
};
  
  export default Register;