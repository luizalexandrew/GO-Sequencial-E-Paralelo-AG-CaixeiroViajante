package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cidade struct {
	id                  int
	latitude, longitude float64
}

func main() {
	fileCities := readCity()
	createArrayOfCities(fileCities)
}

func createArrayOfCities(fileCities *bufio.Reader) {

	line, err := Readln(fileCities)
	var basenameOpts = []cidade{}
	for err == nil {
		id, latitude, longitude := convertLine(line)

		cidade := cidade{id: id, latitude: latitude, longitude: longitude}

		basenameOpts = append(basenameOpts, cidade)

		line, err = Readln(fileCities)
	}

	fmt.Println(basenameOpts[0])

	// id, latitude, longitude := convertLine(line)

	// for err == nil {
	// 	fmt.Println(id, latitude, longitude)
	// 	line, err = Readln(fileCities)
	// }
}

func convertLine(line string) (id int, latitude float64, longitude float64) {

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

func Readln(r *bufio.Reader) (string, error) {
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
