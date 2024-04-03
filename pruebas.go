package main

import "fmt"

func seis() {}

func siete() {}

type Red struct {
	ID       int
	Nombre   string
	Tipo     string
	Subredes []string
}

// Función para buscar una red por nombre
func buscarRedPorNombre(redes []Red, nombre string) *Red {
	for _, red := range redes {
		if red.Nombre == nombre {
			return &red
		}
	}
	return nil // Si no se encuentra la red, devolvemos nil
}

func main3() {
	// Lista de redes de ejemplo
	redes := []Red{
		{ID: 1, Nombre: "Red1", Tipo: "LAN", Subredes: []string{"Subred1", "Subred2"}},
		{ID: 2, Nombre: "Red2", Tipo: "WAN", Subredes: []string{"Subred3", "Subred4"}},
		{ID: 3, Nombre: "Red3", Tipo: "LAN", Subredes: []string{"Subred5", "Subred6"}},
	}

	// Buscar una red por nombre
	nombre := "Red2"
	redEncontrada := buscarRedPorNombre(redes, nombre)

	// Verificar si se encontró la red
	if redEncontrada != nil {
		fmt.Printf("Red encontrada: %+v\n", *redEncontrada)
	} else {
		fmt.Println("Red no encontrada.")
	}
}
