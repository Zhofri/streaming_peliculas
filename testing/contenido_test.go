package test

import (
	"streamingpeliculas/internal/contenido"
	"testing"
)

func TestSetTitulo(t *testing.T) {
	c := &contenido.Contenido{}
	err := c.SetTitulo("Película de Acción")

	if err != nil {
		t.Errorf("SetTitulo() falló: %v", err)
	}

	if c.GetTitulo() != "Película de Acción" {
		t.Errorf("SetTitulo() no asignó correctamente el título")
	}
}

func TestSetTitulo_EmptyValue(t *testing.T) {
	c := &contenido.Contenido{}
	err := c.SetTitulo("")

	if err == nil {
		t.Error("SetTitulo() debería fallar con un título vacío")
	}
}

func TestSetDuracion(t *testing.T) {
	c := &contenido.Contenido{}
	err := c.SetDuracion(120)

	if err != nil {
		t.Errorf("SetDuracion() falló: %v", err)
	}

	if c.GetDuracion() != 120 {
		t.Errorf("SetDuracion() no asignó correctamente la duración")
	}
}
