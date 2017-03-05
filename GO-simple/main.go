package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pmylund/sortutil"
)

type city struct {
	id                  int
	latitude, longitude float64
}

type chromosome struct {
	id      int
	fitness float64
	cities  []city
}

func main() {

	if len(os.Args) > 4 {
		fileDirectory, populationSize, generations, crossingRate, mutation := readArgs()
		searchInstance(fileDirectory, populationSize, generations, crossingRate, mutation)

	} else {
		fmt.Println("Passe os argumentos para executar o experimento")
		fmt.Println("ARGS: fileDirectory, populationSize, generations, mutation")
	}
}

func searchInstance(fileDirectory, populationSizeString, generations, crossingRate, mutation string) {
	fileCities := readCity(fileDirectory)
	cities := getArrayOfCities(fileCities)

	populationSize, err := strconv.Atoi(populationSizeString)
	if err != nil {
		fmt.Println("O tamanho da polulação é inválido\n", err)
		os.Exit(1)
	}

	primeiraGeracao := createInitialPopulationWithFitness(cities, populationSize)

	sortutil.AscByField(primeiraGeracao, "fitness")
	natutalSelection(populationSize, primeiraGeracao)
}

func natutalSelection(populationSize int, geracao []chromosome) {
	var selectedIndividuals []chromosome

	selectedIndividuals = elitism(populationSize, geracao)

}

func elitism(populationSize int, geracao []chromosome) []chromosome {
	var percent = (populationSize) * 25 / 100
	if percent%2 != 0 {
		percent++
	}
	return geracao[0:int(percent)]
}

func calculateTotalFitness(populationSize int, geracao []chromosome) float64 {
	var totalFitness float64
	for _, valor := range geracao {
		totalFitness += valor.fitness
	}
	return totalFitness
}

func createInitialPopulationWithFitness(cities []city, populationSize int) []chromosome {
	primeiraGeracao := []chromosome{}

	for index := 0; index < populationSize; index++ {
		primeiraGeracao = append(primeiraGeracao, createChromosomeOfInitialPopulation(cities, index))
	}

	return primeiraGeracao
}

func crossOver() {
	//OX Crossover - CrossOver Uniforme
	//cruzamento Uniforme Baseado em Ordem
}

func createChromosomeOfInitialPopulation(cities []city, index int) chromosome {
	individuo := shuffle(cities)
	fitness := calculateFitness(shuffle(cities))
	return chromosome{id: index, fitness: fitness, cities: individuo}
}

func readArgs() (string, string, string, string, string) {
	//fileDirectory populationSize generations TaxadeCruzamento crossingRate mutation
	return os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5]
}

func calculateFitness(cities []city) float64 {

	var length = len(cities) - 2
	var fitness float64

	for index := 0; index <= length; index++ {
		fitness += calculateDistanceCoordenate(cities[index], cities[index+1])
	}
	fitness += calculateDistanceCoordenate(cities[length+1], cities[0])
	return fitness

}

func calculateDistanceCoordenate(cidadeOrigem, cidadeDestino city) float64 {

	return 6371 * math.Acos(math.Cos(math.Pi*(90-cidadeDestino.latitude)/180)*math.Cos((90-cidadeOrigem.latitude)*math.Pi/180)+math.Sin((90-cidadeDestino.latitude)*math.Pi/180)*math.Sin((90-cidadeOrigem.latitude)*math.Pi/180)*math.Cos((cidadeOrigem.longitude-cidadeDestino.longitude)*math.Pi/180))

}

func getArrayOfCities(fileCities *bufio.Reader) []city {

	line, err := readLine(fileCities)
	var cities = []city{}
	for err == nil {
		cities = addCity(cities, line)
		line, err = readLine(fileCities)
	}

	rand.Seed(time.Now().UnixNano())
	cities = shuffle(cities)

	return cities
}

func shuffle(cities []city) []city {
	for index := range cities {
		rand := rand.Intn(index + 1)
		cities[index], cities[rand] = cities[rand], cities[index]
	}
	return cities
}

func addCity(cities []city, line string) []city {
	id, latitude, longitude := convertLineOfCity(line)

	cidade := city{id: id, latitude: latitude, longitude: longitude}
	cities = append(cities, cidade)
	return cities
}

func convertLineOfCity(line string) (id int, latitude float64, longitude float64) {

	var err error

	idString := strings.Split(line, " ")[0]
	latitudeString := strings.Split(line, " ")[1]
	longitudeString := strings.Split(line, " ")[2]

	id, err = strconv.Atoi(idString)
	latitude, err = strconv.ParseFloat(latitudeString, 64)
	longitude, err = strconv.ParseFloat(longitudeString, 64)

	if err != nil {
		fmt.Printf("Erro na conversao de uma das linhas das cidades: %v\n", err)
		os.Exit(1)
	}

	return id, latitude, longitude
}

func readCity(fileDirectory string) *bufio.Reader {
	fileCities, err := os.Open(fileDirectory)
	if err != nil {
		fmt.Printf("Erro ao abrir o Arquivo: %v\n", err)
		os.Exit(1)
	}
	return bufio.NewReader(fileCities)
}

func readLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
