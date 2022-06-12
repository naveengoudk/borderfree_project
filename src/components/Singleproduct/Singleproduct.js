import React, { useContext, useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import Popup from "../Deletepopup/Popup";
import Home from "../Home/Home";
import Summary from "../summary/Summary";
import "./Singleproduct.css";
import { store } from "../../App";

export default function Singleproduct() {
  const [loggedin, setLoggedin] = useContext(store);
  const [popup, setPopup] = useState(false);
  const [product, setProduct] = useState([]);
  const [isVisible, setIsVisible] = useState(false);
  const navigate = useNavigate();
  const id = useParams();

  const showSummary = () => {
    setIsVisible(!isVisible);
  };

  const showPopup = () => {
    setPopup(!popup);
  };

  useEffect(() => {
    fetch(`https://borderfreserver.herokuapp.com/getoneproducts/${id.id}`)
      .then((resp) => resp.json())
      .then((data) => setProduct(data));
  }, []);

  if (loggedin.loggedin) {
    return (
      <>
        <Home loggedin={loggedin} setLoggedin={setLoggedin} />
        <div className="Singleproduct__container">
          <div className="Singleproduct__imgcontainer">
            <img src="/images/img2.webp"></img>
          </div>
          <div className="Singleproduct__detailscontainer">
            <div>
              <h1>{product.name}</h1>
            </div>
            <div>
              <h1>Rs. {product.price}</h1>
            </div>
            <div className="product__details">
              <h2>Product Description</h2>
              <p>{product.description}</p>
            </div>
            <div className="singleproduct__buttons">
              <button onClick={showPopup}>Delete</button>
              <button onClick={showSummary}>Update</button>
            </div>
          </div>
        </div>
        <Summary id={id.id} handleUpdate={showSummary} isVisible={isVisible} />
        <Popup id={id.id} popup={popup} setPopup={setPopup} />
      </>
    );
  } else {
    navigate("/");
  }
}
