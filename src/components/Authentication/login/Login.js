import React, { useContext, useState } from "react";
import "./Login.css";
import { Link, useNavigate } from "react-router-dom";
import { store } from "../../../App";

export default function Login() {
  const [loggedin, setLoggedin] = useContext(store);
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const submitHandler = (e) => {
    e.preventDefault();

    fetch("https://borderfreserver.herokuapp.com/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email,
        password: password,
      }),
    })
      .then((resp) => resp.json())
      .then((data) => {
        if (data === "User not found" || data === "Wrong password") {
          alert(data);
        } else {
          alert("Login Successful");
          setLoggedin({ loggedin: true, ...data });
          navigate("/products");
        }
      });
  };
  // console.log(loggedin);
  return (
    <div className="login__body">
      <div class="center">
        <h1>Login</h1>
        <form onSubmit={submitHandler} method="post">
          <div class="txt_field">
            <input
              onChange={(e) => {
                setEmail(e.target.value);
              }}
              type="text"
              required
            />
            <span></span>
            <label>Email</label>
          </div>
          <div class="txt_field">
            <input
              onChange={(e) => {
                setPassword(e.target.value);
              }}
              type="password"
              required
            />
            <span></span>
            <label>Password</label>
          </div>
          <div class="pass">
            <Link to={"/forgetpassword"}>Forgot Password?</Link>
          </div>
          <input type="submit" value="Login" />
          <div class="signup_link">
            Not a member? <Link to="/signup">Signup</Link>
          </div>
        </form>
      </div>
    </div>
  );
}
