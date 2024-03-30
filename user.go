package main

import "fmt"

func user(arr []int) []int {

	a := make([]int, 5)
	return a
}

func dos() {

}

type Producto struct {
	ID         int
	Nombre     string
	Precio     float64
	Existencia int
}

// Función para buscar un producto por nombre
func buscarProductoPorNombre(productos []Producto, nombre string) *Producto {
	for _, producto := range productos {
		if producto.Nombre == nombre {
			return &producto
		}
	}
	return nil // Si no se encuentra el producto, devolvemos nil
}

func main6() {
	// Lista de productos de ejemplo
	productos := []Producto{
		{ID: 1, Nombre: "Producto1", Precio: 10.99, Existencia: 100},
		{ID: 2, Nombre: "Producto2", Precio: 15.99, Existencia: 50},
		{ID: 3, Nombre: "Producto3", Precio: 20.99, Existencia: 200},
	}

	// Buscar un producto por nombre
	nombre := "Producto2"
	productoEncontrado := buscarProductoPorNombre(productos, nombre)

	// Verificar si se encontró el producto
	if productoEncontrado != nil {
		fmt.Printf("Producto encontrado: %+v\n", *productoEncontrado)
	} else {
		fmt.Println("Producto no encontrado.")
	}
}
