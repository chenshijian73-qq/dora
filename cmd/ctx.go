package cmd

import (
	"fmt"
	"github.com/chenshijian73-qq/doraemon/internal"
	"github.com/spf13/cobra"
)

var ctx = &cobra.Command{
	Use:   "ctx [Config File]",
	Short: "Change config file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			internal.ListConfigs()
			return
		}

		internal.SetConfig(args[0])
		fmt.Printf("👉 changed config to [%s.yaml]...\n", args[0])
	},
	ValidArgsFunction: func(_ *cobra.Command, _ []string, _ string) (res []string, _ cobra.ShellCompDirective) {
		for _, info := range internal.Configs {
			res = append(res, fmt.Sprintf("%s\t%s", info.Name, info.Path))
		}
		return res, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(ctx)
}
