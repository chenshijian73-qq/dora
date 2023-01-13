package tpl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig"
	common "github.com/chenshijian73-qq/doraemon/pkg"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"path"
	"text/template"
	"unicode"
)

func RenderFIle(dataFileName, templateFileName, outputFileName string) {
	template, err := template.New(path.Base(templateFileName)).Funcs(template.FuncMap(MyFuncMap)).Funcs(template.FuncMap(sprig.TxtFuncMap())).ParseFiles(templateFileName)
	common.PrintErrWithPrefixAndExit("new template error", err)
	dataFile, err := os.Open(dataFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer dataFile.Close()

	byteValue, err := ioutil.ReadAll(dataFile)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err = ToJSON(byteValue)
	var data interface{}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println(err)
	}
	generateFile(template, outputFileName, data)
}

func generateFile(template *template.Template, outputFileName string, data interface{}) {
	os.MkdirAll(path.Dir(outputFileName), os.ModePerm)
	outputFile, err := os.Create(outputFileName)
	fmt.Println("Generating file : " + outputFileName)
	defer outputFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = template.Execute(outputFile, data)
	if err != nil {
		fmt.Println(err)
	}
}

func ToJSON(data []byte) ([]byte, error) {
	if hasJSONPrefix(data) {
		return data, nil
	}
	return yaml.YAMLToJSON(data)
}

var jsonPrefix = []byte("{")

func hasJSONPrefix(buf []byte) bool {
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	return bytes.HasPrefix(trim, jsonPrefix)
}
