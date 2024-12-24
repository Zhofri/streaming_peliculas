package usuarios

import (
	"errors"
	"regexp"
	"unicode"
)

// ValidarCorreo valida si un correo tiene el formato correcto.
func ValidarCorreo(correo string) error {
	patron := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(patron, correo)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("el formato del correo no es válido")
	}
	return nil
}

// ValidarContrasena valida si una contraseña cumple con los requisitos.
func ValidarContrasena(contrasena string) error {
	if len(contrasena) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}

	var tieneLetra, tieneNumero bool
	for _, char := range contrasena {
		switch {
		case unicode.IsLetter(char):
			tieneLetra = true
		case unicode.IsDigit(char):
			tieneNumero = true
		}
	}

	if !tieneLetra || !tieneNumero {
		return errors.New("la contraseña debe contener al menos una letra y un número")
	}
	return nil
}

// ValidarUsuario valida los datos de un usuario.
func ValidarUsuario(nombre, correo, contrasena string) error {
	if nombre == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	if err := ValidarCorreo(correo); err != nil {
		return err
	}

	if err := ValidarContrasena(contrasena); err != nil {
		return err
	}

	return nil
}
