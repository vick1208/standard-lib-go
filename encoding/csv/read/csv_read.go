package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvStr := "eko,stefanus,soegianto\n" +
		"budi,pratama,nugraha\n" +
		"joko,morro,diah"
	reader := csv.NewReader(strings.NewReader(csvStr))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
