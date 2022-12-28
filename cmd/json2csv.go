package cmd

import (
	"github.com/chenshijian73-qq/doraemon/internal"
	"github.com/spf13/cobra"
)

var (
	jsonFile string
	outFile  string
)

var json2csv = &cobra.Command{
	Use:   "json2csv",
	Short: "Generate csv from the json data",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Json2Csv(jsonFile, outFile)
	},
}

func init() {
	json2csv.PersistentFlags().StringVarP(&jsonFile, "jsonFile", "j", "", "指定 json 文件")
	json2csv.PersistentFlags().StringVarP(&outFile, "outFile", "o", "", "指定 输出文件名")
	rootCmd.AddCommand(json2csv)
}
