package contenido

import (
	"encoding/json"
	"errors"
	
)

// Estructura que representa un contenido
type Contenido struct {
	ID          int    `json:"id"`          // Identificador único del contenido
	Titulo      string `json:"titulo"`      // Título del contenido
	Categoria   string `json:"categoria"`   // Categoría del contenido (e.g., acción, comedia)
	Descripcion string `json:"descripcion"` // Descripción del contenido
}

// Lista predeterminada de contenido disponible
var ContenidoDefault = []Contenido{
	{ID: 1, Titulo: "Película Aventura", Categoria: "Aventura", Descripcion: "Una película emocionante llena de acción y paisajes impresionantes."},
	{ID: 2, Titulo: "Comedia Familiar", Categoria: "Comedia", Descripcion: "Diversión para toda la familia con esta hilarante película."},
	{ID: 3, Titulo: "Documental Naturaleza", Categoria: "Documental", Descripcion: "Explora los secretos más fascinantes de la naturaleza."},
	{ID: 4, Titulo: "Thriller Psicológico", Categoria: "Suspenso", Descripcion: "Un thriller que te mantendrá en el borde del asiento."},
	{ID: 5, Titulo: "Drama Épico", Categoria: "Drama", Descripcion: "Una historia conmovedora que tocará tu corazón."},
	{ID: 6, Titulo: "Ciencia Ficción", Categoria: "Sci-Fi", Descripcion: "Aventúrate a mundos desconocidos con esta película futurista."},
	{ID: 7, Titulo: "Romance Clásico", Categoria: "Romance", Descripcion: "Un romance que perdura a través del tiempo."},
}

// Función para obtener la lista completa de contenido
func ObtenerContenido() []Contenido {
	return ContenidoDefault
}

// Función para buscar contenido por ID
func BuscarContenidoPorID(id int) (*Contenido, error) {
	for _, contenido := range ContenidoDefault {
		if contenido.ID == id {
			return &contenido, nil
		}
	}
	return nil, errors.New("contenido no encontrado")
}

// Función para buscar contenido por categoría
func BuscarContenidoPorCategoria(categoria string) []Contenido {
	var contenidoFiltrado []Contenido
	for _, contenido := range ContenidoDefault {
		if contenido.Categoria == categoria {
			contenidoFiltrado = append(contenidoFiltrado, contenido)
		}
	}
	return contenidoFiltrado
}

// Función para agregar un nuevo contenido
func AgregarNuevoContenido(id int, titulo, categoria, descripcion string) error {
	// Verificar si ya existe un contenido con el mismo ID
	for _, contenido := range ContenidoDefault {
		if contenido.ID == id {
			return errors.New("ya existe un contenido con ese ID")
		}
	}

	// Crear y agregar el nuevo contenido
	nuevoContenido := Contenido{
		ID:          id,
		Titulo:      titulo,
		Categoria:   categoria,
		Descripcion: descripcion,
	}
	ContenidoDefault = append(ContenidoDefault, nuevoContenido)
	return nil
}

// Función para eliminar contenido por ID
func EliminarContenidoPorID(id int) error {
	for i, contenido := range ContenidoDefault {
		if contenido.ID == id {
			// Eliminar el contenido de la lista
			ContenidoDefault = append(ContenidoDefault[:i], ContenidoDefault[i+1:]...)
			return nil
		}
	}
	return errors.New("contenido no encontrado para eliminar")
}

// Función para serializar el contenido a JSON
func SerializarContenido(contenidos []Contenido) (string, error) {
	contenidoJSON, err := json.Marshal(contenidos)
	if err != nil {
		return "", err
	}
	return string(contenidoJSON), nil
}
