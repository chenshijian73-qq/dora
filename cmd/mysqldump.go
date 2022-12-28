package cmd

import (
	"fmt"
	"github.com/chenshijian73-qq/doraemon/internal/machine"
	"github.com/chenshijian73-qq/doraemon/internal/paas"
	"github.com/spf13/cobra"
)

var host, port, user, passwd, dumpPath string

var mysqldump = &cobra.Command{
	Use:     "mysqldump [Args]",
	Short:   "Dump Mariadb Databases",
	Example: "dora mysqldump -H mysql_host -u mysql -p 123456 -P 3306 db1 db2 db3 -o backup/",
	Run: func(cmd *cobra.Command, args []string) {
		paas.Dump(host, port, user, passwd, dumpPath, args)
	},
	ValidArgsFunction: func(_ *cobra.Command, _ []string, _ string) (res []string, _ cobra.ShellCompDirective) {
		for _, info := range machine.Configs {
			res = append(res, fmt.Sprintf("%s\t%s", info.Name, info.Path))
		}
		return res, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	mysqldump.PersistentFlags().StringVarP(&host, "host", "H", "127.0.0.1", "mysql host")
	mysqldump.PersistentFlags().StringVarP(&port, "port", "P", "3306", "mysql port")
	mysqldump.PersistentFlags().StringVarP(&user, "user", "u", "mysql", "mysql user")
	mysqldump.PersistentFlags().StringVarP(&passwd, "passwd", "p", "mysql123", "mysql passwd")
	mysqldump.PersistentFlags().StringVarP(&dumpPath, "dumpPath", "o", "test/", "dump path")

	rootCmd.AddCommand(mysqldump)
}
