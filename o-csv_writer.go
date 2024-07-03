package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Eko", "Kurniawan", "Kannedy"})
	_ = writer.Write([]string{"Andi", "Yanto"})
	_ = writer.Write([]string{"Endang", "Sukaesih"})

	writer.Flush()
}
