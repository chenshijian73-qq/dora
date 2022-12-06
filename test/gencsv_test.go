package test

import (
	"flag"
	"fmt"
	"github.com/chenshijian73-qq/Doraemon/internal"
	"os"
	"testing"
)

var (
	inputFile   string
	outputFile  string
	outputDelim string
	printHeader bool
	keys        internal.StringArray
)

func init() {
	flag.StringVar(&inputFile, "in", "", "/path/to/input.json (optional; default is stdin)")
	flag.StringVar(&outputFile, "out", "", "/path/to/output.csv (optional; default is stdout)")
	flag.StringVar(&outputDelim, "d", ",", "delimiter used for output values")
	flag.BoolVar(&printHeader, "p", false, "prints header to output")
	flag.Var(&keys, "k", "fields to output")
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func Test_json2csv(t *testing.T) {
	fmt.Println(inputFile, outputFile, keys)
	internal.Json2Csv(inputFile, outputFile)
}
