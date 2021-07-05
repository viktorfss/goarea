package goarea

import "math"

// Pi Constate Pi
const Pi = 3.1415926535

// Circ Calcula a área do círculo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect Calcula a área do retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Cálcula a área do triangulo equilátero
// Ps: Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2.0
}
