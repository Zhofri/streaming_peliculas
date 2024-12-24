package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"streamingpeliculas/internal/usuarios"
	
)

// RegistroUsuarioHandler maneja la solicitud de registro de un nuevo usuario
func RegistroUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar que la solicitud sea POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Leer los datos de la solicitud
	var usuario usuarios.Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "No se pudo leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}

	// Deserializar los datos JSON en una estructura Usuario
	err = json.Unmarshal(body, &usuario)
	if err != nil {
		http.Error(w, "Error al procesar los datos", http.StatusBadRequest)
		return
	}

	// Validar los datos de entrada
	if usuario.Nombre == "" || usuario.Correo == "" || usuario.Contraseña == "" {
		http.Error(w, "Faltan datos requeridos", http.StatusBadRequest)
		return
	}

	// Validación de formato de correo (simplificada)
	if !isValidEmail(usuario.Correo) {
		http.Error(w, "Correo inválido", http.StatusBadRequest)
		return
	}

	// Validar que la contraseña tenga al menos 6 caracteres y contenga números y letras
	if len(usuario.Contraseña) < 6 || !containsNumber(usuario.Contraseña) || !containsLetter(usuario.Contraseña) {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres y contener números y letras", http.StatusBadRequest)
		return
	}

	// Guardar el nuevo usuario en el archivo JSON
	err = guardarUsuario(usuario)
	if err != nil {
		http.Error(w, "No se pudo guardar el usuario", http.StatusInternalServerError)
		return
	}

	// Enviar respuesta exitosa
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuario registrado exitosamente"))
}

// Función para validar el formato del correo (simplificada)
func isValidEmail(email string) bool {
	// Aquí podemos agregar una expresión regular más robusta para validar correos electrónicos
	return len(email) > 0
}

// Función para verificar si una cadena contiene al menos un número
func containsNumber(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

// Función para verificar si una cadena contiene al menos una letra
func containsLetter(s string) bool {
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			return true
		}
	}
	return false
}

// Función para guardar el usuario en un archivo JSON
func guardarUsuario(usuario usuarios.Usuario) error {
	// Leer el archivo existente
	usuariosFile := "usuarios.json"
	data, err := ioutil.ReadFile(usuariosFile)
	if err != nil {
		return fmt.Errorf("error al leer el archivo de usuarios: %v", err)
	}

	// Deserializar el archivo JSON en una lista de usuarios
	var usuariosList []usuarios.Usuario
	err = json.Unmarshal(data, &usuariosList)
	if err != nil {
		return fmt.Errorf("error al procesar los datos del archivo: %v", err)
	}

	// Agregar el nuevo usuario
	usuariosList = append(usuariosList, usuario)

	// Serializar la lista de usuarios de nuevo a JSON
	usuariosData, err := json.MarshalIndent(usuariosList, "", "  ")
	if err != nil {
		return fmt.Errorf("error al convertir los datos de usuarios a JSON: %v", err)
	}

	// Escribir los datos actualizados al archivo
	err = ioutil.WriteFile(usuariosFile, usuariosData, 0644)
	if err != nil {
		return fmt.Errorf("error al guardar el archivo de usuarios: %v", err)
	}

	return nil
}
