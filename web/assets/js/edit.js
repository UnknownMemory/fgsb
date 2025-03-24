const form = document.querySelector("#edit-scoreboard");

const UpdateScoreboard = async () => {
  try {
    const res = await fetch("http://localhost:8080/api/scoreboard/update", {
      method: "POST",
      body: new FormData(form),
    });
    await res;
  } catch (e) {
    console.error(e);
  }
};

form.addEventListener("submit", (e) => {
  e.preventDefault();
  UpdateScoreboard();
});
