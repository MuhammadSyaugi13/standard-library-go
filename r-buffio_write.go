package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello world\n")
	_, _ = writer.WriteString("test buff write\n")
	writer.Flush()
}
