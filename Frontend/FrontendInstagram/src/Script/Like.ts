import { getUserByName } from "../utils/userService";

const likeButtons = document.querySelectorAll("#Like");
const postIds = document.querySelectorAll("#ID");

async function checkUserLike(postId: number, token: string | null): Promise<boolean> {
    if (!token) return false;
    
    try {
        const response = await fetch('http://localhost:8080/api/posts/check-like', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                post_id: postId,
                token: token
            })
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        return data.has_liked;
    } catch (error) {
        console.error("Error al verificar like:", error);
        return false;
    }
}

function updateLikeIcon(button: Element, isLiked: boolean) {
    button.innerHTML = isLiked ? `
        <svg
            aria-label="Unlike"
            class="h-6 w-6 fill-red-500 cursor-pointer transition-all duration-200 ease-in-out"
            fill="currentColor"
            height="24"
            role="img"
            viewBox="0 0 48 48"
            width="24"
        >
            <title>Unlike</title>
            <path
                d="M34.6 3.1c-4.5 0-7.9 1.8-10.6 5.6-2.7-3.7-6.1-5.5-10.6-5.5C6 3.1 0 9.6 0 17.6c0 7.3 5.4 12 10.6 16.5.6.5 1.3 1.1 1.9 1.7l2.3 2c4.4 3.9 6.6 5.9 7.6 6.5.5.3 1.1.5 1.6.5s1.1-.2 1.6-.5c1-.6 2.8-2.2 7.8-6.8l2-1.8c.7-.6 1.3-1.2 2-1.7C42.7 29.6 48 25 48 17.6c0-8-6-14.5-13.4-14.5z"
            ></path>
        </svg>
    ` : `
        <svg
            aria-label="Like"
            class="h-6 w-6 hover:fill-gray-500 cursor-pointer transition-all duration-200 ease-in-out"
            fill="currentColor"
            height="24"
            role="img"
            viewBox="0 0 24 24"
            width="24"
        >
            <title>Like</title>
            <path
                d="M16.792 3.904A4.989 4.989 0 0 1 21.5 9.122c0 3.072-2.652 4.959-5.197 7.222-2.512 2.243-3.865 3.469-4.303 3.752-.477-.309-2.143-1.823-4.303-3.752C5.141 14.072 2.5 12.167 2.5 9.122a4.989 4.989 0 0 1 4.708-5.218 4.21 4.21 0 0 1 3.675 1.941c.84 1.175.98 1.763 1.12 1.763s.278-.588 1.11-1.766a4.17 4.17 0 0 1 3.679-1.938m0-2a6.04 6.04 0 0 0-4.797 2.127 6.052 6.052 0 0 0-4.787-2.127A6.985 6.985 0 0 0 .5 9.122c0 3.61 2.55 5.827 5.015 7.97.283.246.569.494.853.747l1.027.918a44.998 44.998 0 0 0 3.518 3.018 2 2 0 0 0 2.174 0 45.263 45.263 0 0 0 3.626-3.115l.922-.824c.293-.26.59-.519.885-.774 2.334-2.025 4.98-4.32 4.98-7.94a6.985 6.985 0 0 0-6.708-7.218Z"
            ></path>
        </svg>
    `;
}

async function handleLikeClick(button: Element, postId: number) {
    try {
        const user = await getUserByName(localStorage.getItem("user"));
        if (!user?.user?.user_id) {
            throw new Error("Usuario no autenticado");
        }

        const response = await fetch("http://localhost:8080/api/like", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                post_id: postId,
                user_id: user.user.user_id
            })
        });

        if (!response.ok) {
            throw new Error(`Error HTTP: ${response.status}`);
        }

        const data = await response.json();
        if(data){
            window.location.reload()
        }

        const hasLiked = await checkUserLike(postId, localStorage.getItem("user"));
        updateLikeIcon(button, hasLiked);
        
    } catch (error) {
        console.error("Error al manejar like:", error);
    }
}

async function initializeLikes() {
    const token = localStorage.getItem("user");
    if (!token) return;

    for (let i = 0; i < likeButtons.length; i++) {
        try {
            const postIdText = postIds[i].textContent;
            if (!postIdText) continue;

            const postId = parseInt(postIdText);
            if (isNaN(postId)) continue;

            const hasLiked = await checkUserLike(postId, token);
            updateLikeIcon(likeButtons[i], hasLiked);
        } catch (error) {
            console.error(`Error al inicializar like ${i}:`, error);
        }
    }
}

function setupLikeButtons() {
    likeButtons.forEach((button, i) => {
        const postIdText = postIds[i].textContent;
        if (!postIdText) return;

        const postId = parseInt(postIdText);
        if (isNaN(postId)) return;

        button.addEventListener("click", () => handleLikeClick(button, postId));
    });
}

document.addEventListener("DOMContentLoaded", () => {
    initializeLikes();
    setupLikeButtons();
});