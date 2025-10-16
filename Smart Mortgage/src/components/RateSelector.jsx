export default function RateSelector({
  selectedRateType,
  setSelectedRateType,
  rateDescription
}) {
  return (
    <div onClick={() => setSelectedRateType(selectedRateType)}>
      <h2>{selectedRateType} </h2>
      <p>{rateDescription} </p>
    </div>
  );
}
