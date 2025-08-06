interface ApiResponse {
  data: any[];
  user: any[];
}

export async function ObtenerPublicaciones(): Promise<ApiResponse> {
    try {
        const response = await fetch("http://localhost:8080/api/ObtenerPublicaciones");
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data: ApiResponse = await response.json();
        return data;
    } catch (e) {
        console.error("Error fetching publicaciones:", (e as Error).message);
        return { data: [], user: [] }; // Retorna un objeto con la estructura esperada
    }
}