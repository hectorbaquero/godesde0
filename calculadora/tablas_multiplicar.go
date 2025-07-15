package calculadora

import (
	"fmt"
)

func TablasMultiplicar(a int) string {
	tabla := fmt.Sprintf("Tabla de Multiplicar del %d: \n", a)
	for i := 1; i <= 10; i++ {
		x := i * a
		tabla += fmt.Sprintf("%d * %d = %d \n", a, i, x)
	}
	return tabla
}
