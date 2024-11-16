document.addEventListener("DOMContentLoaded", function () {
  const cityElement = document.getElementById("city");
  const tempElement = document.getElementById("temperature");
  const dateElement = document.getElementById("date");

  function updateWeather(city = "Campinas") {
    fetch(`/weather?city=${city}`)
      .then(response => {
        if (!response.ok) {
          throw new Error("Failed to fetch weather data");
        }
        return response.json();
      })
      .then(data => {
        cityElement.textContent = data.city;
        tempElement.textContent = `${data.temperature.toFixed(1)}°`;
        dateElement.textContent = new Date().toLocaleDateString();
      })
      .catch(error => {
        console.error("Error fetching weather data:", error);
        cityElement.textContent = "Error loading data";
        tempElement.textContent = "--°";
      });
  }

  updateWeather();

});
