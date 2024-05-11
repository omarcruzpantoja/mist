

import Left from './left';


import "./App.css";

const App = ():JSX.Element => {
    
  console.log('rendering app');

  return (
    <div className="container">
      <h1>Welcome to Tauri!</h1>


      <Left />
    </div>

  );
}

export default App;
