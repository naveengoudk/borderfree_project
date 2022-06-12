import React, { useState } from "react";
import "./Fp.css";
import { Link, useNavigate } from "react-router-dom";

export default function Fp() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const submitHandler = (e) => {
    e.preventDefault();
    fetch("https://borderfreserver.herokuapp.com/forgetpassword", {
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
        if (data == "User not found") {
          alert("User not found");
        } else {
          alert("Password updated");
          navigate("/");
        }
      });
  };
  return (
    <div className="login__body">
      <div class="center">
        <h1>Recover Password</h1>
        <form onSubmit={submitHandler} method="post">
          <div class="txt_field">
            <input
              onChange={(e) => {
                setEmail(e.target.value);
              }}
              type="email"
              required
            />
            <span></span>
            <label>Confirm Email</label>
          </div>
          <div class="txt_field">
            <input
              onChange={(e) => {
                setPassword(e.target.value);
              }}
              type="text"
              required
            />
            <span></span>
            <label>New Password</label>
          </div>
          <input type="submit" value="Reset Password" />
          <div class="signup_link">
            go back to login? <Link to="/">Login</Link>
          </div>
        </form>
      </div>
    </div>
  );
}
