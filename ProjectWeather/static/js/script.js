document.addEventListener("DOMContentLoaded", function(){
    fetch("/weather?city=campinas")
    .then(response => response.json())
    .then(data => {
        document.getElementById("city").textContent=data.city;
        document.getElementById("temperature").textContent = `${data.temperature}°`;
        document.getElementById("date").textContent = new Date().toLocaleDateString();
    })
    .catch(error => console.error("Error fetching weather data:", error));

});