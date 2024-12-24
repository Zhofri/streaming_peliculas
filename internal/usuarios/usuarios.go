package usuarios

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"
	"streamingpeliculas/config"
)

// Usuario representa un usuario del sistema.
type Usuario struct {
	Nombre          string   `json:"nombre"`
	Correo          string   `json:"correo"`
	Contraseña      string   `json:"contrasena"`
	PeliculasVistas []string `json:"peliculasVistas"`
}

// RegistrarUsuario registra un nuevo usuario en la base de datos
func RegistrarUsuario(nombre, correo, password string) error {
	query := `
		INSERT INTO Usuarios (Nombre, Correo, Password)
		VALUES (@Nombre, @Correo, @Password)`
	_, err := config.DB.Exec(query,
		sql.Named("Nombre", nombre),
		sql.Named("Correo", correo),
		sql.Named("Password", password),
	)
	if err != nil {
		log.Printf("Error al registrar usuario: %v", err)
		return err
	}
	log.Println("Usuario registrado correctamente")
	return nil
}


// NuevoUsuario crea una nueva instancia de Usuario.
func NuevoUsuario(nombre, correo, contrasena string) (*Usuario, error) {
	if err := ValidarUsuario(nombre, correo, contrasena); err != nil {
		return nil, err
	}

	return &Usuario{
		Nombre:          nombre,
		Correo:          correo,
		Contraseña:      contrasena,
		PeliculasVistas: []string{},
	}, nil
}

// GetNombre obtiene el nombre del usuario.
func (u *Usuario) GetNombre() string {
	return u.Nombre
}

// SetNombre establece un nuevo nombre para el usuario.
func (u *Usuario) SetNombre(nombre string) {
	u.Nombre = nombre
}

// GetCorreo obtiene el correo del usuario.
func (u *Usuario) GetCorreo() string {
	return u.Correo
}

// GetContrasena obtiene la contraseña del usuario.
func (u *Usuario) GetContrasena() string {
	return u.Contraseña
}

// GetPeliculasVistas obtiene la lista de películas vistas por el usuario.
func (u *Usuario) GetPeliculasVistas() []string {
	return u.PeliculasVistas
}

// AddPeliculaVista agrega una película a la lista de películas vistas del usuario.
func (u *Usuario) AddPeliculaVista(titulo string) {
	u.PeliculasVistas = append(u.PeliculasVistas, titulo)
}

// GuardarUsuario guarda un usuario en un archivo JSON.
func GuardarUsuario(usuario *Usuario) error {
	usuarios, err := CargarUsuarios()
	if err != nil {
		usuarios = []*Usuario{}
	}

	usuarios = append(usuarios, usuario)

	file, err := os.Create("usuarios.json")
	if err != nil {
		return errors.New("error al guardar el archivo JSON: " + err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(usuarios); err != nil {
		return errors.New("error al codificar los usuarios en JSON: " + err.Error())
	}

	return nil
}

// CargarUsuarios carga todos los usuarios desde el archivo JSON.
func CargarUsuarios() ([]*Usuario, error) {
	file, err := os.Open("usuarios.json")
	if err != nil {
		return nil, errors.New("error al abrir el archivo JSON: " + err.Error())
	}
	defer file.Close()

	var usuarios []*Usuario
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&usuarios); err != nil {
		return nil, errors.New("error al decodificar el archivo JSON: " + err.Error())
	}

	return usuarios, nil
}

// BuscarUsuario busca un usuario por correo.
func BuscarUsuario(correo string) (*Usuario, error) {
	usuarios, err := CargarUsuarios()
	if err != nil {
		return nil, err
	}

	for _, u := range usuarios {
		if u.Correo == correo {
			return u, nil
		}
	}

	return nil, errors.New("usuario no encontrado")
}
