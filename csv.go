package main

import "os"

type GenericFunction func(args ...interface{}) interface{}

func PopulateCSV(function GenericFunction, filename string, elements int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for i := 0; i <= elements; i++ {

	}
}
