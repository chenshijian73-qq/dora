package internal

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/yukithm/json2csv"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func Json2Csv(jsonFile string, outFile string) {
	b := &bytes.Buffer{}
	wr := json2csv.NewCSVWriter(b)
	j, _ := os.ReadFile(jsonFile)
	var data []map[string]interface{}
	err := json.Unmarshal(j, &data)
	if err != nil {
		log.Fatal(err)
	}
	// convert json to CSV
	csvdata, err := json2csv.JSON2CSV(data)
	if err != nil {
		log.Fatal(err)
	}
	// CSV bytes convert & writing...
	err = wr.WriteCSV(csvdata)
	if err != nil {
		log.Fatal(err)
	}
	wr.Flush()
	got := b.String()

	//Following line prints CSV
	println(got)

	// create file and append if you want
	createFileAppendText(outFile, got)
}

func createFileAppendText(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func YamlToCsv(yamlFIle string, outFile string) (string, error) {

	yamlConfig, err := ioutil.ReadFile(yamlFIle)
	if err != nil {
		return "", err
	}
	// Parse the YAML string into a map
	var yamlMap map[string]interface{}
	if err := yaml.Unmarshal(yamlConfig, &yamlMap); err != nil {
		return "", err
	}

	// Create a new CSV writer
	csvBuf := &bytes.Buffer{}
	csvWriter := csv.NewWriter(csvBuf)

	// Write the headers to the CSV
	var headers []string
	for key := range yamlMap {
		headers = append(headers, key)
	}
	if err := csvWriter.Write(headers); err != nil {
		return "", err
	}

	// Write the data to the CSV
	var data []string
	for _, value := range yamlMap {
		data = append(data, fmt.Sprintf("%v", value))
	}
	if err := csvWriter.Write(data); err != nil {
		return "", err
	}

	// Flush and return the CSV string
	csvWriter.Flush()
	createFileAppendText(outFile, csvBuf.String())
	return csvBuf.String(), nil
}
