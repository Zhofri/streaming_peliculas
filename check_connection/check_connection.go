package main

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
	connectionString := fmt.Sprintf("sqlserver://%s:%d?databse=%s&integratedSecurity=true",
		server, port, database)


	// Intentar abrir la conexión
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatalf("No se pudo establecer la conexión con la base de datos: %v", err)
	} else {
		fmt.Println("Conexión exitosa con la base de datos!")
	}
}
