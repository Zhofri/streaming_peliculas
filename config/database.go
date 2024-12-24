package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // Driver para SQL Server
)

// Configuración de conexión a la base de datos
const (
	server   = "localhost"        // Cambiar por el nombre/host del servidor
	port     = 1433               // Cambiar si el puerto del servidor SQL es diferente
	database = "StreamingPeliculasDB" // Nombre de la base de datos
)

// DB es una variable global para acceder a la conexión de la base de datos
var DB *sql.DB

// ConectarBaseDatos establece la conexión con SQL Server
func InitDB() {
	// Construcción de la cadena de conexión con autenticación de Windows
	// NOTA: Usamos "integratedSecurity=true" para utilizar las credenciales de Windows
	connectionString := fmt.Sprintf("sqlserver://%s:%d?database=%s&integratedSecurity=true",
		server, port, database)

	// Abrir la conexión
	var err error
	DB, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	// Verificar la conexión
	err = DB.Ping()
	if err != nil {
		log.Fatalf("No se pudo establecer la conexión con la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos establecida exitosamente")
}

// CerrarBaseDatos cierra la conexión con la base de datos
func CerrarBaseDatos() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error al cerrar la conexión con la base de datos: %v", err)
		} else {
			log.Println("Conexión con la base de datos cerrada exitosamente")
		}
	}
}
