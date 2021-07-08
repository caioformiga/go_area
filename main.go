package teste_go

import "math"

const Pi = 3.1416

func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}
