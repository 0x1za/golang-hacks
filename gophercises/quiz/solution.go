package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	"bufio"
	"time"
)

func main()  {
	quiz("gophercises/quiz/problems.csv")
}

func quiz(file_path string){
	correct := 0
	// Load the csv file from path
	f, _ := os.Open(file_path)

	// Create new reader.
	r := csv.NewReader(f)
	result, _ := r.ReadAll()
	questions := strconv.Itoa(len(result))
	fmt.Println("You have a total of " + questions +" questions 😁")
	// Start SDTIN reader.
	scanner := bufio.NewScanner(os.Stdin)
	start := time.Now()
	for i := range result{
		duration := int(time.Since(start).Seconds())
		if duration < 30 {
			fmt.Print("What is "+result[i][0]+"? ")
			answer := result[i][1]
			scanner.Scan()
			input := scanner.Text()

			if answer == input {
				correct += 1
			}
		} else {
			fmt.Println("Time Up!")
			break
		}
	}
	println("Your score is " + strconv.Itoa(correct) +"/"+questions)
}