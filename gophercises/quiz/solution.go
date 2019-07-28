package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
)

func main()  {
	readCSV("gophercises/quiz/problems.csv")
}

// Function to read CSV from DIR.
func readCSV(file_path string){
	// Load the csv file from path
	f, _ := os.Open(file_path)

	// Create new reader.
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		// Stop at EOF
		if err == io.EOF {
			break
		} else if err != nil{
			panic(err)
		}

		fmt.Println(record[0], record[1])
	}
}