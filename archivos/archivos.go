package archivos

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func CrearArchivo(ruta, nombre string) string {
	rutaCompleta := filepath.Join(ruta, nombre)
	archivo, err := os.Create(rutaCompleta)
	if err != nil {
		fmt.Printf("Ocurrio un error en la creación del archivo: \n %s \n", err.Error())
		return ""
	}
	defer archivo.Close()
	return archivo.Name()
}

func GuardarDatoArchivoNuevo(ruta_archivo, dato string) {
	archivo, err := os.OpenFile(ruta_archivo, os.O_APPEND|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Printf("Ocurrio un error al guardar un dato en un archivo nuevo: \n %s \n", err.Error())
		return
	}
	defer archivo.Close()

	_, err = archivo.WriteString(dato)
	if err != nil {
		fmt.Printf("Ocurrio un error al guardar un dato en un archivo nuevo: \n %s \n", err.Error())
		return
	}
	defer archivo.Close()
}

func LeerArchivo(ruta_archivo string) string {
	archivo, err := os.Open(ruta_archivo)
	if err != nil {
		fmt.Printf("Ocurrio un error al leer el archivo: \n %s \n", err.Error())
		return ""
	}
	defer archivo.Close()
	var resultado string
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		resultado += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ocurrio un error al leer el archivo y asignarlo al string: \n %s \n", err.Error())
		return ""
	}

	return resultado
}
