package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	
)

// Struct para la respuesta de error
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Manejador para errores generales
func ErrorHandling(w http.ResponseWriter, r *http.Request) {
	// Establecer el encabezado de la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Obtener el tipo de error desde la consulta (query params)
	errorType := r.URL.Query().Get("error")

	// Si no se pasa un tipo de error, devolver error genérico
	if errorType == "" {
		http.Error(w, "Error desconocido", http.StatusBadRequest)
		return
	}

	// Definir la respuesta de error basada en el tipo de error
	var response ErrorResponse

	switch errorType {
	case "login":
		// Error de login: credenciales incorrectas
		response = ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Credenciales incorrectas, por favor verifica tu correo o contraseña.",
		}
	case "registro":
		// Error de registro: datos incompletos o inválidos
		response = ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Datos inválidos para el registro, por favor revisa tu correo y contraseña.",
		}
	case "contenido":
		// Error de contenido: no encontrado o inválido
		response = ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Contenido no encontrado, por favor intenta con otro título.",
		}
	default:
		// Error desconocido
		response = ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Ocurrió un error desconocido, por favor intenta nuevamente más tarde.",
		}
	}

	// Establecer el código de estado HTTP correspondiente
	w.WriteHeader(response.Code)

	// Convertir la respuesta en formato JSON y enviarla
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Si hubo un error al serializar la respuesta, devolver un error genérico
		http.Error(w, "Error al procesar la solicitud", http.StatusInternalServerError)
	}
}

// Función para manejar errores internos (opcional, se puede usar si se desea personalizar más errores)
func InternalServerError(w http.ResponseWriter, err error) {
	// Responder con error interno y mensaje detallado
	http.Error(w, fmt.Sprintf("Error interno: %v", err), http.StatusInternalServerError)
}
