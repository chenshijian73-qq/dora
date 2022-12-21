package internal

import (
	"database/sql"
	"errors"
	"fmt"
	common "github.com/chenshijian73-qq/doraemon/pkg"
	"github.com/go-sql-driver/mysql"
	"github.com/jamf/go-mysqldump"
	"os"
	"path"
	"time"
)

type DumpConfig struct {
	MysqlConfig *mysql.Config
	DumpPath    string
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

	fmt.Printf("File is saved to %s", filePath)
}

func Register(db *sql.DB, dir, format string) (*mysqldump.Data, error, string) {
	if !isDir(dir) {
		return nil, errors.New("invalid directory"), ""
	}

	name := time.Now().Format(format)
	p := path.Join(dir, name+".sql")

	// Check dump directory
	if e, _ := exists(p); e {
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

func exists(p string) (bool, os.FileInfo) {
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

func isDir(p string) bool {
	if e, fi := exists(p); e {
		return fi.Mode().IsDir()
	}
	return false
}
