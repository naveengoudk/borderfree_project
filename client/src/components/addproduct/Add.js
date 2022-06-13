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
  const [image, setImage] = useState("");
  const [description, setDescription] = useState("");
  const token = localStorage.getItem("token");

  const submitHandler = (e) => {
    e.preventDefault();
    fetch("http://localhost:8000/addproduct", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        name: name,
        image: image,
        price: price,
        description: description,
      }),
    })
      .then((res) => res.json())
      .then((data) => console.log(data));
    navigate("/products");
  };

  if (token) {
    return (
      <>
        <Home loggedin={loggedin} setLoggedin={setLoggedin} />
        <div className="Add__container">
          <div className="form__container">
            <form onSubmit={submitHandler}>
              <input
                onChange={(e) => {
                  setImage(e.target.value);
                }}
                name="image"
                type={"text"}
                placeholder="Image url"
              ></input>
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
