package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler para la ruta principal.
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok", "message": "Servidor funcionando correctamente"}
	json.NewEncoder(w).Encode(response)
}

// Inicia el servidor web.
func StartServer() {
	http.HandleFunc("/health", healthCheck)

	fmt.Println("Servidor web escuchando en http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
