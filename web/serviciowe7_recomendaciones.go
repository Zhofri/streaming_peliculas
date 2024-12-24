package web

import (
	"encoding/json"
	
	"net/http"
	
)

// Estructura para la respuesta de recomendaciones
type RecomendacionesResponse struct {
	Recomendaciones []string `json:"recomendaciones"`
}

// servicioWe7Recomendaciones maneja las solicitudes para generar recomendaciones para un usuario
func servicioWe7Recomendaciones(w http.ResponseWriter, r *http.Request) {
	// Establecer el tipo de contenido como JSON
	w.Header().Set("Content-Type", "application/json")

	// Obtener el parámetro 'usuario_id' desde la URL
	usuarioID := r.URL.Query().Get("usuario_id")
	if usuarioID == "" {
		http.Error(w, "Falta el parámetro usuario_id", http.StatusBadRequest)
		return
	}

	
	

	}


// generarRecomendaciones toma un historial de contenido y genera recomendaciones
func generarRecomendaciones(historial []string) []string {
	// Lista simulada de títulos recomendados
	titulosDisponibles := []string{
		"Película A", "Película B", "Película C", "Película D", "Película E",
		"Película F", "Película G", "Película H", "Película I", "Película J",
	}

	// Lógica simple para filtrar recomendaciones (por ejemplo, títulos no vistos)
	recomendaciones := []string{}
	for _, titulo := range titulosDisponibles {
		// Si el título no está en el historial, es una recomendación
		if !contiene(historial, titulo) {
			recomendaciones = append(recomendaciones, titulo)
		}
	}

	return recomendaciones
}

// contiene verifica si un título ya está en el historial
func contiene(historial []string, titulo string) bool {
	for _, h := range historial {
		if h == titulo {
			return true
		}
	}
	return false
}

// respondWithJSON serializa la respuesta en formato JSON y la escribe en el ResponseWriter
func respondWithJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
