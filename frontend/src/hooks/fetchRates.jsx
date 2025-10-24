import { useEffect, useState } from "react";

export default function fetchRates(type) {
  const [data, setData] = useState([]);

  useEffect(() => {
    fetch(`http://localhost:8001/rates/${type}`)
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
