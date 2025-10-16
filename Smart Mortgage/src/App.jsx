import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import Header from "./components/header";
import RateSelector from "./components/RateSelector";

function App() {
  const [selectedRateType, setSelectedRateType] = useState("");

  return (
    <>
      <Header></Header>
      <RateSelector
        selectedRateType={"First Time Buyer"}
        setSelectedRateType={setSelectedRateType}
        rateDescription={"Designed for people buying a home for the first time. Usually offers lower rates to help new buyers get on the property ladder."
} 
      />
      <RateSelector
        selectedRateType={"Remortgage"}
        setSelectedRateType={setSelectedRateType}
        rateDescription={"For homeowners looking to switch or renew their existing mortgage. Rates may differ depending on your current loan and equity."}
      />
    {selectedRateType}
    </>
  );
}

export default App;
