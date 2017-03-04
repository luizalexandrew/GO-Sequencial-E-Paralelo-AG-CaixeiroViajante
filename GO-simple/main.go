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
)

type cidade struct {
	id                  int
	latitude, longitude float64
}

func main() {
	fileCities := readCity()
	cidades := getArrayOfCities(fileCities)

	calculateFitness(cidades)
}

func calculateFitness(cidades []cidade) {

	var length = len(cidades) - 2
	var fitness float64
	for index := 0; index <= length; index++ {
		fitness += calculateDistanceCoordenate(cidades[index], cidades[index+1])
	}
	fitness += calculateDistanceCoordenate(cidades[length+1], cidades[0])

	fmt.Print("Total ", fitness, "\n")

}

func calculateDistanceCoordenate(cidadeOrigem, cidadeDestino cidade) float64 {

	return 6371 * math.Acos(math.Cos(math.Pi*(90-cidadeDestino.latitude)/180)*math.Cos((90-cidadeOrigem.latitude)*math.Pi/180)+math.Sin((90-cidadeDestino.latitude)*math.Pi/180)*math.Sin((90-cidadeOrigem.latitude)*math.Pi/180)*math.Cos((cidadeOrigem.longitude-cidadeDestino.longitude)*math.Pi/180))

}

func getArrayOfCities(fileCities *bufio.Reader) []cidade {

	line, err := readLine(fileCities)
	var cidades = []cidade{}
	for err == nil {
		cidades = addCity(cidades, line)
		line, err = readLine(fileCities)
	}

	rand.Seed(time.Now().UnixNano())
	shuffle(cidades)

	return cidades
}

func shuffle(a []cidade) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func addCity(cidades []cidade, line string) []cidade {
	id, latitude, longitude := convertLineOfCitye(line)

	cidade := cidade{id: id, latitude: latitude, longitude: longitude}
	cidades = append(cidades, cidade)
	return cidades
}

func convertLineOfCitye(line string) (id int, latitude float64, longitude float64) {

	idString := strings.Split(line, " ")[0]
	latitudeString := strings.Split(line, " ")[1]
	longitudeString := strings.Split(line, " ")[2]

	id, _ = strconv.Atoi(idString)
	latitude, _ = strconv.ParseFloat(latitudeString, 64)
	longitude, _ = strconv.ParseFloat(longitudeString, 64)

	return id, latitude, longitude
}

func readCity() *bufio.Reader {
	fileCities, err := os.Open("./cidades.txt")
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
