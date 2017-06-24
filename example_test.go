package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

// exampleData holds the contents of testdata/example.csv.
var exampleData [][]string

func init() {
	f, err := os.Open("testdata/example.csv")
	if err != nil {
		panic(fmt.Sprintf("error opening example data: %v", err))
	}
	exampleData, err = csv.NewReader(f).ReadAll()
	if err != nil {
		panic(fmt.Sprintf("error reading example data: %v", err))
	}
	f.Close()
}
