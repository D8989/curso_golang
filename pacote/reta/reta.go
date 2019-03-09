package main

import "math"

// iniciando com letra maiúscula é PUBLICO (visibilidade dentro e fora do pacote.)
// iniciando com letra minuscula é PRIVADO (visivel apenas dentro do pacote)

//Por exemplo...
// Ponto 			- gerará alfo público
// ponto ou _Ponto 	- gerará algo privado

// Ponto representa um ponto no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsavel por calcular a distância entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
