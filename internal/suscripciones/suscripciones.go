package suscripciones

import (
	"errors"
	"fmt"
)

type Suscripcion struct {
	id        int    // ID de la suscripción
	plan      string // Plan de la suscripción
	usuarioID int    // ID del usuario asociado
}

// Implementación de la interfaz InfoPrinter.
func (s Suscripcion) PrintInfo() {
	fmt.Printf("Suscripción: Plan %s, Usuario ID: %d\n", s.plan, s.usuarioID)
}

// Setter para el plan de la suscripción
func (s *Suscripcion) SetPlan(plan string) error {
	if plan == "" {
		return errors.New("el plan no puede estar vacío")
	}
	s.plan = plan
	return nil
}

// Setter para el ID del usuario asociado
func (s *Suscripcion) SetUsuarioID(usuarioID int) error {
	if usuarioID <= 0 {
		return errors.New("el ID de usuario debe ser un valor positivo")
	}
	s.usuarioID = usuarioID
	return nil
}

// Función que crea una suscripción predeterminada para el usuario
func CreateDefaultSubscription(usuarioID int) Suscripcion {
	return Suscripcion{
		plan:      "Premium", // Puedes elegir cualquier plan como predeterminado
		usuarioID: usuarioID,
	}
}
