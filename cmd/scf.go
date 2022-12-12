package cmd

import (
	"github.com/chenshijian73-qq/Doraemon/internal"
	"github.com/spf13/cobra"
)

var copy2Group bool

var scf = &cobra.Command{
	Use:   "scf [-t] FILE/DIR|SERVER:PATH SERVER:PATH|FILE/DIR",
	Short: "Copies files between hosts on a network",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			_ = cmd.Help()
		} else {
			internal.Copy(args, copy2Group)
		}
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) (res []string, _ cobra.ShellCompDirective) {
		for _, s := range internal.ListServers(true) {
			res = append(res, s.Name)
		}
		return res, cobra.ShellCompDirectiveDefault
	},
}

func init() {
	scf.PersistentFlags().BoolVarP(&copy2Group, "tag", "t", false, "server tag")
	rootCmd.AddCommand(scf)
}
