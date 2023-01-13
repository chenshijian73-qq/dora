package tpl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/ghodss/yaml"
	"html"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"unicode"
)

var MyFuncMap = map[string]interface{}{
	"ToLower":      strings.ToLower,
	"ToUpper":      strings.ToUpper,
	"ToGetterName": ToGetterName,
	"ToSetterName": ToSetterName,
	"ToSelector":   ToSelector,
	"ToClassName":  ToClassName,
	"escapeHtml":   EscapeHtml,
	"escapeQuote":  EscapeQuote,
	"ToImport":     ToImport,
}

func ToGetterName(name string) string {
	return "get" + strings.Title(name)
}
func ToSetterName(name string) string {
	return "Set" + strings.Title(name)
}
func ToSelector(name string) string {
	if name == "isEmail" {
		name = "isEmailAndGmail"
	}
	var first = string(name[2])
	if strings.ToLower(string(name[3])) == string(name[3]) {
		first = strings.ToLower(string(name[2]))
	}

	return first + name[3:] + "Validator"
}
func ToClassName(name string) string {
	if name == "isEmail" {
		name = "isEmailAndGmail"
	}
	return strings.ToUpper(string(name[2])) + name[3:] + "ValidatorDirective"
}
func ToImport(name string) string {

	return name[:len(name)-3]
}
func EscapeHtml(name string) string {
	return html.EscapeString(name)
}
func EscapeQuote(name string) string {
	return strings.Replace(name, "'", "\\'", 1)
}
func TestName(t *testing.T) {
	dataFileName := "servers.yaml"
	templateFileName := "servers.tpl"
	template, err := template.New(path.Base(templateFileName)).Funcs(template.FuncMap(MyFuncMap)).Funcs(template.FuncMap(sprig.TxtFuncMap())).ParseFiles(templateFileName)
	if err != nil {
		fmt.Println(err)
	}

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
	outputFileName := strings.TrimSuffix(dataFileName, filepath.Ext(dataFileName)) + ".generated.yaml"
	generateFile(template, "./", outputFileName, data)
}

func generateFile(template *template.Template, outputDirectory string, outputFileName string, data interface{}) {
	absOutputFileName := path.Join(outputDirectory, outputFileName)
	os.MkdirAll(path.Dir(absOutputFileName), os.ModePerm)
	outputFile, err := os.Create(absOutputFileName)
	fmt.Println("Generating file : " + absOutputFileName)
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
