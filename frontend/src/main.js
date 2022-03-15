const formElem = document.getElementById("chooseTarget");

// Setup the auth function
window.auth = function (name) {
  // Call App.Auth(name)
  window.go.main.App.Auth(name).then((result) => {
    // Update result with data back from App.Auth()
    document.getElementById("result").innerText = result;
  });
};

formElem.onsubmit = function (e) {
  e.preventDefault();
  const name = document.querySelector('input[name="choice"]:checked').value;
  window.auth(name);
}
