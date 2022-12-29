package paas

import (
	"bufio"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"strings"
	"testing"
)

func Test_redisDump(t *testing.T) {
	// 连接到 Redis 服务器
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 循环读取用户输入的命令并执行
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("redis> ")
		if !scanner.Scan() {
			break
		}
		fmt.Println("hello")
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Println("line: ", line)
		fields := strings.Fields(line)
		if len(fields) == 0 {
			fmt.Println("Invalid command")
			continue
		}
		cmd := fields[0]
		var args []interface{}
		if len(fields) > 1 {
			args = make([]interface{}, len(fields)-1)
			for i, field := range fields[1:] {
				args[i] = field
			}
		}

		// 执行命令
		result, err := conn.Do(cmd, redis.Args{}.AddFlat(args)...)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 输出命令的结果
		fmt.Println(result)
	}
}
func Test_scan(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
