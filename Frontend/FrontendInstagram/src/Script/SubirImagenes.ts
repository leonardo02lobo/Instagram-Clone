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
            console.log(data)
        } catch (err) {
            console.error('Error al subir la imagen:', err);
        }
    });
}

volver?.addEventListener("click", (e) => {
    e.preventDefault()
    history.back()
})