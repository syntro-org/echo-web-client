document.addEventListener("DOMContentLoaded", function () {
  var toAuthButton = document.getElementById("toAuth");
  var toRegButton = document.getElementById("toReg");
  var authPanel = document.getElementById("auth");
  var regPanel = document.getElementById("reg");

  toAuthButton.addEventListener("click", function () {
    authPanel.classList.add("active");
    regPanel.classList.remove("active");
  });

  toRegButton.addEventListener("click", function () {
    authPanel.classList.remove("active");
    reg.classList.add("active");
  });
});