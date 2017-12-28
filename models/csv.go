package models

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ReportCSV(r Report) {
	file, err := os.OpenFile("C:\\Users\\Public\\C2NET\\Report\\csv\\productionreport.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	arr := []string{}

	arr = append(arr, strconv.Itoa(r.id))
	arr = append(arr, strconv.Itoa(r.expected))
	arr = append(arr, strconv.Itoa(r.measured))
	arr = append(arr, strconv.Itoa(r.eventcode))
	arr = append(arr, r.problem)
	arr = append(arr, r.notes)

	err = writer.Write(arr)
	checkError("Cannot write to file", err)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
