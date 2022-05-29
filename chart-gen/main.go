package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/proullon/ramsql/driver"
	"golang.org/x/image/colornames"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type PerformancePoint struct {
	ArraySize   float64 `db:"array_size"`
	Performance float64 `db:"performance"`
}

func main() {
	db, err := sqlx.Open("ramsql", "benchmarks")
	handleError(err)
	query := `
		CREATE TABLE benchmarks (
			array_size INT,
			method_name TEXT,
			iterations_number TEXT,
			performance TEXT,
			performance_metric TEXT,
			memory TEXT,
			memory_metric TEXT,
			allocations TEXT,
			allocations_metric TEXT
		);
	`
	_, err = db.Exec(query)
	handleError(err)
	reader := bufio.NewScanner(os.Stdin)
	methods := map[string]struct{}{}

	for reader.Scan() {
		columns := strings.Split(reader.Text(), ",")
		var arraySize int
		arraySize, err = strconv.Atoi(columns[0])
		handleError(err)

		_, err = db.Exec(`
		INSERT INTO benchmarks (
			array_size, method_name, iterations_number,
			performance, performance_metric, memory,
			memory_metric, allocations, allocations_metric
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`, arraySize,
			columns[1], columns[2], columns[3], columns[4],
			columns[5], columns[6], columns[7], columns[8])
		handleError(err)

		methods[columns[1]] = struct{}{}
	}

	methodNames := make([]string, 0, len(methods))

	for key := range methods {
		methodNames = append(methodNames, key)
	}

	sort.Strings(methodNames)

	pl := plot.New()
	pl.Title.Text = "Benchmarks"

	query = `
			SELECT array_size, performance
			FROM benchmarks
			WHERE method_name = ?
			ORDER BY array_size`

	for i, method := range methodNames {
		if method == "YAML" {
			continue
		}

		var xys plotter.XYs
		var points []PerformancePoint
		err = db.Select(&points, query, method)
		handleError(err)

		if points[0].Performance > 1 {
			continue
		}

		fmt.Println(method)

		for _, point := range points {
			xys = append(xys, plotter.XY{X: point.ArraySize, Y: point.Performance})
		}

		line, err := plotter.NewLine(xys)
		handleError(err)
		line.Color = colornames.Map[colornames.Names[len(colornames.Names)-i-8]]
		line.Width = 1.5

		pl.Add(line)
	}

	err = pl.Save(15*vg.Centimeter, 15*vg.Centimeter, "plot.png")
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
