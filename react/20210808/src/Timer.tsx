import React, {useState, useEffect} from 'react'

const LIMIT = 60;

export default function Timer() {
  const [timeLeft, setTimeLeft] = useState(LIMIT);
  const reset = () => {
    setTimeLeft(LIMIT);
  };
  const tick = () => {
    console.log('tick');
    setTimeLeft(prevTime => (prevTime === 0 ? LIMIT : prevTime -1));
  }

  useEffect(()=>{
    console.log('create Timer');
    const timerId = setInterval(tick, 1000);

    return () => {
      console.log('cleanup Timer');
      clearInterval(timerId);
    }
  }, []);

  return (
    <div>
      <p>time: {timeLeft}</p>
      <button onClick={reset}>reset</button>
    </div>
  )
}