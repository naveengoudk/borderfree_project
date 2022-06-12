import React, { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import Home from "../Home/Home";
import "./Add.css";
import { store } from "../../App";

export default function Add() {
  const [loggedin, setLoggedin] = useContext(store);
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [price, setPrice] = useState("");
  const [description, setDescription] = useState("");

  const submitHandler = (e) => {
    e.preventDefault();
    console.log(name, price, description);
    fetch("https://borderfreserver.herokuapp.com/addproduct", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        user: loggedin.email,
        name: name,
        price: price,
        description: description,
      }),
    });
    navigate("/products");
  };

  if (loggedin.loggedin) {
    return (
      <>
        <Home loggedin={loggedin} setLoggedin={setLoggedin} />
        <div className="Add__container">
          <div className="form__container">
            <form onSubmit={submitHandler}>
              <label for="name">Name</label>
              <input
                onChange={(e) => {
                  setName(e.target.value);
                }}
                name="name"
                type="text"
                required
              ></input>
              <label for="price">Price</label>
              <input
                onChange={(e) => {
                  setPrice(e.target.value);
                }}
                name="price"
                type="text"
                // min={"0"}
                required
              ></input>
              <label for="description">Description</label>
              <textarea
                onChange={(e) => {
                  setDescription(e.target.value);
                }}
                rows="8"
                cols={"10"}
                name="description"
                type="text"
                required
              ></textarea>
              <button>Submit</button>
            </form>
          </div>
        </div>
      </>
    );
  } else {
    navigate("/");
  }
}
