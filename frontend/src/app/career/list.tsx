"use client";

import { useEffect, useState } from "react";

const List = () => {
  const [num, setNum] = useState(0);
  useEffect(() => {
    console.log("Career Sample Test");
    fetch("http://localhost:8080/max")
      .then((response) => response.json())
      .then((json) => {
        console.log(json);
        setNum(json.result);
      })
      .catch((error) => console.error(error));
  }, []);

  return (
    <div>
      <h1>Career Sample Test</h1>
      <h2>{num}</h2>
    </div>
  );
};

export default List;
