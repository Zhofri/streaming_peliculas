package contenido

import (
	"errors"
	"fmt"
	"time"
)

// Contenido representa un contenido del sistema.
type Contenido struct {
	id        int     //
	titulo    string // Título del contenido
	categoria string // Categoría del contenido
	duracion  int    // Duración del contenido en minutos
}

// Implementación de la interfaz InfoPrinter.
func (c Contenido) PrintInfo() {
	fmt.Printf("Contenido: %s, Categoría: %s, Duración: %d minutos\n", c.titulo, c.categoria, c.duracion)
}

// Setter para el título del contenido
func (c *Contenido) SetTitulo(titulo string) error {
	if titulo == "" {
		return errors.New("el título no puede estar vacío")
	}
	c.titulo = titulo
	return nil
}

// Setter para la categoría del contenido
func (c *Contenido) SetCategoria(categoria string) error {
	if categoria == "" {
		return errors.New("la categoría no puede estar vacía")
	}
	c.categoria = categoria
	return nil
}

// Setter para la duración del contenido
func (c *Contenido) SetDuracion(duracion int) error {
	if duracion <= 0 {
		return errors.New("la duración debe ser un valor positivo")
	}
	c.duracion = duracion
	return nil
}
// GetTitulo obtiene el título del contenido
func (c *Contenido) GetTitulo() string {
	return c.titulo
}

// GetDuracion obtiene la duración del contenido
func (c *Contenido) GetDuracion() int {
	return c.duracion
}
// Prueba de Goroutines para la concurrencia.
func TestGoroutines() {
	ch := make(chan string)

	go func() {
		fmt.Println("Goroutine 1 ejecutando...")
		time.Sleep(1 * time.Second)
		ch <- "Resultado de Goroutine 1"
	}()

	go func() {
		fmt.Println("Goroutine 2 ejecutando...")
		time.Sleep(2 * time.Second)
		ch <- "Resultado de Goroutine 2"
	}()

	// Esperar y recibir resultados de los canales
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
	close(ch)
	fmt.Println("Todas las Goroutines finalizadas.")

	
}
