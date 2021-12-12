package main

import "math"

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	//Forma reduzida de riar uma var
	area := PI * math.Pow(raio, 2)

	println("A área da circunferênia é", area)
}
