package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\n with new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(line)
		fmt.Println(string(line))
	}

}
