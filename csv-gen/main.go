package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	regex, err := regexp.Compile("\\s+")
	handleError(err)
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		str := string(regex.ReplaceAll([]byte(reader.Text()), []byte(",")))

		if !strings.HasPrefix(str, "Benchmark") {
			continue
		}

		fmt.Println(str)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
