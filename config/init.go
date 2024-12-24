package config

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Configuración inicial del servidor web y otras utilidades.
func InitServer() *http.Server {
	port := getEnv("APP_PORT", "8080")
	log.Printf("Iniciando el servidor en el puerto %s...\n", port)

	server := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	return server
}

// Obtiene una variable de entorno o un valor predeterminado si no existe.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Función de inicialización general.
func InitApp() {
	log.Println("Configuración inicial del sistema...")
	checkDirectories()
}

// Verifica la existencia de los directorios necesarios y los crea si no existen.
func checkDirectories() {
	requiredDirs := []string{"data"}
	for _, dir := range requiredDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.Mkdir(dir, os.ModePerm)
			if err != nil {
				log.Fatalf("Error al crear el directorio %s: %v", dir, err)
			}
			log.Printf("Directorio creado: %s\n", dir)
		}
	}
	log.Println("Todos los directorios necesarios están configurados.")
}
