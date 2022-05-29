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
	regexFinder, err := regexp.Compile("Benchmark(\\d+)(\\w+)\\-\\d+")
	handleError(err)
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		str := string(regex.ReplaceAll([]byte(reader.Text()), []byte(",")))

		if !strings.HasPrefix(str, "Benchmark") {
			continue
		}

		strs := regexFinder.FindAllStringSubmatch(str, -1)
		parts := strings.Split(str, ",")
		data := append(strs[0][1:], parts[1:]...)
		columns := strings.Join(data, ",")

		fmt.Println(columns)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
