import RateCard from "../components/RateCard";

const Home = () => {
  return (
    <div className="main-container">
      <div className="cards-container">
        <RateCard
          title="First Time Buyer"
          description="Designed for people buying a home for the first time."
          route="/first-time-buyer"
        />
        <RateCard
          title="Remortgage"
          description="For homeowners looking to switch or renew their existing mortgage."
          route="/remortgage"
        />
      </div>
    </div>
  );
};

export default Home;
