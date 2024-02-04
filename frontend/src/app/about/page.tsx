"use client";

import { useEffect } from "react";

const AboutPage = () => {
  useEffect(() => {
    console.log("AboutPage Sample Test");
    fetch("http://localhost:8080/min")
      .then((response) => response.json())
      .then((json) => console.log(json))
      .catch((error) => console.error(error));
  }, []);
  return (
    <div>
      <h1>AboutPage Sample Test</h1>
    </div>
  );
};
export default AboutPage;
