import React, { useState } from "react";
import "./Signup.css";
import { Link, useNavigate } from "react-router-dom";

export default function Signup() {
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const submitHandler = (e) => {
    e.preventDefault();
    fetch("http://localhost:8000/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: name,
        email: email,
        password: password,
      }),
    });
    alert("Signup Successful");
    navigate("/");
  };
  return (
    <div className="login__body">
      <div class="center">
        <h1>Signup</h1>
        <form onSubmit={submitHandler} method="post">
          <div class="txt_field">
            <input
              onChange={(e) => {
                setName(e.target.value);
              }}
              type="text"
              required
            />
            <span></span>
            <label>Username</label>
          </div>
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
          <input type="submit" value="Register" />
          <div class="signup_link">
            go back to login? <Link to="/">Login</Link>
          </div>
        </form>
      </div>
    </div>
  );
}
