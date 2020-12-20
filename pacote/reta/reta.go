package main

import "math"

// inicialicando com letra maiuscula e PUBLIC
// inicializando com letra minuscuia e PRIVATE (dentro do PACOTE)

// exemplo:
// Ponto 			- gerara algo publico
// ponto ou _Ponto 	- gerara algo privado

// Ponto representa uma coordenada no ponto cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia retorna a distancia entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
