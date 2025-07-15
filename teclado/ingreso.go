package teclado

import (
	"bufio"
	"os"
)

func Ingreso() string {
	scanner := bufio.NewScanner(os.Stdin)
	var texto string
	if scanner.Scan() {
		texto = scanner.Text()
	}
	return texto
}
