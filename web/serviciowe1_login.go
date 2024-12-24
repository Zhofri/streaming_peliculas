package web

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "streamingpeliculas/internal/usuarios"
)

// Estructura para la respuesta del servicio
type responde struct {
    Message string `json:"message"`
}

// Servicio Web 1: Login de usuario
func Login(w http.ResponseWriter, r *http.Request) {
    // Establecer el encabezado para JSON
    w.Header().Set("Content-Type", "application/json")

    // Definir la estructura para la solicitud de login
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Decodificar los datos JSON del cuerpo de la solicitud
    err := json.NewDecoder(r.Body).Decode(&loginData)
    if err != nil {
        http.Error(w, "Error al leer los datos JSON", http.StatusBadRequest)
        return
    }

    // Leer el archivo de usuarios para verificar las credenciales
    usuariosFile, err := ioutil.ReadFile("usuarios.json")
    if err != nil {
        http.Error(w, "Error al leer los datos de usuarios", http.StatusInternalServerError)
        return
    }

    var usuariosList []usuarios.Usuario
    err = json.Unmarshal(usuariosFile, &usuariosList)
    if err != nil {
        http.Error(w, "Error al procesar los datos de usuarios", http.StatusInternalServerError)
        return
    }

    

    
}
