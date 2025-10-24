import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./App.css";
import Header from "./components/Header";
import Home from "./pages/Home";
import FirstTimeBuyer from "./pages/FirstTimeBuyer";

function App() {
  return (
    <BrowserRouter>
      <Header />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/first-time-buyer" element={<FirstTimeBuyer />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
