package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	//prueft ob die Datei "Testfile.csv vorhanden ist dann öffnet er, wenn sie da ist."
	// der wert wird beiden variablen zugewiesen: wenn err = nil so die fehler meldung.
	file, err := os.Open("writefile.csv")
	if err != nil {
		// es wird eine fehlermeldung ausgeben wenn der wert nil ist, die anwendung wird beendet und geschlossen.
		fmt.Println("Error:", err)
		return
	}
	// csv wird geschlossen.
	defer file.Close()
	// bezeichner reader oeffnet durch ubergabe von variable "file"= os.Open(testfile.csv) datei.
	//und speichert in sich die inhalte der csv-datei
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, row := range rows {
		for _, col := range row{
		fmt.Print(col," ")
	}

	fmt.Println()
}
}