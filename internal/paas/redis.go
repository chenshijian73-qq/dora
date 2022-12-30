package paas

import (
	"bufio"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"strings"
)

func RedisCli(host, port string) {
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := redis.Dial("tcp", addr)
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
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			fmt.Println("Invalid command")
			continue
		}
		cmd, args := fields[0], fields[1:]
		var cmdArgs []interface{}
		for _, arg := range args {
			cmdArgs = append(cmdArgs, arg)
		}
		result, err := conn.Do(cmd, cmdArgs...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if result == nil {
			fmt.Println("Key not found")
		} else {
			switch v := result.(type) {
			case string:
				fmt.Println(v)
			case []byte:
				fmt.Println(string(v))
				// ...
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			break
		}
	}
}
