import { useEffect, useState } from "react";

export default function MortgageData() {
  const [data, setData] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8001/rates/first_time_buyer")
      .then((res) => res.json())
      .then((json) => {
        console.log(json);
        setData(json);
      })
      .catch((err) => console.error("Error fetching data:", err));
  }, []);

  return (
    <pre>{JSON.stringify(data, null, 2)}</pre> // prints nicely on screen
  );
}
