package test

import (
	"streamingpeliculas/internal/usuarios"
	"testing"
)

func TestSetNombre(t *testing.T) {
	u := &usuarios.Usuario{}
	err := u.SetNombre("Carlos Pérez")

	if err != nil {
		t.Errorf("SetNombre() falló: %v", err)
	}

	if u.GetNombre() != "Carlos Pérez" {
		t.Errorf("SetNombre() no asignó correctamente el nombre")
	}
}

func TestSetPassword(t *testing.T) {
	u := &usuarios.Usuario{}
	err := u.SetPassword("123456")

	if err != nil {
		t.Errorf("SetPassword() falló: %v", err)
	}

	if u.GetPassword() != "123456" {
		t.Errorf("SetPassword() no asignó correctamente la contraseña")
	}
}

func TestSetEmail_Invalid(t *testing.T) {
	u := &usuarios.Usuario{}
	err := u.SetEmail("")

	if err == nil {
		t.Error("SetEmail() debería fallar con un email vacío")
	}
}
