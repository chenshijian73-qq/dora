package simplified_cmd

import (
	"bufio"
	"fmt"
	common "github.com/chenshijian73-qq/doraemon/pkg"
	"log"
	"os"
	"os/exec"
	"strings"
)

func KubectlSimpleConfig() {
	// 快捷配置内容写进 ～/.dora_kubectl
	doraKubectlConfig(Kubectl_shortcut)

	dirname, _ := os.UserHomeDir()

	doraPath := fmt.Sprintf("%s/.dora_kubectl", dirname)
	// source 配置文件的命令
	doraSource := fmt.Sprintf("\nsource %s", doraPath)

	out, _ := exec.Command("bash", "-c", "echo $SHELL").CombinedOutput()
	shell := strings.Replace(string(out), "\n", "", -1)

	if shell == "/bin/bash" {
		zshPath := fmt.Sprintf("%s/.bashrc", dirname)
		addConfigToShell(zshPath, doraSource)
	} else if shell == "/bin/zsh" {
		zshPath := fmt.Sprintf("%s/.zshrc", dirname)
		addConfigToShell(zshPath, doraSource)
	} else {
		log.Fatal("(｡ì _ í｡) Sorry，cant not get which shell is using。")
	}
}

func doraKubectlConfig(content string) {
	dirname, err := os.UserHomeDir()
	common.CheckAndExit(err)

	doraPath := fmt.Sprintf("%s/.dora_kubectl", dirname)
	doraFile, err := os.OpenFile(doraPath, os.O_CREATE|os.O_WRONLY, 0777)
	common.CheckAndExit(err)
	defer doraFile.Close()
	writer := bufio.NewWriter(doraFile)
	writer.Write([]byte(content))
	writer.Flush()
}

func addConfigToShell(shellConfigPath, config string) {
	checkSourceCmd := fmt.Sprintf("cat %s |grep '%s'", shellConfigPath, ".dora_kubectl")
	sourceIsExist, _ := exec.Command("bash", "-c", checkSourceCmd).CombinedOutput()
	if string(sourceIsExist) == "" {
		configFile, err := os.OpenFile(shellConfigPath, os.O_WRONLY|os.O_APPEND, 0666)
		defer configFile.Close()
		common.CheckErr(err)
		shellConfigWriter := bufio.NewWriter(configFile)
		shellConfigWriter.Write([]byte(config))
		shellConfigWriter.Flush()

		sourceCmd := fmt.Sprintf("source %s", shellConfigPath)
		cmd := exec.Command("bash", "-c", sourceCmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("combined out:\n%s\n", string(out))
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}
}
