package main

import (
	"log"
	"os"
	"strconv"

	"github.com/jreyeshdez/triangleClassifier/cmd/classifier"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("incorrect number of arguments")
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("error parsing: %s ", os.Args[1])
	}

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatalf("error parsing: %s ", os.Args[2])
	}

	c, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("error parsing: %s ", os.Args[3])
	}

	triangle := classifier.NewTriangle(a, b, c)

	result, err := triangle.GetType()
	if err != nil {
		log.Fatalf(err.Error())
	}

	if result == "" {
		log.Fatalf("type of triangle could not be classified")
	}
	log.Println("the type of triangle for the given input is: " + result)

}
