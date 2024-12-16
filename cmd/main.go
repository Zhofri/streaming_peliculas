package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"streamingpeliculas/internal/contenido"
	"streamingpeliculas/internal/suscripciones"
	"streamingpeliculas/internal/usuarios"
	"streamingpeliculas/pkg"
	"streamingpeliculas/web"
)

func main() {
	// Mensaje de bienvenida
	fmt.Println("=== Bienvenido al Sistema de Gestión de StreamingPelículas ===")

	// Crear un usuario utilizando setters con manejo de errores
	user := &usuarios.Usuario{}
	if err := user.SetNombre("Juan Pérez"); err != nil {
		log.Println("Error al asignar nombre:", err)
	}
	if err := user.SetPassword("abc123"); err != nil {
		log.Println("Error al asignar contraseña:", err)
	}
	if err := user.SetEmail("juan.perez@example.com"); err != nil {
		log.Println("Error al asignar email:", err)
	}

	// Crear un contenido utilizando setters con manejo de errores
	video := &contenido.Contenido{}
	if err := video.SetTitulo("Documental de Naturaleza"); err != nil {
		log.Println("Error al asignar título:", err)
	}
	if err := video.SetCategoria("Documental"); err != nil {
		log.Println("Error al asignar categoría:", err)
	}
	if err := video.SetDuracion(90); err != nil {
		log.Println("Error al asignar duración:", err)
	}

	// Crear una suscripción utilizando setters con manejo de errores
	sub := &suscripciones.Suscripcion{}
	if err := sub.SetPlan("Premium"); err != nil {
		log.Println("Error al asignar plan:", err)
	}
	if err := sub.SetUsuarioID(1); err != nil {
		log.Println("Error al asignar ID de usuario:", err)
	}

	// Uso de la interfaz para imprimir información de los objetos creados
	var elementos []pkg.InfoPrinter = []pkg.InfoPrinter{user, video, sub}
	for _, elem := range elementos {
		elem.PrintInfo()
	}

	// Pruebas concurrentes
	fmt.Println("\n=== Pruebas Concurrentes ===")
	contenido.TestGoroutines()

	// Simulación de servicio web
	fmt.Println("\n=== Iniciando Servidor Web ===")
	go web.StartServer()

	time.Sleep(2 * time.Second)
	fmt.Println("\nRealizando prueba del servicio web...")

	// Prueba de serialización de datos (JSON)
	data, _ := json.Marshal(user)
	fmt.Println("Datos serializados del usuario (JSON):", string(data))

	// Llamada HTTP para probar el servidor
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		log.Println("Error al probar el servidor:", err)
	} else {
		fmt.Println("Respuesta del servidor:", resp.Status)
	}

	// Mensaje final de éxito
	fmt.Println("\n=== Sistema ejecutado exitosamente ===")
}
