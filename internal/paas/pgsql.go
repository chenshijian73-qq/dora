package paas

import (
	"database/sql"
	"errors"
	"fmt"
	common "github.com/chenshijian73-qq/doraemon/pkg"
	_ "github.com/lib/pq"
	"os"
	"path"
	"time"
)

type PgDumpConfig struct {
	Host     string
	Port     int
	DB       string
	Username string
	Password string
	Outpath  string
}

func PgDump(host, db, user, passwd, outpath string, port int) {
	config := PgDumpConfig{
		Host:     host,
		Port:     port,
		DB:       db,
		Username: user,
		Password: passwd,
		Outpath:  outpath,
	}
	PgDumpExec(&config)
}

func PgDumpExec(config *PgDumpConfig) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.Username, config.Password, config.DB)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Printf("Error connecting to database: %s\n", err)
		return
	}
	defer db.Close()

	// Open a file for writing the dump
	filePath, err := IsDirVaild(config.Outpath, config.DB+"-"+"20060102T150405")
	common.CheckAndExit(err)
	dumpFile, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating dump file: %s\n", err)
		return
	}
	defer dumpFile.Close()

	// Get a list of all tables in the database
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	if err != nil {
		fmt.Printf("Error getting table list: %s\n", err)
		return
	}
	defer rows.Close()

	// Iterate through each table and dump its data
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			fmt.Printf("Error scanning table name: %s\n", err)
			return
		}

		// Dump the data for the table
		dataRows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
		if err != nil {
			fmt.Printf("Error getting data for table %s: %s\n", tableName, err)
			return
		}
		defer dataRows.Close()

		// Get the column names for the table
		columnNames, err := dataRows.Columns()
		if err != nil {
			fmt.Printf("Error getting column names for table %s: %s\n", tableName, err)
			return
		}

		// Write the CREATE TABLE statement to the dump file
		_, err = fmt.Fprintf(dumpFile, "CREATE TABLE %s (\n", tableName)
		if err != nil {
			fmt.Printf("Error writing to dump file: %s\n", err)
			return
		}

		// Write the column names to the CREATE TABLE statement
		for i, columnName := range columnNames {
			_, err = fmt.Fprintf(dumpFile, "  %s", columnName)
			if err != nil {
				fmt.Printf("Error writing to dump file: %s\n", err)
				return
			}
			if i < len(columnNames)-1 {
				_, err = fmt.Fprintf(dumpFile, ",\n")
			} else {
				_, err = fmt.Fprintf(dumpFile, "\n")
			}
			if err != nil {
				fmt.Printf("Error writing to dump file: %s\n", err)
				return
			}
		}
		_, err = fmt.Fprintf(dumpFile, ");\n\n")
		if err != nil {
			fmt.Printf("Error writing to dump file: %s\n", err)
			return
		}

		// Write the INSERT statements for the data
		for dataRows.Next() {
			// Scan the data into a slice of pointers to the values
			values := make([]interface{}, len(columnNames))
			valuePtrs := make([]interface{}, len(columnNames))
			for i := range values {
				valuePtrs[i] = &values[i]
			}
			if err := dataRows.Scan(valuePtrs...); err != nil {
				fmt.Printf("Error scanning data row: %s\n", err)
				return
			}

			// Write the INSERT statement for the row
			_, err = fmt.Fprintf(dumpFile, "INSERT INTO %s (", tableName)
			if err != nil {
				fmt.Printf("Error writing to dump file: %s\n", err)
				return
			}
			for i, columnName := range columnNames {
				_, err = fmt.Fprintf(dumpFile, "%s", columnName)
				if err != nil {
					fmt.Printf("Error writing to dump file: %s\n", err)
					return
				}
				if i < len(columnNames)-1 {
					_, err = fmt.Fprintf(dumpFile, ", ")
				}
				if err != nil {
					fmt.Printf("Error writing to dump file: %s\n", err)
					return
				}
			}
			_, err = fmt.Fprintf(dumpFile, ") VALUES (")
			if err != nil {
				fmt.Printf("Error writing to dump file: %s\n", err)
				return
			}
			for i, value := range values {
				// Convert the value to a string and write it to the INSERT statement
				_, err = fmt.Fprintf(dumpFile, "%v", value)
				if err != nil {
					fmt.Printf("Error writing to dump file: %s\n", err)
					return
				}
				if i < len(values)-1 {
					_, err = fmt.Fprintf(dumpFile, ", ")
					if err != nil {
						fmt.Printf("Error writing to dump file: %s\n", err)
						return
					}
				}
			}
			_, err = fmt.Fprintf(dumpFile, ");\n")
			if err != nil {
				fmt.Printf("Error writing to dump file: %s\n", err)
				return
			}
		}
		_, err = fmt.Fprintf(dumpFile, "\n")
		if err != nil {
			fmt.Printf("Error writing to dump file: %s\n", err)
			return
		}
	}

	fmt.Printf("Successfully dumped database to %s\n", dumpFile.Name())
}

func IsDirVaild(dir, format string) (string, error) {
	if !IsDir(dir) {
		return "", errors.New("invalid directory")
	}

	name := time.Now().Format(format)
	p := path.Join(dir, name+".sql")

	// Check dump directory
	if e, _ := Exists(p); e {
		return "", errors.New("Dump '" + name + "' already exists.")
	}

	return p, nil
}
