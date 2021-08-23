import React, {useState, useMemo} from 'react'
import Timer from './Timer'

function App() {

  const [visible, setVisible] = useState(true);
  const [count1, setCount1] = useState(0);
  const [count2, setCount2] = useState(0);

  const dabble = (count: number): number => {
    let i = 0;
    while(i < 1000000000) i++;
    return count * 2;
  };

  const dabbleCount = useMemo(()=> dabble(count1), [count1] );

  return (
    <div>
      <button onClick ={() => setVisible(!visible)}>toggle</button>
      {visible ? <Timer /> : ''}

      <p>Count1: {dabbleCount}</p>
      <button onClick={()=>setCount1(count1+1)}>increment1</button>

      <p>Count2: {count2}</p>
      <button onClick={()=>setCount2(count2+1)}>increment2</button>
    </div>
  );
}

export default App;
