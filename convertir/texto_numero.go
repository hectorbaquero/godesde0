package convertir

import (
	"fmt"
	"strconv"
)

func Texto_numero(tx string) int {

	numero, err := strconv.Atoi(tx)

	if err != nil {
		fmt.Println("El valor registrado no es numerico")
		return 0
	}

	return numero

}
