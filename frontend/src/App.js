import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect,
} from "react-router-dom";
import "./App.css";
import { LoginForm } from "./components";
import { Container, Row, Col } from "reactstrap";

function App() {
  return (
    <Container>
      <p>Basic exmaple of react using OAuth JWTs</p>
      <Row>
        <Col>
          <Router>
            <nav>
              <ul>
                <li>
                  <Link to="/">Protected</Link>
                </li>
                <li>
                  <Link to="/login">Login</Link>
                </li>
                <li>
                  <Link to="/signup">Sign up</Link>
                </li>
              </ul>
            </nav>
            <Switch>
              <Route path="/login">
                <LoginForm
                  heading="Login"
                  to="/"
                  submit={authService.authenticate}
                ></LoginForm>
              </Route>
              <Route path="/signup">
                <LoginForm
                  heading="Sign Up"
                  to="/login"
                  submit={userService.create}
                ></LoginForm>
              </Route>
              <PrivateRoute path="/">
                <ProtectedPage></ProtectedPage>
              </PrivateRoute>
            </Switch>
          </Router>
        </Col>
      </Row>
    </Container>
  );
}

function ProtectedPage() {
  return <h3>Protected</h3>;
}

// A wrapper for <Route> that redirects to the login
// screen if you're not yet authenticated.
function PrivateRoute({ children, ...rest }) {
  return (
    <Route
      {...rest}
      render={({ location }) =>
        authService.isAuthenticated ? (
          children
        ) : (
          <Redirect
            to={{
              pathname: "/login",
              state: { from: location },
            }}
          />
        )
      }
    />
  );
}

const userService = {
  async create(creds) {
    try {
      const response = await fetch("http://localhost:8080/user", {
        method: "post",
        body: JSON.stringify(creds),
      });
      if (response.status !== 200) {
        throw new Error(response.status);
      }
      return await response.json();
    } catch (error) {
      console.log(error);
      throw new Error(error);
    }
  },
};

const authService = {
  isAuthenticated: false,
  token: null,
  authenticate(creds) {
    authService.isAuthenticated = false;
    console.log(creds);
    return fetch("http://localhost:8080/auth/login", {
      method: "post",
      body: JSON.stringify(creds),
    })
      .then(function (response) {
        if (response.status !== 200) {
          return;
        }

        return response.json().then(function (data) {
          console.log(data);
          authService.isAuthenticated = true;
          authService.token = data.token;
        });
      })
      .catch(function (error) {
        console.log(error);
      });
  },
  validate() {
    fetch("http://localhost:8080/auth/validate", {
      method: "post",
      body: { token: authService.token },
    })
      .then(function (response) {
        if (response.status !== 201) {
          authService.isAuthenticated = false;
          return;
        }

        response.json().then(function (data) {
          console.log(data);
          authService.isAuthenticated = true;
          authService.token = data.token;
        });
      })
      .catch(function (error) {
        console.log(error);
      });
  },
  refresh() {
    fetch("http://localhost:8080/auth/refresh", {
      method: "post",
      body: { token: authService.token },
    })
      .then(function (response) {
        if (response.status !== 201) {
          return;
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  },
  signout() {
    authService.isAuthenticated = false;
    authService.token = null;
  },
};

export default App;
