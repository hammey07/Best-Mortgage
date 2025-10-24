import { useEffect, useState } from "react";
import MortgageRateCard from "../components/MortgageRateCard";
import "./FirstTimeBuyer.scss";

const FirstTimeBuyer = () => {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8001/rates/first-time-buyer")
      .then((res) => res.json())
      .then(setData);
  }, []);

  if (!data) return <p className="loading">Loading...</p>;

  return (
    <div className="first-time-buyer">
      <div className="page-header">
        <h1>First Time Buyer Rates</h1>
        <p className="page-subtitle">Compare mortgage rates from top lenders</p>
      </div>
      <div className="rates-grid">
        {data.map((item, index) => (
          <MortgageRateCard key={index} item={item} />
        ))}
      </div>
    </div>
  );
};

export default FirstTimeBuyer;
