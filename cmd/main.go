package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"streamingpeliculas/internal/contenido"
	"streamingpeliculas/internal/usuarios"
	"strings"
	"streamingpeliculas/web"
	"net/http"
	"streamingpeliculas/config"
	"streamingpeliculas/migracion"  // Asegúrate de que la ruta del import sea correcta 

)

func main() {

	// Llamar a la función que migrará los datos
	fmt.Println("Iniciando migración de datos...")
	migracion.MigrarDatos()



    config.InitDB()
    defer func() {
        if config.DB != nil {
            err := config.DB.Close()
            if err != nil {
                log.Printf("Error al cerrar la base de datos: %v", err)
            }
        }
    }()
	




	fmt.Println("Sistema de Streaming Películas en ejecución...")
	// Registrar el servicio de login
    http.HandleFunc("/login", web.Login)
    

	log.Println("Iniciando el sistema StreamingPelículas...")
	fmt.Println("=== Bienvenido al Sistema de Gestión de StreamingPelículas ===")

	go  func ()  {
		log.Fatal(http.ListenAndServe(":8081", nil))  // Escuchar en el puerto 8080
	}()
	
	// Crear o acceder a cuenta de usuario
	user := createOrAccessAccount()

	
	// Obtener contenido disponible
	availableContent := contenido.ObtenerContenido()

	// Mostrar contenido disponible
	fmt.Println("\nBIENVENIDO :D ¿Qué te gustaría ver hoy?")
	for i, content := range availableContent {
		fmt.Printf("%d. %s\n", i+1, content.Titulo)
	}

	// Menú principal
	for {
		fmt.Println("\nOpciones: 'ver' para elegir contenido, 'salir' para salir del sistema.")
		var action string
		fmt.Scanln(&action)

		switch action {
		case "ver":
			seleccionarContenido(user, availableContent)
		case "salir":
			fmt.Println("Saliendo del sistema... ¡Hasta pronto!")
			return
		default:
			fmt.Println("Opción no válida. Inténtalo nuevamente.")
		}
	}
}



// Crear o acceder a cuenta de usuario
func createOrAccessAccount() *usuarios.Usuario {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("¿Ya tienes una cuenta? (si/no)")

	for {
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		switch answer {
		case "si":
			return accederCuenta(reader)
		case "no":
			return crearCuenta(reader)
		default:
			fmt.Println("Respuesta no válida. Por favor, responde con 'si' o 'no'.")
		}
	}
}

// Acceder a cuenta existente
func accederCuenta(reader *bufio.Reader) *usuarios.Usuario {
	for {
		fmt.Print("Introduce tu correo: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		user, err := usuarios.BuscarUsuario(email)
		if err == nil {
			fmt.Print("Introduce tu contraseña: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			if user.Contraseña == password {
				fmt.Println("Acceso exitoso. ¡Bienvenido!")
				return user
			}
			fmt.Println("Contraseña incorrecta. Intenta de nuevo.")
		} else {
			fmt.Println("No se encontró una cuenta con ese correo. ¿Deseas intentar de nuevo o crear una nueva cuenta? (intentar/crear)")
			option, _ := reader.ReadString('\n')
			option = strings.TrimSpace(option)

			if option == "crear" {
				return crearCuenta(reader)
			}
		}
	}
}

// Crear nueva cuenta
func crearCuenta(reader *bufio.Reader) *usuarios.Usuario {
	fmt.Println("Creación de nueva cuenta.")

	fmt.Print("Introduce tu nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Introduce tu correo: ")
	var correo string
	for {
		correo, _ = reader.ReadString('\n')
		correo = strings.TrimSpace(correo)
		if validarCorreo(correo) {
			break
		}
		fmt.Println("Correo no válido. Inténtalo nuevamente.")
	}

	fmt.Print("Introduce tu contraseña: ")
	var contrasena string
	for {
		contrasena, _ = reader.ReadString('\n')
		contrasena = strings.TrimSpace(contrasena)
		if validarContrasena(contrasena) {
			break
		}
		fmt.Println("Contraseña no válida. Debe contener al menos 6 caracteres, incluyendo letras y números.")
	}

	user, _ := usuarios.NuevoUsuario(nombre, correo, contrasena)
	if err := usuarios.GuardarUsuario(user); err != nil {
		log.Fatalf("Error al guardar usuario: %v", err)
	}

	fmt.Println("Cuenta creada exitosamente. ¡Bienvenido!")
	return user
}

// Seleccionar contenido
func seleccionarContenido(user *usuarios.Usuario, availableContent []contenido.Contenido) {
	fmt.Println("\nSelecciona el número del contenido que deseas ver:")
	for i, content := range availableContent {
		fmt.Printf("%d. %s\n", i+1, content.Titulo)
	}

	var seleccion int
	fmt.Scanln(&seleccion)

	if seleccion > 0 && seleccion <= len(availableContent) {
		selectedContent := availableContent[seleccion-1]
		fmt.Printf("Has seleccionado: %s\n", selectedContent.Titulo)
		user.PeliculasVistas = append(user.PeliculasVistas, selectedContent.Titulo)

		if err := usuarios.GuardarUsuario(user); err != nil {
			fmt.Println("Error al guardar usuario:", err)
		}
	} else {
		fmt.Println("Selección inválida. Intenta de nuevo.")
	}
}

// Validación del formato del correo
func validarCorreo(correo string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(correo)
}

// Validación de la contraseña
func validarContrasena(contrasena string) bool {
	if len(contrasena) < 6 {
		return false
	}

	hasLetter := false
	hasNumber := false
	for _, char := range contrasena {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasLetter = true
		}
		if char >= '0' && char <= '9' {
			hasNumber = true
		}
	}
	return hasLetter && hasNumber
}
