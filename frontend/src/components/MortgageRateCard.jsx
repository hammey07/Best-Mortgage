import "./MortgageRateCard.scss";

function MortgageRateCard({ item }) {
  return (
    <div className="mortgage-rate-card">
      <div className="card-header">
        <h2 className="lender-name">{item.lender}</h2>
      </div>
      <div className="card-content">
        <div className="rate-info">
          <div className="rate-primary">
            <span className="rate-value">{item.rate}</span>
            <span className="rate-label">Rate</span>
          </div>
          <div className="aprc-info">
            <span className="aprc-value">{item.APRC}</span>
            <span className="aprc-label">APRC</span>
          </div>
        </div>

        <div className="product-details">
          <div className="detail-item">
            <span className="detail-label">Product:</span>
            <span className="detail-value">{item.product}</span>
          </div>
          <div className="detail-item">
            <span className="detail-label">Max LTV:</span>
            <span className="detail-value">{item.max_LTV}</span>
          </div>
          <div className="detail-item">
            <span className="detail-label">Monthly Payment:</span>
            <span className="detail-value monthly-payment">
              {item.monthly_payment}
            </span>
          </div>
        </div>
      </div>
    </div>
  );
}

export default MortgageRateCard;
