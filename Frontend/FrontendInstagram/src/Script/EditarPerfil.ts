import { getUserByName } from "../utils/userService";

// Contador de caracteres para la biografía
const textareaBio = document.getElementById("bio") as HTMLTextAreaElement;
const contador = document.getElementById("contador-caracteres") as HTMLSpanElement;
const imagenPerfil = document.getElementById("imagenPerfil") as HTMLImageElement;
const nombreUsuario = document.getElementById("nombreUsuario") as HTMLElement;
let dataUsuario:any = null

async function loadUserData() {
    try {
        const username = localStorage.getItem("user");
        if (!username) {
            throw new Error("No se encontró usuario en localStorage");
        }

        dataUsuario = await getUserByName(username);

        console.log(dataUsuario)
        if (dataUsuario) {
            if (imagenPerfil) {
                console.log(dataUsuario.user.profile_pic.String)
                imagenPerfil.src = (dataUsuario.user.profile_pic.String == "") ? "/public/Assets/fotosinperfil.png" : dataUsuario.user.profile_pic.String
            }
            if (nombreUsuario && dataUsuario.user.username) {
                nombreUsuario.textContent = dataUsuario.user.username;
            }
            if (textareaBio && dataUsuario.bio) {
                textareaBio.value = dataUsuario.bio;
                updateCharacterCount();
            }
        }
    } catch (error) {
        console.error("Error al cargar datos del usuario:", error);
    }
}

function updateCharacterCount() {
    if (textareaBio && contador) {
        contador.textContent = `${textareaBio.value.length}/150`;
    }
}

document.addEventListener("DOMContentLoaded", loadUserData);

if (textareaBio && contador) {
    textareaBio.addEventListener("input", updateCharacterCount);
}

const inputFoto = document.getElementById("input-foto") as HTMLInputElement;
const previewFoto = document.getElementById("preview-foto") as HTMLImageElement;

if (inputFoto && previewFoto) {
    inputFoto.addEventListener("change", (e: Event) => {
        const target = e.target as HTMLInputElement;
        const file = target.files?.[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = (event: ProgressEvent<FileReader>) => {
                if (event.target?.result) {
                    previewFoto.src = event.target.result as string;
                }
            };
            reader.readAsDataURL(file);
        }
    });
}

const formEditarPerfil = document.getElementById("form-editar-perfil") as HTMLFormElement;
if (formEditarPerfil && textareaBio && inputFoto) {
    formEditarPerfil.addEventListener("submit", async (e: SubmitEvent) => {
        e.preventDefault();

        const formData = new FormData();

        if (inputFoto.files?.[0]) {
            formData.append("image", inputFoto.files[0]);
        }
        try {
            const res = await fetch('http://localhost:3000/images/upload', {
                method: 'POST',
                body: formData
            });
            const data = await res.json()
            ActualizarPerfil(data)
            if(data){
                localStorage.removeItem("user")
                window.location.href ="/"
            }
        } catch (err) {
            console.error('Error al subir la imagen:', err);
        }
        console.log(textareaBio.value)
    });
}

async function ActualizarPerfil(data:any) {
    try {
        const res = await fetch('http://localhost:8080/api/ActualizarPerfil', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                bio: textareaBio.value,
                profile_pic: data.url,
                user_id: dataUsuario.user.user_id
            })
        });
        
        if (!res.ok) {
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        
        const datos = await res.json();
        return datos;
    } catch (err) {
        console.error('Error al actualizar el perfil:', err);
        throw err; 
    }
}