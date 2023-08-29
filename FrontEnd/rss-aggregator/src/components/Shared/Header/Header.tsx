import { Link } from "react-router-dom";
import { Fragment } from "react";

import { useUserContext } from "../../../hooks/useUserContext";
import { ButtonGroup, Dropdown, DropdownButton } from "react-bootstrap";
import { UserContextType } from "../../../contexts/UserContext";

export default function Header() {
  const { isLogged } = useUserContext() as UserContextType;

  return (
    <div
      className="d-flex flex-column flex-shrink-0 p-3 bg-light fixed-top"
      style={{ width: "280px", minHeight: "100vh", textAlign: "left" }}
    >
      <Link className="nav-link" to="/">
        RssAggregator
      </Link>
      <hr />
      <ul className="nav nav-pills flex-column mb-auto">
        {isLogged && (
          <Fragment>
            <li className="nav-item">
              <Link className="nav-link text-dark" to="/posts">
                Posts
              </Link>
            </li>
            <li className="nav-item">
              <Link className="nav-link text-dark" to="/feed">
                Create feed
              </Link>
            </li>
          </Fragment>
        )}
      </ul>
      <hr />
      {isLogged && (
        <DropdownButton
          as={ButtonGroup}
          key="danger"
          drop="up"
          id={`actions-dropdown`}
          variant="light"
          title={
            isLogged && (
              <div className="d-inline-flex text-wrap">
                <strong>Actions</strong>
              </div>
            )
          }
        >
          <Dropdown.Item eventKey="2">
            <Link className="nav-link text-dark" to="/logout">
              Logout
            </Link>
          </Dropdown.Item>
        </DropdownButton>
      )}
    </div>
  );
}