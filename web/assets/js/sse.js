const eventSource = new EventSource("http://localhost:8080/api/scoreboard/events");
eventSource.onmessage = function (event) {
  console.log(event.data);
};

eventSource.onerror = (err) => {
  console.error("EventSource failed: ", err);
};
