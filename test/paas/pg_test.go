package paas

import (
	"github.com/chenshijian73-qq/doraemon/internal/paas"
	_ "github.com/lib/pq"
	"testing"
)

func Test_pgdump(t *testing.T) {
	paas.PgDump("127.0.0.1", "cci", "postgres", "coding123", "sql", 5432)
	//fmt.Println("Backup created successfully")
}
