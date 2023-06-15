import { signal } from "@preact/signals";
import "./app.css";

export function App() {
  const count = signal(0);

  const onIncrementClick = () => {
    count.value++;
  };

  return (
    <>
      <div>count is : {count}</div>
      <button onClick={onIncrementClick}>Increment</button>
    </>
  );
}
