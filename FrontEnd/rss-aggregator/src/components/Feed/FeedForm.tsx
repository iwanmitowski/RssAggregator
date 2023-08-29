import { useState } from "react";
import { Button, Form } from "react-bootstrap";
import { Feed } from "./interfaces";
import { CatchError } from "../Shared/interfaces";
import { useNavigate } from "react-router";
  
import * as feedService from "../../services/feedService";

const FeedForm: React.FC = () => {
  const navigate = useNavigate();

  const [feed, setFeed] = useState<Feed>({
    name: "",
    url: ""
  });

  const [error, setError] =  useState<string>("");

  const onInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFeed((prevState: Feed) => {
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
    
    feedService
      .create(feed)
      .then(() => {
        navigate(`/posts`);
      })
      .catch((error: CatchError) => {
        setError(error.message);
      });
  };

return <div>
  <h1>Create feed</h1>
  <Form onSubmit={onFormSubmit}>
    <Form.Group className="form-group mb-3" controlId="name">
      <Form.Label>Name</Form.Label>
      <Form.Control
          type="name"
          name="name"
          value={feed.name}
          placeholder="Enter feed name"
          onChange={onInputChange}
          required
      />
    </Form.Group>
    <Form.Group className="form-group mb-3" controlId="url">
      <Form.Label>Url</Form.Label>
      <Form.Control
          type="url"
          name="url"
          value={feed.url}
          placeholder="Enter feed url"
          onChange={onInputChange}
          required
      />
    </Form.Group>
    {error && (
      <div className="text-danger mb-3">
        {error.split("\n").map((message, key) => {
          return <div key={key}>{message}</div>;
        })}
      </div>
    )}
    <Button
      variant="dark"
      type="submit"
    >
      {"Create feed"}
    </Button>
  </Form>
</div>;
};

export default FeedForm;