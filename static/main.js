function toggleColors() {
  const menu = document.getElementById("colorsMenu");
  if (menu) menu.classList.toggle("show");
}

const colorInputs = document.querySelectorAll('input[name="color"]');

colorInputs.forEach(input => {
  input.addEventListener("change", () => {
    const menu = document.getElementById("colorsMenu");
    const button = document.querySelector(".color-button");

    if (button) {
      button.style.background = input.value;
      button.textContent = "";
    }

    if (menu) menu.classList.remove("show");
  });
});

const textarea = document.querySelector("textarea");

if (textarea) {
  textarea.addEventListener("input", () => {
    if (textarea.value.length > 0) {
      textarea.style.textAlign = "left";
    } else {
      textarea.style.textAlign = "center";
    }
  });
}

window.onload = () => {
  const textarea = document.querySelector("textarea");

  if (textarea && performance.navigation.type === 1) {
    textarea.value = "";
  }
};