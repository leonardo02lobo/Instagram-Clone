const Registrar = document.getElementById("Registrar") as HTMLFormElement;

Registrar.addEventListener("submit", async (e) => {
    e.preventDefault();
    const formData = new FormData(Registrar);
    const data = Object.fromEntries(formData.entries());
    const result = await fetch("http://localhost:8080/api/registrar", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    
    const datos = await result.json()

    console.log(datos)
});