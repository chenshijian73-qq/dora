package cmd

import (
	"github.com/chenshijian73-qq/Doraemon/internal"
	"github.com/spf13/cobra"
	"log"
)

var (
	yamlFile string
)

var yaml2csv = &cobra.Command{
	Use:   "yaml2csv",
	Short: "generate csv from the yaml data",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := internal.YamlToCsv(yamlFile, outFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	yaml2csv.PersistentFlags().StringVarP(&yamlFile, "yamlFile", "y", "", "指定 yaml 文件")
	yaml2csv.PersistentFlags().StringVarP(&outFile, "outFile", "o", "", "指定 输出文件名")
	rootCmd.AddCommand(yaml2csv)
}
