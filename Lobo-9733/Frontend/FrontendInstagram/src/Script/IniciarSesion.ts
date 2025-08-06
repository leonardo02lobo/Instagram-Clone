const loginForm = document.getElementById("loginForm") as HTMLFormElement

loginForm.addEventListener("submit", async (e) => {
    e.preventDefault()

    const formData = new FormData(loginForm)
    const data = Object.fromEntries(formData.entries())

    const result = await fetch("http://localhost:8080/api/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })

    const datos = await result.json()
    if(result.ok){
        localStorage.setItem("user", datos.Token)
        window.location.href = "/"
    }
})