package main

import (
	"os"
	"net/http"
	"io"
	"strconv"
)

func main() {

	err := downloadFile()

	if err != nil {
		panic(err)
	}
}

func downloadFile() error {
	for i:=1; i<=24; i++ {
		println (string(i))
		url := "http://math.slu.edu/~freeman/Discrete_Lecture_"+ strconv.Itoa(i) +".pdf"
		// Create the file
		out, err := os.Create(strconv.Itoa(i) +".pdf")
		if err != nil {
			return err
		}
		defer out.Close()

		// Get the data
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
	}
	return nil
}