package cmd

import (
	"fmt"
	"github.com/chenshijian73-qq/doraemon/internal/machine"
	"github.com/chenshijian73-qq/doraemon/internal/paas"
	"github.com/spf13/cobra"
)

var redis_host, redis_port string

var rediscli = &cobra.Command{
	Use:     "rediscli [Args]",
	Short:   "redis client",
	Example: "dora rediscli -H redis_host -P 6379",
	Run: func(cmd *cobra.Command, args []string) {
		paas.RedisCli(redis_host, redis_port)
	},
	ValidArgsFunction: func(_ *cobra.Command, _ []string, _ string) (res []string, _ cobra.ShellCompDirective) {
		for _, info := range machine.Configs {
			res = append(res, fmt.Sprintf("%s\t%s", info.Name, info.Path))
		}
		return res, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rediscli.PersistentFlags().StringVarP(&redis_host, "host", "H", "127.0.0.1", "pg host")
	rediscli.PersistentFlags().StringVarP(&redis_port, "port", "P", "6379", "pg port")

	rootCmd.AddCommand(rediscli)
}
