package web

import (
	

	"net/http"
	"streamingpeliculas/internal/contenido"
	
)

// Estructura para la respuesta del servicio
type ListaContenidoResponse struct {
	Contenido []contenido.Contenido `json:"contenido"`
	Mensaje   string                `json:"mensaje"`
}

// Servicio web 3: Listar contenido
func ListarContenido(w http.ResponseWriter, r *http.Request) {
	// Seteamos los headers para que la respuesta sea en formato JSON
	w.Header().Set("Content-Type", "application/json")

	// Comprobamos si el usuario está autenticado
	usuarioID := r.URL.Query().Get("usuario_id") // Suponiendo que el ID de usuario se pasa como parámetro
	if usuarioID == "" {
		http.Error(w, "Usuario no autenticado", http.StatusUnauthorized)
		return
	}

	


	

	
}