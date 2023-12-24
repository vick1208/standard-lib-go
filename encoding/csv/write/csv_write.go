package main

import (
	"encoding/csv"
	"os"
)

func main() {
	write := csv.NewWriter(os.Stdout)

	_ = write.Write([]string{"Eko", "Soegianto", "Wibowo"})
	_ = write.Write([]string{"Obaja", "Tanto", "Kurniawan"})
	_ = write.Write([]string{"Debora", "Putri", "Susanto"})

	write.Flush()
}
