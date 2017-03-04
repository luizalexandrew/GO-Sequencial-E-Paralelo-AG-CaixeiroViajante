package main

import (
	"fmt"
	"math"
)

func main() {

	// var cidades = [][]int{{1, 4}, {1, 1}, {3, 4}, {4, 2}, {4, 6}, {5, 4}, {7, 2}, {7, 6}}

	cidades := [][][]int{
		{{40, 10}, {2, 6}, {3, 4}, {55, 67}, {4, 6}, {5, 4}, {7, 2}, {7, 6}},
		{{55, 67}, {4, 6}, {5, 4}, {40, 10}, {2, 6}, {3, 4}, {7, 2}, {7, 6}},
		{{55, 67}, {4, 6}, {5, 4}, {7, 2}, {7, 6}, {40, 10}, {2, 6}, {3, 4}},
		{{55, 67}, {4, 6}, {5, 4}, {7, 2}, {40, 10}, {2, 6}, {3, 4}, {7, 6}},
		{{3, 4}, {55, 67}, {4, 6}, {5, 4}, {7, 2}, {7, 6}, {40, 10}, {2, 6}},
		{{3, 4}, {55, 67}, {4, 6}, {40, 10}, {2, 6}, {5, 4}, {7, 2}, {7, 6}},
		{{4, 6}, {5, 4}, {7, 2}, {7, 6}, {40, 10}, {2, 6}, {3, 4}, {55, 67}}}

	fitness(cidades)

}

func fitness(rota [][][]int) {

	for index := 0; index < 20000; index++ {
		for _, valor := range rota {
			go calcularFitness(valor)
		}
	}

}

func calcularFitness(rota [][]int) {

	var length = len(rota) - 2
	var fitness float64
	for index := 0; index <= length; index++ {
		fitness += calcularDistancia(rota[index], rota[index+1])
	}
	fitness += calcularDistancia(rota[length+1], rota[0])

	fmt.Print("Total ", fitness, "\n")

}

func calcularDistancia(cidade1 []int, cidade2 []int) float64 {
	return math.Sqrt(math.Pow(float64((cidade2[0])-(cidade1[0])), 2) + math.Pow(float64((cidade2[1])-(cidade1[1])), 2))
}

// c := [][][]int{
// 	{{2, 2}, {2, 6}, {3, 4}, {4, 2}, {4, 6}, {5, 4}, {7, 2}, {7, 6}},
// 	{{4, 2}, {4, 6}, {5, 4}, {2, 2}, {2, 6}, {3, 4}, {7, 2}, {7, 6}},
// 	{{4, 2}, {4, 6}, {5, 4}, {7, 2}, {7, 6}, {2, 2}, {2, 6}, {3, 4}},
// 	{{4, 2}, {4, 6}, {5, 4}, {7, 2}, {2, 2}, {2, 6}, {3, 4}, {7, 6}},
// 	{{3, 4}, {4, 2}, {4, 6}, {5, 4}, {7, 2}, {7, 6}, {2, 2}, {2, 6}},
// 	{{4, 6}, {5, 4}, {7, 2}, {7, 6}, {2, 2}, {2, 6}, {3, 4}, {4, 2}},
// 	{{4, 2}, {4, 6}, {5, 4}, {7, 2}, {7, 6}, {2, 2}, {2, 6}, {3, 4}}}
