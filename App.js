import React, { useState, useEffect } from "react";
import axios from "axios";
import { StyleSheet, View, Text, Button } from 'react-native';
import Register from './components/register';

const App = () => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const apiKey = process.env.EXPO_PUBLIC_API_KEY;

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
        setData(response.data); // Tallenna API:n vastaus
        setLoading(false);
      } catch (err) {
        setError("API-kutsua ei voitu suorittaa.");
        setLoading(false);
      }
    };

    fetchData();
  }, []);
  

  if (loading) return <Text>Ladataan...</Text>;
  if (error) return <Text>{error}</Text>;

  return (    
    <>
      <Register></Register>
      <Text>
      <Text>API-Football Testeri</Text>
      <Text>{JSON.stringify(data, null, 2)}</Text> {/* Näytetään raakadata */}
      </Text>
    </>

  );
};

export default App;
