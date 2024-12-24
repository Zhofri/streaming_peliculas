Nombre del Proyecto: Streaming Management System
Descripción del Proyecto
Este proyecto tiene como objetivo el desarrollo de un sistema de gestión de streaming que permite a los usuarios registrarse, autenticar sus cuentas y ver contenido de películas y series en línea. El sistema incluye la creación de servicios web, la integración de la serialización de datos en formato JSON, y la implementación de varias funcionalidades utilizando el lenguaje de programación Go.

Funcionalidades Principales del Código
Registro de Usuario: Los usuarios pueden crear una cuenta proporcionando su nombre, correo electrónico y una contraseña segura.
Autenticación de Usuario: Los usuarios registrados pueden iniciar sesión mediante la validación de su correo y contraseña.
Visualización de Contenido: Los usuarios pueden ver una lista de contenido disponible (películas y series) tras iniciar sesión.
Servicios Web: El sistema está compuesto por 8 servicios web que gestionan distintas funcionalidades como registro, autenticación, visualización de contenido, entre otros.
Serialización JSON: Los datos de las cuentas de usuario se almacenan en formato JSON, lo que permite una fácil integración y transmisión de los datos.
Gestión de Errores: El sistema maneja errores de entrada, como contraseñas incorrectas o correos mal formateados, de manera eficiente.
Concurrencia y Escalabilidad: El sistema está diseñado para ser escalable y manejar múltiples solicitudes simultáneamente utilizando goroutines y canales.
Miembros del Grupo
Zhofri Joel Guaman Quichimbo: Desarrollador Principal - Responsable del desarrollo del sistema de autenticación, manejo de servicios web y gestión de datos en JSON.

Objetivo del Programa
El objetivo principal de este proyecto es crear una plataforma de streaming segura y eficiente, que permita a los usuarios registrarse, iniciar sesión y consumir contenido en línea. Además, el sistema debe ser escalable y capaz de manejar múltiples usuarios simultáneamente.

Fecha de Creación
Fecha de creación: 22 de diciembre de 2024

Tecnologías Utilizadas
Lenguaje de Programación: Go
Base de Datos: JSON (para almacenamiento temporal), SQL Server (para futuro almacenamiento)
Servicios Web: Go (para implementación de los servicios)
Interfaz de Usuario: HTML (para la visualización del contenido en el futuro)

Pruebas de Servicios Web: Los servicios están disponibles en endpoints definidos en los archivos correspondientes en web/.
Conocimientos Integrados
Unidad 1
Sintaxis básica, condicionales y estructuras de control.
Unidad 2
Arrays, slices y maps para gestionar datos.
Unidad 3
Encapsulación mediante structs y métodos.
Unidad 4
Concurrencia con goroutines y manejo de canales.

Licencia
Este proyecto se desarrolla con fines educativos y no tiene fines comerciales.


NUESTRO CODIGO ESTA SUBIDO EN "MIGRACION-DATOS"
