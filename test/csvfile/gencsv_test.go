package csvfile

import (
	"flag"
	"fmt"
	"github.com/chenshijian73-qq/doraemon/internal"
	common "github.com/chenshijian73-qq/doraemon/pkg"
	"os"
	"testing"
)

var (
	inputFile   string
	outputFile  string
	outputDelim string
	printHeader bool
)

func init() {
	flag.StringVar(&inputFile, "in", "", "/path/to/input.json (optional; default is stdin)")
	flag.StringVar(&outputFile, "out", "", "/path/to/output.csv (optional; default is stdout)")
	flag.StringVar(&outputDelim, "d", ",", "delimiter used for output values")
	flag.BoolVar(&printHeader, "p", false, "prints header to output")
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

// go test -v gencsv_test.go -run Test_json2csv -in input.json -out output.csv
func Test_json2csv(t *testing.T) {
	fmt.Println(inputFile, outputFile)
	internal.Json2Csv(inputFile, outputFile)
}

// go test -v gencsv_test.go -run Test_yaml2csv -in input.yaml -out output.csv
func Test_yaml2csv(t *testing.T) {
	fmt.Println(inputFile, outputFile)
	err := internal.YamlToCsv(inputFile, outputFile)
	common.CheckErr(err)
}
