// Get input + focus
let nameElement = document.getElementById("name");
nameElement.focus();

// Setup the auth function
window.auth = function () {

  // Get name
  let name = nameElement.value;

  // Call App.Auth(name)
  window.go.main.App.Auth(name).then((result) => {
    // Update result with data back from App.Auth()
    document.getElementById("result").innerText = result;
  });
};

nameElement.onkeydown = function (e) {
  if (e.code === "Enter") {
    window.auth();
  }
};
