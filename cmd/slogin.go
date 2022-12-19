package cmd

import (
	"fmt"
	"github.com/chenshijian73-qq/doraemon/internal"
	"github.com/spf13/cobra"
	"path/filepath"
)

var slogin = &cobra.Command{
	Use:   "slogin SERVER_NAME",
	Short: "Login server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			internal.SingleLogin(args[0])
		} else {
			_ = cmd.Help()
		}
	},
	ValidArgsFunction: func(_ *cobra.Command, _ []string, _ string) (res []string, _ cobra.ShellCompDirective) {
		for _, s := range internal.ListServers(true) {
			res = append(res, fmt.Sprintf("%s\tfrom %s(%s)", s.Name, filepath.Base(s.ConfigPath), s.Name))
		}
		return res, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(slogin)
}
