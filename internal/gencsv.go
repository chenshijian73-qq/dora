package internal

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	common "github.com/chenshijian73-qq/doraemon/pkg"
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
	common.CheckErr(err)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	if _, err = f.WriteString(text); err != nil {
		common.CheckErr(err)
	}
}

func YamlToCsv(yamlFile string, csvFile string) error {
	// Parse the YAML string into a map
	yamlData, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return err
	}

	var data interface{}
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return err
	}

	csvData, err := flatten(data)
	if err != nil {
		return err
	}

	file, err := os.Create(csvFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range csvData {
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}
	return err
}

func flatten(data interface{}) ([][]string, error) {
	// 将 yaml 文件中的数据转换为二维 string 数组
	// 使用递归或者循环来处理嵌套的数组
	// 在这里，我们假设 data 是一个嵌套的数组，每个元素都是 map[string]interface{} 类型
	// 例如：
	// [
	//   {"name": "Alice", "age": 30},
	//   {"name": "Bob", "age": 35},
	//   {"name": "Charlie", "age": 40},
	// ]
	// 你需要将它转换为如下的二维 string 数组：
	// [
	//   {"name", "age"},
	//   {"Alice", "30"},
	//   {"Bob", "35"},
	//   {"Charlie", "40"},
	// ]
	// 具体的实现方法可以是这样的：
	keys := make([]string, 0)
	records := make([][]string, 0)
	for _, item := range data.([]interface{}) {
		record := make([]string, 0)
		for key, value := range item.(map[string]interface{}) {
			keys = append(keys, key)
			switch value.(type) {
			case string:
				record = append(record, value.(string))
			case int:
				record = append(record, fmt.Sprintf("%d", value))
			case float64:
				record = append(record, fmt.Sprintf("%f", value))
			default:
				return nil, fmt.Errorf("unsupported value type: %T", value)
			}
		}
		if len(records) == 0 {
			records = append(records, keys)
		}
		records = append(records, record)
	}
	return records, nil
}
