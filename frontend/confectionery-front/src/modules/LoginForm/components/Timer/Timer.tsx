import "./Timer.scss";
type TTimerProps = {
  seconds: number;
  isVisible?: boolean;
};

export function Timer({ seconds, isVisible = false }: TTimerProps) {
  return (
    <>
      <div
        className="timer"
        style={isVisible ? { display: "block" } : { display: "none" }}
      >
        <div
          className="line"
          style={
            isVisible
              ? { animation: `countdown ${seconds}s linear forwards` }
              : {}
          }
        ></div>
      </div>
    </>
  );
}
