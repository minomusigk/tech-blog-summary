import { useState } from 'react';
import { Configuration, GetPingApiFactory } from './api/axios';

const App = (): JSX.Element => {
  const [message, setMessage] = useState('null');

  const apiConfig = new Configuration({
    basePath: 'http://localhost:8080',
  });
  const client = GetPingApiFactory(apiConfig);

  const onPingButtonClick = async () => {
    const result = await client.getPing();
    setMessage(result.data.message);
  };

  return (
    <div>
      <button onClick={() => onPingButtonClick()}>ping</button>
      <div>{message}</div>
    </div>
  );
};

export default App;
