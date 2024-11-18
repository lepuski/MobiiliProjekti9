import React, { useState, useEffect } from "react";
import axios from "axios";

const App = () => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const apiKey = `"${process.env.REACT_APP_API_KEY}"`;

  useEffect(() => {
    // Testataan API-kutsu ilman parametrejä
    const fetchData = async () => {
      try {
        const response = await axios.get("https://v3.football.api-sports.io/status", {
          headers: {
            "x-rapidapi-host": "v3.football.api-sports.io",
            "x-rapidapi-key": apiKey, // Lisää API-avaimesi tähän
          },
        });
        console.log(process.env);
        setData(response.data); // Tallenna API:n vastaus
        setLoading(false);
      } catch (err) {
        setError("API-kutsua ei voitu suorittaa.");
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <p>Ladataan...</p>;
  if (error) return <p>{error}</p>;

  return (
    <div>
      <h1>API-Football Testeri</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre> {/* Näytetään raakadata */}
    </div>
  );
};

export default App;
