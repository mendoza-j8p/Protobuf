package main

import (
	"github.com/mendoza-j8p/Protobuf" // Importa el paquete generado
	"fmt"
)

func main() {
person := &Protobuf.Person{
	Name:    "John Doe",
	Age:     30,
	Hobbies: []string{"Reading", "Coding"},
}

fmt.Println(person)
}
