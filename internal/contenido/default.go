package contenido

import "fmt"

// Función para listar contenido predeterminado
func ListarContenidoDefault() {
	fmt.Println("Contenido disponible:")
	for _, contenido := range ContenidoDefault {
		fmt.Printf("ID: %d, Título: %s, Categoría: %s\n", contenido.ID, contenido.Titulo, contenido.Categoria)
	}
}

// Función para obtener los detalles de un contenido por ID
func ObtenerDetallesContenido(id int) *Contenido {
	for _, contenido := range ContenidoDefault {
		if contenido.ID == id {
			return &contenido
		}
	}
	return nil
}
