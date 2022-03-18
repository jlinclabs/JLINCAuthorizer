const formElem = document.getElementById("chooseTarget");

// Setup the authz function
window.authz = function (service) {
  window.go.main.App.Authz(service).then((result) => {
    document.getElementById("result").innerText = result;
  });
};

formElem.onsubmit = function (e) {
  e.preventDefault();
  const service = document.querySelector('input[name="choice"]:checked').value;
  window.authz(service);
}
