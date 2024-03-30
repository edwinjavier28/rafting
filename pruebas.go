package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func seis() {}

func siete() {}
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
}
