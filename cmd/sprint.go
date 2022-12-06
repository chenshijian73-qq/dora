package cmd

import (
	"fmt"
	"github.com/chenshijian73-qq/Doraemon/internal"
	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	serverSort bool
	serverName string
)
var sprint = &cobra.Command{
	Use:   "sprint [SERVER_NAME]",
	Short: "Print server list",
	Run: func(cmd *cobra.Command, args []string) {
		if serverName != "" {
			internal.PrintServerDetail(serverName)
		} else {
			internal.PrintServers(serverSort)
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
	sprint.PersistentFlags().BoolVarP(&serverSort, "sort", "S", false, "sort server list")
	sprint.PersistentFlags().StringVarP(&serverName, "server_name", "s", "", "print server detail")
	rootCmd.AddCommand(sprint)
}
