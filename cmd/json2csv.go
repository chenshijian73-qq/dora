package cmd

import (
	"github.com/chenshijian73-qq/Doraemon/internal"
	"github.com/spf13/cobra"
)

var (
	jsonFile string
	outFile  string
)

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "generate csv from the json data",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Json2Csv(jsonFile, outFile)
	},
}

func init() {
	csvCmd.PersistentFlags().StringVarP(&jsonFile, "jsonFile", "j", "", "指定 json 文件")
	csvCmd.PersistentFlags().StringVarP(&outFile, "outFile", "o", "", "指定 输出文件名")
	rootCmd.AddCommand(csvCmd)
}
