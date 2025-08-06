import { getUserByName } from "../utils/userService";

const uploadForm = document.getElementById('uploadForm');
const volver = document.getElementById("volver")

if (uploadForm) {
    uploadForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const imageInput = document.getElementById('imageInput') as HTMLInputElement | null;
        let file = null;
        if (imageInput && imageInput.files) {
            file = imageInput.files[0];
        }
        const formData = new FormData();
        if (file) {
            formData.append('image', file);
        }
        
        try {
            const res = await fetch('http://localhost:3000/images/upload', {
                method: 'POST',
                body: formData
            });
            const data = await res.json()
            CrearPublicacion(data)
        } catch (err) {
            console.error('Error al subir la imagen:', err);
        }
    });
}

volver?.addEventListener("click", (e) => {
    e.preventDefault()
    history.back()
})


async function CrearPublicacion(data: any): Promise<void>{
    const texto = document.getElementById("texto") as HTMLInputElement
    const user = await getUserByName(localStorage.getItem("user"))
    const result = await fetch("http://localhost:8080/api/CrearPublicacion",{
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            "user_id": user.user.user_id,
            "caption": texto?.value,
            "media_url": data.url,
            "media_type": "image"
        })
    })
    const datos = await result.json()
    if(datos){
        window.location.href = "/"
    }
}