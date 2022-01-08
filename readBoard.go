package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func ReadBoard(csvfilename string, byCol bool) (r [][]string) {
	csvfile, err := os.Open(csvfilename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	// Parse the file
	nline := 0

	scanner := bufio.NewScanner(csvfile)
	for scanner.Scan() {
		line := scanner.Text()
		err = scanner.Err()
		// Read each line from csv

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		nline++
		s := strings.Fields(line)
		r = append(r, s)
	}
	if byCol {
		r = transpose(r)
	}
	return
}
