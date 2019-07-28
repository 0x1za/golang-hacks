package main

import (
	"fmt"
	"os"
	"encoding/csv"
)

func main()  {
	quiz("gophercises/quiz/problems.csv")
}

func quiz(file_path string){
	// Load the csv file from path
	f, _ := os.Open(file_path)

	// Create new reader.
	r := csv.NewReader(f)
	result, _ := r.ReadAll()
	for i := range result{
		fmt.Println(result[i][0], result[i][1])
	}
}