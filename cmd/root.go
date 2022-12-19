package cmd

import (
	"github.com/chenshijian73-qq/doraemon/internal"
	common "github.com/chenshijian73-qq/doraemon/pkg"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	completionShell string
	showVersion     bool
)

var rootCmd = &cobra.Command{
	Use:   "dora",
	Short: "Four dimensional pocket",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	internal.LoadConfig()

	if os.Args[0] != rootCmd.Name() {
		subCmd, _, err := rootCmd.Find([]string{filepath.Base(os.Args[0])})
		if err == nil && subCmd.Name() != rootCmd.Name() {
			// if find a subcommand, we need to remove the subcommand from the parent command
			// to ensure that the '__complete' command takes effect
			rootCmd.RemoveCommand(subCmd)
			// reset args for the subcommand
			if len(os.Args) > 1 {
				subCmd.SetArgs(os.Args[1:])
			}
			// execute subcommand
			common.CheckAndExit(subCmd.Execute())
			return
		}
	}
	common.CheckAndExit(rootCmd.Execute())
}

func init() {
}
