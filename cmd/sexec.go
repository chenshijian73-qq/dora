package cmd

import (
	"fmt"
	"github.com/chenshijian73-qq/doraemon/internal/machine"
	"github.com/spf13/cobra"
	"path/filepath"
)

var execGroup bool

var commnad string

var sexec = &cobra.Command{
	Use:   "sexec [OPTIONS] SERVER|TAG -c COMMAND",
	Short: "Batch exec command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			_ = cmd.Help()
		} else {
			machine.Exec(commnad, args[0], execGroup, false)
		}
	},
	ValidArgsFunction: func(_ *cobra.Command, _ []string, _ string) (res []string, _ cobra.ShellCompDirective) {
		for _, s := range machine.ListServers(true) {
			res = append(res, fmt.Sprintf("%s\tfrom %s(%s)", s.Name, filepath.Base(s.ConfigPath), s.Name))
		}
		return res, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	sexec.PersistentFlags().BoolVarP(&execGroup, "tag", "t", false, "server tag")
	sexec.PersistentFlags().StringVarP(&commnad, "command", "c", "", "give the command to execute")
	rootCmd.AddCommand(sexec)
}
