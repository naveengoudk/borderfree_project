import React, { useContext, useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import Home from "../Home/Home";
import "./products.css";
import { store } from "../../App";

export default function Products() {
  const [loggedin, setLoggedin] = useContext(store);
  const navigate = useNavigate();
  const userEmail = loggedin.email;
  const [products, setProducts] = useState([]);
  useEffect(() => {
    fetch(`https://borderfreserver.herokuapp.com/getproducts/${userEmail}`)
      .then((resp) => resp.json())
      .then((data) => setProducts(data));
  }, []);

  if (loggedin.loggedin) {
    return (
      <>
        <Home loggedin={loggedin} setLoggedin={setLoggedin} />
        <div className="landingpage__container">
          {products == null ? (
            <h1>No products available</h1>
          ) : (
            products.map((product) => {
              return (
                <Link
                  className="link__container"
                  to={`/productdetails/${product._id}`}
                >
                  <div className="container">
                    <div className="image__container">
                      <img src="images/img2.webp"></img>
                    </div>
                    <div className="name__container">
                      <h1>{product.name}</h1>
                    </div>
                    <div className="description__container">
                      <p>Rs. {product.price}</p>
                    </div>
                  </div>
                </Link>
              );
            })
          )}
        </div>
      </>
    );
  } else {
    navigate("/");
  }
}
