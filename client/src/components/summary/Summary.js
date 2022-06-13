import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./summary.css";

export default function Summary({ isVisible, handleUpdate, id }) {
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [price, setPrice] = useState("");
  const [description, setDescription] = useState("");

  const submitHandler = (e) => {
    e.preventDefault();
    fetch(`http://localhost:8000/updateproduct/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: name,
        price: price,
        description: description,
      }),
    });
    navigate("/products");
  };

  if (isVisible) {
    return (
      <div className="summary__conatiner">
        <div className="summary__leftdiv">
          <header className="summary__header">
            <div className="summary__header__innnerdiv">
              <h1>Update</h1>
              <i onClick={handleUpdate} class="fa-solid fa-xmark fa-2xl"></i>
            </div>
          </header>
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
      </div>
    );
  } else {
    return null;
  }
}
