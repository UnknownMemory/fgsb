const eventSource = new EventSource("http://localhost:8080/api/scoreboard/events");

eventSource.onmessage = function (event) {
  const data = JSON.parse(event.data);
  for (const key in data) {
    const element = document.querySelector(`[data-${key}]`);
    if (element !== null) {
      element.textContent = data[key];
    }
  }
};

eventSource.onerror = (err) => {
  console.error("EventSource failed: ", err);
};
