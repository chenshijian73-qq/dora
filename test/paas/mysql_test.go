package paas

import (
	"database/sql"
	"fmt"
	"github.com/chenshijian73-qq/doraemon/internal/paas"
	common "github.com/chenshijian73-qq/doraemon/pkg"
	"github.com/go-sql-driver/mysql"
	"github.com/jamf/go-mysqldump"
	"os"
	"testing"
)

func Test_mysqldump1(t *testing.T) {
	mysqlconfig := mysql.NewConfig()
	mysqlconfig.User = "coding"
	mysqlconfig.Passwd = "coding123"
	mysqlconfig.DBName = "coding_testing"
	mysqlconfig.Net = "tcp"
	mysqlconfig.Addr = "127.0.0.1:13306"
	mysqlconfig.ParseTime = true
	wd, _ := os.Getwd()

	dumpPath := fmt.Sprintf("%s/%s", wd, "sql")
	config := paas.DumpConfig{
		MysqlConfig: mysqlconfig,
		DumpPath:    dumpPath,
	}

	dumpFilenameFormat := fmt.Sprintf("%s-20060102T150405", config.MysqlConfig.DBName)
	db, err := sql.Open("mysql", config.MysqlConfig.FormatDSN())
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return
	}
	dumper, err := mysqldump.Register(db, config.DumpPath, dumpFilenameFormat)
	if err != nil {
		fmt.Println("Error registering databse:", err)
		return
	}
	// Dump database to file
	err = dumper.Dump()
	if err != nil {
		fmt.Println("Error dumping:", err)
		return
	}
	fmt.Printf("File is saved to %s", dumpFilenameFormat)

	// Close dumper, connected database and file stream.

	common.CheckErr(dumper.Close())
}

func Test_mysqldump2(t *testing.T) {
	DBGroup := [...]string{"test", "workflow"}
	mysqlconfig := mysql.NewConfig()
	mysqlconfig.User = "coding"
	mysqlconfig.Passwd = "coding123"
	mysqlconfig.DBName = ""
	mysqlconfig.Net = "tcp"
	mysqlconfig.Addr = "127.0.0.1:13306"
	mysqlconfig.ParseTime = true

	wd, _ := os.Getwd()
	dumpPath := fmt.Sprintf("%s/%s", wd, "sql")

	for i := 0; i < len(DBGroup); i++ {
		mysqlconfig.DBName = DBGroup[i]
		config := paas.DumpConfig{
			MysqlConfig: mysqlconfig,
			DumpPath:    dumpPath,
		}
		paas.Mysqldump(config)
	}
}

func Test_mysqldump3(t *testing.T) {
	wd, _ := os.Getwd()
	dumpPath := fmt.Sprintf("%s/%s", wd, "sql")
	DBGroup := []string{"test", "workflow"}
	paas.Dump("127.0.0.1", "13306", "coding", "coding123", dumpPath, DBGroup)
}
