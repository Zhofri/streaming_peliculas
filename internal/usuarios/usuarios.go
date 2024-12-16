package usuarios

import (
	"errors"
	"fmt"
)

type Usuario struct {
	id int    // ID del usuario
	nombre   string // Nombre del usuario
	password string // Contraseña del usuario
	email    string // Correo electrónico del usuario
}

// Implementación de la interfaz InfoPrinter.
func (u Usuario) PrintInfo() {
	fmt.Printf("Usuario: %s, Email: %s\n", u.nombre, u.email)
}

// Setter para el nombre del usuario
func (u *Usuario) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	u.nombre = nombre
	return nil
}

// Setter para la contraseña del usuario
func (u *Usuario) SetPassword(password string) error {
	if len(password) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}
	u.password = password
	return nil
}

// Setter para el correo electrónico del usuario
func (u *Usuario) SetEmail(email string) error {
	if email == "" {
		return errors.New("el correo electrónico no puede estar vacío")
	}
	u.email = email
	return nil
}


// GetNombre obtiene el nombre del usuario.
func (u *Usuario) GetNombre() string {
	return u.nombre
}

// GetPassword obtiene la contraseña del usuario.
func (u *Usuario) GetPassword() string {
	return u.password
}

// GetEmail obtiene el correo electrónico del usuario.
func (u *Usuario) GetEmail() string {
	return u.email
}