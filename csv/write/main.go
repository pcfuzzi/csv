package main

import (
	"encoding/csv"
	"fmt"
	"os"
)
func main () {

	testdata := [][]string{
		{"Name",     "Nachname",    "Alter",  "Haarfarbe",    "Beruf"},
		{"Adriano",  "Christmann",  "53",     "Braun",        "Devops"},
		{"Tom",      "van Dujin",   "29",     "Braun",        "Programierer",},
		{"Michael",  "Leipner",     "29",     "Dunkelblond",  "IT-Berater",},
		{"Serdar",   "Sadkoglu",    "29",     "Dunkelbraun",  "Programierer",},
	}

	file, err := os.Create ("writefile.csv")
	if err != nil {
		fmt.Print("Error:", err)
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, record := range testdata{
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error: ", err)
			panic(err)

		}
	}
writer.Flush()

if err := writer.Error(); err != nil{
	fmt.Print("Error: ", err)
	panic(err)
}
}


