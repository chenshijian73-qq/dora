package cmd

import (
	tpl2 "github.com/chenshijian73-qq/doraemon/internal/tpl"
	"github.com/spf13/cobra"
)

var (
	dataFile string
	tplFile  string
	outFile  string
)

var tpl = &cobra.Command{
	Use:   "tpl",
	Short: "Generate yaml from the dataFile and the tplFile",
	Run: func(cmd *cobra.Command, args []string) {
		tpl2.RenderFIle(dataFile, tplFile, outFile)
	},
}

func init() {
	tpl.PersistentFlags().StringVarP(&dataFile, "dataFile", "d", "", "指定 数据文件，json 或 yaml 文件")
	tpl.PersistentFlags().StringVarP(&tplFile, "tplFile", "t", "", "指定 模版文件")
	tpl.PersistentFlags().StringVarP(&outFile, "outFile", "o", "", "指定 输出文件名")
	rootCmd.AddCommand(tpl)
}
