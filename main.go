package main

import (
	"fmt"

	"github.com/HectorBaquero/godesde0/archivos"
	"github.com/HectorBaquero/godesde0/calculadora"
	"github.com/HectorBaquero/godesde0/convertir"
	"github.com/HectorBaquero/godesde0/teclado"
)

func main() {
	fmt.Println("Hola Hector, eres el mejor :)")
	fmt.Println("Digita el número de la operación que deseas realizar:")
	fmt.Println("1 Para multiplicar dos numeros")
	fmt.Println("2 Para conocer una tabla de multiplicar")
	opcion := teclado.Ingreso()
	switch opcion {
	case "1":
		nombre_archivo := archivos.CrearArchivo(".", "multiplicacion.txt")
		multiplicacion(nombre_archivo)
		fmt.Println(archivos.LeerArchivo(nombre_archivo))
	case "2":
		nombre_archivo := archivos.CrearArchivo(".", "tabla.txt")
		tabla(nombre_archivo)
		fmt.Println(archivos.LeerArchivo(nombre_archivo))
	default:
		fmt.Println("El valor digitado no corresponde a una opción valida")

	}

}

func multiplicacion(nombre_archivo string) {
	fmt.Println("Por favor ingresa el primer número a multiplicar:")

	ingreso1 := teclado.Ingreso()
	numero1 := convertir.Texto_numero(ingreso1)

	fmt.Println("Por favor ingresa el segundo número a multiplicar:")

	ingreso2 := teclado.Ingreso()
	numero2 := convertir.Texto_numero(ingreso2)

	resultado := calculadora.Multiplicar(numero1, numero2)

	mensaje := fmt.Sprintf("El resultado de multiplicar %d * %d es %d\n", numero1, numero2, resultado)
	archivos.GuardarDatoArchivoNuevo(nombre_archivo, mensaje)
}

func tabla(nombre_archivo string) {
	fmt.Println("Por favor digita el número de la tabla de multiplicar que deseas conocer:")
	ingreso1 := teclado.Ingreso()
	numero1 := convertir.Texto_numero(ingreso1)
	resultado := calculadora.TablasMultiplicar(numero1)
	archivos.GuardarDatoArchivoNuevo(nombre_archivo, resultado)

}
