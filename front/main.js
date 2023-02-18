function sendForm() {
    let form = {
        nome: document.getElementById("nome").value,
        email: document.getElementById("email").value,
        password: document.getElementById("html").value,
        description: document.getElementById("description").value,
    }
    fetch("http://localhost:3000/api/form", {
        method: "POST",
        body: JSON.stringify(form),
    })
        .then((response) => response.json())
        .then((data) => console.log(data))

}
