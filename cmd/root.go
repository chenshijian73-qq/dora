package cmd

import (
	common "github.com/chenshijian73-qq/Doraemon/pkg"
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
		if completionShell != "" {
			GenCompletion(cmd, completionShell)
			return
		}
		_ = cmd.Help()
	},
}

func GenCompletion(cmd *cobra.Command, shell string) {
	switch shell {
	case "bash":
		_ = cmd.GenBashCompletion(os.Stdout)
	case "zsh":
		_ = cmd.GenZshCompletion(os.Stdout)
	case "fish":
		_ = cmd.GenFishCompletion(os.Stdout, true)
	case "powershell":
		_ = cmd.GenPowerShellCompletionWithDesc(os.Stdout)
	}
}

func Execute() {
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
	rootCmd.PersistentFlags().StringVar(&completionShell, "completion", "", "generate shell completion")
}
