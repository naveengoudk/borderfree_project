import { createContext, useState } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./App.css";
import Add from "./components/addproduct/Add";
import Fp from "./components/Authentication/forgetpassword/Fp";
import Login from "./components/Authentication/login/Login";
import Signup from "./components/Authentication/Signup/Signup";
import Home from "./components/Home/Home";
import Products from "./components/landingpage/Products";
import Singleproduct from "./components/Singleproduct/Singleproduct";

export const store = createContext();

function App() {
  const [loggedin, setLoggedin] = useState({ loggedin: false });
  return (
    <div className="App">
      <store.Provider value={[loggedin, setLoggedin]}>
        <Router>
          <Routes>
            <Route path="/" element={<Login />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="/forgetpassword" element={<Fp />}></Route>
            <Route path="/products" element={<Products />}></Route>
            <Route path="/addproduct" element={<Add />}></Route>
            <Route
              path={`/productdetails/:id`}
              element={<Singleproduct />}
            ></Route>
          </Routes>
        </Router>
      </store.Provider>
    </div>
  );
}

export default App;
