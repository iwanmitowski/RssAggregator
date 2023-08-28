import { Button, Form } from "react-bootstrap";
import { User } from "./interfaces";

interface Props {
  onFormSubmit: (e: React.FormEvent) => void
  onInputChange: (e: React.ChangeEvent<HTMLInputElement>) => void
  user: User,
}
  
const UserForm: React.FC<Props> = (props) => {
const { onFormSubmit, onInputChange, user } = props

return <div>
  <Form onSubmit={onFormSubmit}>
    <Form.Group className="form-group mb-3" controlId="name">
    <Form.Label>Name</Form.Label>
    <Form.Control
        type="name"
        name="name"
        value={user.name}
        placeholder="Enter address"
        onChange={onInputChange}
        required
    />
    </Form.Group>
    {/* {error && (
      <div className="text-danger mb-3">
        {error.split("\n").map((message, key) => {
          return <div key={key}>{message}</div>;
        })}
      </div>
    )} */}
    <Button
      variant="dark"
      type="submit"
    >
      {"Register"}
    </Button>
  </Form>
</div>;
};

export default UserForm;