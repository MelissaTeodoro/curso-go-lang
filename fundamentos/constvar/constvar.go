package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	// Forma reduzida de riar uma var - Assim criamos uma variável e não constante
	area := PI * m.Pow(raio, 2)

	println("A área da circunferênia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var filhos, casado bool = false, true // Variável filhos recebe false e casado recebe true

	fmt.Println("Tem filhos?", filhos)
	fmt.Println("É casado(a)?", casado)

	// Forma reduzida
	e, f, g := 2, false, "Hello!" // Dessa forma define a tipagem para toda a execução do programa
	fmt.Println(e, f, g)

}
