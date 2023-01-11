package cmd

import (
	"github.com/chenshijian73-qq/doraemon/internal/simplified_cmd"
	"github.com/spf13/cobra"
)

var simple_kctl = &cobra.Command{
	Use:   "simple_kctl",
	Short: "Add simplified kubectl commands to the terminal",
	Run: func(cmd *cobra.Command, args []string) {
		simplified_cmd.KubectlSimpleConfig()
	},
}

func init() {
	rootCmd.AddCommand(simple_kctl)
}
