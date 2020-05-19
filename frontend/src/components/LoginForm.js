import React, { useState } from "react";
import { Col, Form, FormGroup, Label, Input, Button } from "reactstrap";
import { useHistory, useLocation } from "react-router-dom";

export const useInput = (initialValue) => {
  const [values, setValues] = useState(initialValue || {});

  return {
    values,
    setValues,
    bind: {
      values,
      onChange: (e) => {
        const { target } = e;
        const { name, value } = target;
        setValues({ ...values, [name]: value });
      },
    },
  };
};

const LoginForm = (props) => {
  let history = useHistory();
  let location = useLocation();
  let { from } = location.state || { from: { pathname: props.to } };

  const { values, bind } = useInput("");

  const handleSubmit = (evt) => {
    evt.preventDefault();
    console.log(`Submitting Name ${values}`);
    props.submit(values).then(function () {
      history.replace(from);
    });
  };

  return (
    <Form>
      <h2 sm={2}>{props.heading}</h2>
      <FormGroup row>
        <Label for="uname" sm={2}>
          Username
        </Label>
        <Col sm={10}>
          <Input
            type="text"
            name="username"
            id="uname"
            placeholder="Enter Username"
            {...bind}
          />
        </Col>
      </FormGroup>
      <FormGroup row>
        <Label for="pword" sm={2}>
          Password
        </Label>
        <Col sm={10}>
          <Input
            type="password"
            name="password"
            id="pword"
            placeholder="Enter Password"
            {...bind}
          />
        </Col>
      </FormGroup>
      <Button onClick={handleSubmit}>Submit</Button>
    </Form>
  );
};

export default LoginForm;
