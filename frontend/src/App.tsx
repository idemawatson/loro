import React from "react";
import ReactDOM from "react-dom";
import "assets/css/App.css";

const App: React.FC = () => {
  return <div className="main"></div>;
};

ReactDOM.render(<App />, document.querySelector("#root"));

export default App;
