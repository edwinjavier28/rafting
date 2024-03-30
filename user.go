package main

import "fmt"

func user(arr []int) []int {

	a := make([]int, 5)
	return a
}

func dos() {

}

type Usuario struct {
	ID       int
	Username string
	Email    string
}

// Función para buscar un usuario por nombre de usuario
func buscarUsuarioPorUsername(usuarios []Usuario, username string) *Usuario {
	for _, usuario := range usuarios {
		if usuario.Username == username {
			return &usuario
		}
	}
	return nil // Si no se encuentra el usuario, devolvemos nil
}

func main4() {
	// Lista de usuarios de ejemplo
	usuarios := []Usuario{
		{ID: 1, Username: "usuario1", Email: "usuario1@example.com"},
		{ID: 2, Username: "usuario2", Email: "usuario2@example.com"},
		{ID: 3, Username: "usuario3", Email: "usuario3@example.com"},
	}

	// Buscar un usuario por nombre de usuario
	username := "usuario2"
	usuarioEncontrado := buscarUsuarioPorUsername(usuarios, username)

	// Verificar si se encontró el usuario
	if usuarioEncontrado != nil {
		fmt.Printf("Usuario encontrado: %+v\n", *usuarioEncontrado)
	} else {
		fmt.Println("Usuario no encontrado.")
	}
}
