const form = document.querySelector("#edit");

const UpdateScoreboard = async () => {
  try {
    const res = await fetch("http://localhost:8080/api/v1/scoreboard/update", {
      method: "POST",
      body: new FormData(form),
    });
    await res;
  } catch (e) {
    console.error(e);
  }
};

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  await UpdateScoreboard();
});

form.addEventListener("change", (e) => {
  const element = document.querySelector(`[data-${e.target.id}]`)
  element.textContent = e.target.value
})