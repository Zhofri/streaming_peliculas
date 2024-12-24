package migracion

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

type Usuario struct {
	Nombre  string `json:"Nombre"`
	Correo  string `json:"Correo"`
	Password string `json:"Password"`
}

// Función principal para leer los datos desde el archivo JSON y migrarlos a la base de datos
func MigrarDatos() {
	// Abrir el archivo JSON
	file, err := os.Open("usuarios.json")  // Cambia la ruta si es necesario
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	// Leer los datos JSON
	var usuarios []Usuario
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&usuarios)
	if err != nil {
		log.Fatalf("Error al leer el archivo JSON: %v", err)
	}

	// Insertar los datos en la base de datos
	insertarUsuarios(usuarios)
}

// Insertar usuarios en la base de datos SQL Server
func insertarUsuarios(usuarios []Usuario) {
	// Conexión a la base de datos (usar la cadena de conexión adecuada)
	connectionString := "sqlserver://localhost:1433?database=StreamingPeliculasDB&integratedSecurity=true"
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer db.Close()

	// Insertar los datos en la base de datos
	for _, usuario := range usuarios {
		// Verificar si el correo ya existe en la base de datos
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM Usuarios WHERE Correo = @p1", usuario.Correo).Scan(&count)
		if err != nil {
			log.Printf("Error al verificar existencia de correo %s: %v", usuario.Correo, err)
			continue
		}

		// Si el correo ya existe, no insertar
		if count > 0 {
			log.Printf("El correo %s ya existe en la base de datos. No se insertará.\n", usuario.Correo)
			continue
		}

		// Si el correo no existe, insertar el nuevo usuario
		_, err = db.Exec("INSERT INTO Usuarios (Nombre, Correo, Password) VALUES (@p1, @p2, @p3)", usuario.Nombre, usuario.Correo, usuario.Password)
		if err != nil {
			log.Printf("Error al insertar usuario %s: %v", usuario.Nombre, err)
		} else {
			fmt.Printf("Usuario %s insertado exitosamente\n", usuario.Nombre)
		}
	}
}
