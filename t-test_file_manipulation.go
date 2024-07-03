package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFileCsv(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {

		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}

	return message, nil

}

func changeCommaToSemicolon(data string) (string, error) {
	readerCsv := csv.NewReader(strings.NewReader(data))

	var stringData string = ""
	for {
		record, err := readerCsv.Read()
		if err == io.EOF {
			return stringData, err
		}

		for index, value := range record {

			if index == len(record)-1 {
				stringData += value + "\n"
				continue
			}
			stringData += value + ";"

		}
	}
}

func fileClose(file *os.File) {
	fmt.Println("menutup file")
	file.Close()
}

func addToFileLog(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer fileClose(file)

	file.WriteString(message)
	return nil
}

func main() {
	dataCsv, err := readFileCsv("assets/datacsv.csv")
	if err != nil {
		fmt.Println("ini error main :")
		fmt.Println(err)
		panic("berhenti")
	}

	stringData, err := changeCommaToSemicolon(dataCsv)

	if err != io.EOF {
		fmt.Println(err)
		panic("error changeCommaToSemicolon")
	}

	if err == io.EOF {
		fmt.Println(stringData)
		addToFileLog("sample.log", stringData)
	}

}
