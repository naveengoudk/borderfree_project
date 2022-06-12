import React, { useContext, useState } from "react";
import { Link } from "react-router-dom";
import { store } from "../../App";
import { Navigate } from "react-router-dom";
import "./Home.css";

export default function Home({ name, loggedin, setLoggedin }) {
  const [userdata, setuserdata] = useState({});
  // const [name, setName] = useState("");
  const logoutHandler = () => {
    setLoggedin({ loggedin: false });
  };
  return (
    <>
      <div className="home__container">
        <header className="home__header">
          <div className="home__header__insidediv">
            <div className="home__header__title">
              <h1>BorderFree</h1>
            </div>
            <div className="home__header__navigation">
              <button onClick={logoutHandler}>Logout</button>
            </div>
          </div>
          <div className="home__header__userinfo">
            <div className="home__header__userimage">
              <img src="https://lh3.googleusercontent.com/ogw/ADea4I7Jg1mUhjHMgDuy34nUCvmABKEPG3wOr4p2SzlOsg=s32-c-mo"></img>
            </div>
            <h1>{loggedin.name}</h1>
          </div>
        </header>
        <nav className="home__navbar">
          <Link className="linkdiv" to={"#"}>
            <i class="fa-solid fa-house"></i>
          </Link>
          <Link className="linkdiv" to={"/addproduct"}>
            <i class="fa-solid fa-circle-plus"></i>
          </Link>
          <Link className="linkdiv" to={"/products"}>
            <i class="fa-solid fa-list"></i>
          </Link>
        </nav>
        <footer className="home__footer">
          <p>2021 &#169; Borderfree</p>
        </footer>
      </div>
      {/* // )} */}
    </>
  );
}
