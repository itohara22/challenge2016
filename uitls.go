package main

import (
	"encoding/csv"
	"log"
	"os"
)

func Read_file(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Panic(err.Error())
	}

	reader_csv := csv.NewReader(f)
	data, err := reader_csv.ReadAll()
	if err != nil {
		log.Panic(err.Error())
	}

	data_wihtout_header := data[1:]
	return data_wihtout_header
}
