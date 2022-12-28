package cmd

import (
	"fmt"
	"github.com/chenshijian73-qq/doraemon/internal/machine"
	"github.com/chenshijian73-qq/doraemon/internal/paas"
	"github.com/spf13/cobra"
)

var pg_host, pg_user, pg_passwd, dbName, pg_dumpPath string
var pg_port int16

var pgdump = &cobra.Command{
	Use:     "pgdump [Args]",
	Short:   "Dump PostgreSql Database",
	Example: "dora pgdump -H pg_host -u postgres -p 123456 -P 5432 -d db1 -o backup/",
	Run: func(cmd *cobra.Command, args []string) {
		paas.PgDump(pg_host, dbName, pg_user, pg_passwd, pg_dumpPath, int(pg_port))
	},
	ValidArgsFunction: func(_ *cobra.Command, _ []string, _ string) (res []string, _ cobra.ShellCompDirective) {
		for _, info := range machine.Configs {
			res = append(res, fmt.Sprintf("%s\t%s", info.Name, info.Path))
		}
		return res, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	pgdump.PersistentFlags().StringVarP(&pg_host, "host", "H", "127.0.0.1", "pg host")
	pgdump.PersistentFlags().Int16VarP(&pg_port, "port", "P", 5432, "pg port")
	pgdump.PersistentFlags().StringVarP(&pg_user, "user", "u", "postgres", "pg user")
	pgdump.PersistentFlags().StringVarP(&pg_passwd, "passwd", "p", "123456", "pg passwd")
	pgdump.PersistentFlags().StringVarP(&dbName, "dbName", "d", "postgres", "pg dbName")
	pgdump.PersistentFlags().StringVarP(&pg_dumpPath, "dumpPath", "o", "test/", "dump path")

	rootCmd.AddCommand(pgdump)
}
