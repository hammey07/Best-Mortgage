import { useNavigate } from "react-router-dom";

const RateCard = ({ title, description, route }) => {
  const navigate = useNavigate();

  return (
    <div className="rate-card" onClick={() => navigate(route)}>
      <h2>{title}</h2>
      <p>{description}</p>
    </div>
  );
};

export default RateCard;
