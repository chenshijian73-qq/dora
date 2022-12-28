package paas

import (
	"database/sql"
	"errors"
	"fmt"
	common "github.com/chenshijian73-qq/doraemon/pkg"
	"github.com/go-sql-driver/mysql"
	"github.com/jamf/go-mysqldump"
	"log"
	"os"
	"path"
	"time"
)

type DumpConfig struct {
	MysqlConfig *mysql.Config
	DumpPath    string
}

func Dump(host, port, user, passwd, dumppath string, dbGroup []string) {

	mysqlconfig := mysql.NewConfig()
	mysqlconfig.User = user
	mysqlconfig.Passwd = passwd
	mysqlconfig.DBName = ""
	mysqlconfig.Net = "tcp"
	mysqlconfig.Addr = "127.0.0.1:13306"
	mysqlconfig.Addr = fmt.Sprint(host, ":", port)
	mysqlconfig.ParseTime = true

	var dumpconfig DumpConfig

	if len(dbGroup) < 1 {
		log.Fatal("dont config dbName")
	} else if len(dbGroup) == 1 {
		mysqlconfig.DBName = dbGroup[0]
		dumpconfig = DumpConfig{
			MysqlConfig: mysqlconfig,
			DumpPath:    dumppath,
		}
		Mysqldump(dumpconfig)
	} else {
		for i := 0; i < len(dbGroup); i++ {
			mysqlconfig.DBName = dbGroup[i]
			dumpconfig = DumpConfig{
				MysqlConfig: mysqlconfig,
				DumpPath:    dumppath,
			}
			Mysqldump(dumpconfig)
		}
	}
}

func Mysqldump(config DumpConfig) {

	dumpFilenameFormat := fmt.Sprintf("%s-20060102T150405", config.MysqlConfig.DBName)
	fmt.Println(config.MysqlConfig.FormatDSN())
	db, err := sql.Open("mysql", config.MysqlConfig.FormatDSN())
	common.PrintErrWithPrefixAndExit("Error opening database: ", err)

	err = db.Ping()
	common.PrintErrWithPrefixAndExit("Failed to ping mysql: ", err)

	dumper, err, filePath := Register(db, config.DumpPath, dumpFilenameFormat)
	common.PrintErrWithPrefixAndExit("Error registering databse:", err)

	err = dumper.Dump()
	common.PrintErrWithPrefixAndExit("Error dumping:", err)

	err = dumper.Close()
	common.PrintErrWithPrefixAndExit("dump close:", err)

	fmt.Printf("File is saved to %s\n", filePath)
}

func Register(db *sql.DB, dir, format string) (*mysqldump.Data, error, string) {
	if !IsDir(dir) {
		return nil, errors.New("invalid directory"), ""
	}

	name := time.Now().Format(format)
	p := path.Join(dir, name+".sql")

	// Check dump directory
	if e, _ := Exists(p); e {
		return nil, errors.New("Dump '" + name + "' already exists."), ""
	}

	// Create .sql file
	f, err := os.Create(p)

	if err != nil {
		return nil, err, ""
	}

	return &mysqldump.Data{
		Out:        f,
		Connection: db,
	}, nil, p
}

func Exists(p string) (bool, os.FileInfo) {
	f, err := os.Open(p)
	if err != nil {
		return false, nil
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return false, nil
	}
	return true, fi
}

func IsDir(p string) bool {
	if e, fi := Exists(p); e {
		return fi.Mode().IsDir()
	}
	return false
}
