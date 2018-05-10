Desarrollo de una Agenda en GO.

Integrantes:
Borsotti Albano.
Deagustini Agustin.
Soto Yamil.
---------------------------------------------------------------------------------------------------------
La API permite la insercion, modificacion, eliminacion y listar contactos de una agenda.

1-Insertar Contacto.
Mediante Postman:
    POST http://localhost:8080/contactos
    Json del Body:
    {
	"name": "Nombre",
	"lastname": "Apellido",
	"adress": "Direccion",
	"email": "email@dominio.com"
    }
Realizamos la comprobación de la validez de el email por medio de expresiones regulares.


2-Modificar Contacto.
Mediante Postman:
    PUT http://localhost:8080/contactos/"iddelcontacto"
    Json del Body:
    {
	"name": "Nombre",
	"lastname": "Apellido",
	"adress": "Direccion",
	"email": "email@dominio.com"
    }

3-Modificar Contacto.
Mediante Postman:
    DELETE http://localhost:8080/contactos/"iddelcontacto"


4-Listar Contactos.
En el Navegador:
http://localhost:8080/contactos

5-Buscar si existe un contacto con un Nombre.
(Este lo realizamos de modo extra: Funcion ContactSearch)
En el Navegador:
http://localhost:8080/contactos/"NombreABuscar"