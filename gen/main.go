package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

var (
	templatePath string
	outDir       string
	sizes        = []int{16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072}
	funcs        = template.FuncMap{
		"add": addNumbers,
	}
)

func parseFlags() {
	flag.StringVar(&templatePath, "template", "",
		"A template to generate all the benchmarks")
	flag.StringVar(&outDir, "out", "",
		"An output directory for the generated source code.")

	flag.Parse()
}

func addNumbers(a, b int) string {
	return strconv.Itoa(a + b)
}

func main() {
	parseFlags()

	if !filepath.IsAbs(templatePath) {
		var err error
		templatePath, err = filepath.Abs(templatePath)
		handleError(err)
	}

	if !filepath.IsAbs(outDir) {
		var err error
		outDir, err = filepath.Abs(outDir)
		handleError(err)
	}

	tmpl := template.New(filepath.Base(templatePath)).Funcs(funcs)
	_, err := tmpl.ParseFiles(templatePath)
	handleError(err)

	for _, arraySize := range sizes {
		filename := filepath.Join(outDir,
			fmt.Sprintf("bencmark_%d_test.go", arraySize))

		if !filepath.IsAbs(filename) {
			filename, err = filepath.Abs(filename)
			handleError(err)
		}

		file, err := os.Create(filename)
		handleError(err)
		err = tmpl.Execute(file, map[string]interface{}{
			"arraySize": arraySize,
		})
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
