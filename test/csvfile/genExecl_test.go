package csvfile

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/chenshijian73-qq/doraemon/internal/common"
	"github.com/xuri/excelize/v2"
	"log"
	"testing"
)

var (
	yamlFile string
	outFile  string
)

func init() {
	flag.StringVar(&yamlFile, "in", "", "/path/to/input.json (optional; default is stdin)")
	flag.StringVar(&outFile, "out", "", "/path/to/output.csv (optional; default is stdout)")
}

//go test -v genExecl_test.go -in alert.yaml -out alert_rules.xlsx
func Test_yaml2excel(t *testing.T) {
	JsonToExcel(common.YamlToJson(yamlFile))
}
func JsonToExcel(jsonStr []byte) {
	var data map[string]interface{}
	if err := json.Unmarshal(jsonStr, &data); err != nil {
		log.Fatal(err)
	}

	// Create a new Excel file.
	f := excelize.NewFile()

	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "层类")
	f.SetCellValue("Sheet1", "B1", "告警名")
	f.SetCellValue("Sheet1", "C1", "级别")
	f.SetCellValue("Sheet1", "D1", "严重层级")
	//f.SetCellValue("Sheet1", "E1", "GroupName")
	f.SetCellValue("Sheet1", "E1", "表达式")
	f.SetCellValue("Sheet1", "F1", "持续时间")
	f.SetCellValue("Sheet1", "G1", "概要")
	f.SetCellValue("Sheet1", "H1", "描述")
	f.SetCellValue("Sheet1", "I1", "帮助文档")

	// Iterate through the alert data and write to Excel file.
	row := 2
	for _, alertData := range data["alert"].([]interface{}) {
		alert := alertData.(map[string]interface{})
		for _, ruleData := range alert["rules"].([]interface{}) {
			rule := ruleData.(map[string]interface{})
			f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), alert["name"])
			f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), rule["alert"])
			f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), rule["labels"].(map[string]interface{})["level"])
			f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), rule["labels"].(map[string]interface{})["severity"])
			//f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), rule["labels"].(map[string]interface{})["group"])
			f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), rule["expr"])
			f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), rule["for"])
			f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), rule["annotations"].(map[string]interface{})["summary"])
			f.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), rule["annotations"].(map[string]interface{})["description"])
			f.SetCellValue("Sheet1", fmt.Sprintf("I%d", row), rule["labels"].(map[string]interface{})["help_url"])
			row++
		}
	}

	// 获取所有行
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}

	// 遍历第一列的所有单元格，查找相邻且值相同的单元格并合并
	var prevValue string
	startRow := 2
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if i == 1 {
			prevValue = row[0]
			continue
		}
		if row[0] != prevValue {
			endRow := i
			if err := f.MergeCell("Sheet1", fmt.Sprintf("A%d", startRow), fmt.Sprintf("A%d", endRow)); err != nil {
				panic(err)
			}
			prevValue = row[0]
			startRow = i + 1
		}
	}
	endRow := len(rows)
	if err := f.MergeCell("Sheet1", fmt.Sprintf("A%d", startRow), fmt.Sprintf("A%d", endRow)); err != nil {
		panic(err)
	}

	// 设置居中
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			//{Type: "diagonalDown", Color: "A020F0", Style: 7},
			//{Type: "diagonalUp", Color: "A020F0", Style: 8},
		},
	})
	f.SetCellStyle("Sheet1", "A1", fmt.Sprintf("I%d", endRow), style)
	// 设置列宽
	f.SetColWidth("Sheet1", "A", "A", 10)
	f.SetColWidth("Sheet1", "B", "B", 20)
	f.SetColWidth("Sheet1", "C", "D", 10)
	f.SetColWidth("Sheet1", "E", "E", 30)
	f.SetColWidth("Sheet1", "F", "F", 10)
	f.SetColWidth("Sheet1", "G", "I", 40)

	// Save the Excel file.
	if err := f.SaveAs(outFile); err != nil {
		log.Fatal(err)
	}
}
