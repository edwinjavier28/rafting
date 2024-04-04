package main

<<<<<<< HEAD
import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)
=======
import "fmt"
>>>>>>> develop

func seis() {}

func siete() {}
<<<<<<< HEAD
func main9() {
	// Crear una nueva sesión de AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Cambia esto al nombre de tu región
	})

	if err != nil {
		fmt.Println("Error al crear la sesión de AWS:", err)
		return
	}

	// Crear un cliente de S3
	svc := s3.New(sess)

	// Llamar a la función ListBuckets para obtener la lista de buckets
	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error al listar buckets:", err)
		return
	}

	// Buscar un bucket específico
	nombreBucket := "nombre-del-bucket-a-buscar"

	for _, bucket := range result.Buckets {
		if *bucket.Name == nombreBucket {
			fmt.Println("Bucket encontrado:", nombreBucket)
			return
		}
	}

	fmt.Println("Bucket no encontrado:", nombreBucket)
=======

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
>>>>>>> develop
}
