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

func ReadBoard(csvFilename string, byCol bool) (r [][]string) {
	/*
	   :param: csvFilename: The file containing the board description.
	   Each line in the csv file describes one row or one column of the board,
	   depending on argument --bycols.
	   Lines beginning with '#' and blank lines are ignored.
	   Cells are separated by whitespace. Each cell contains:

	       <cell> ::= <terrain> | <terrain> "." <oilwell> | "." <oilwell>
	       <terrain> ::= 1 | 2 | 3
	       <oilwell> ::= <wellcount> | <wellcount> <variant>
	       <wellcount> ::= 1 | 2 | 3
	       <variant> = "x" | "d"

	   If <terrain> is absent, "1" is assumed.
	   If <variant> == "x" then this is a cell ignored for 3-person games.
	   If <variant> == "d" then this cell contains a derrick (used for testing).

	   :return: A list of rows containing tuples corresponding to the columns.
	            Each cell is the string as defined above.
	*/
	csvFile, err := os.Open(csvFilename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvFile.Close()
	// Parse the file
	nline := 0

	scanner := bufio.NewScanner(csvFile)
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
