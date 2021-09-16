package model

// Estructura de una comunidad
type Community struct {
	// Name nombre de una comunidad. Ej Edteam
	Name string
}

// Communities slice de comunidades
type Communities []Community

// Person Estructura de una persona
type Person struct {
	// Name : nombre de una persona. Ej Karen
	Name string

	//Age edad de la persona
	Age uint8

	//Communities comunidades a las que pertenece una persona.
	Communities Communities
}

type Persons []Person
