package shortcut

import (
	"bufio"
	"fmt"
	common "github.com/chenshijian73-qq/doraemon/pkg"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func Test_shortcut(t *testing.T) {
	shortcut := `
alias k='kubectl'
alias kgp='kubectl get pod'
alias kdc='kubectl describe'
alias kdl='kubectl delete '
alias ke='kubectl edit'
alias kgnode='kubectl get nodes'
alias ktnode='kubectl top nodes'
alias anoready='kubectl get pod -A|grep 0/|grep -v Complete'

function noready(){
    local OPTIND 
    while getopts 'n:k:h:' OPT
    do
        case $OPT in
            n)
            namespace=$OPTARG
            ;;
            h)
            echo 'examplel: noready -n ns1'
	return 0
	;;
	?)
	echo "(^_^)v ERROR: 请输入参数 -n"
	return 1
	;;
	esac
	done
	kubectl get pod -n $namespace |grep 0/|grep -v Complete
}

function podimage(){
    local OPTIND
    while getopts 'n:k:h:' OPT
    do
        case $OPT in
            n)
            namespace=$OPTARG
            ;;
            k)
            keyword=$OPTARG
            ;;
            h)
            echo 'examplel: podimage -n ns1 -k name1'
	return 0
	;;
	?)
	echo "(*ˉ︶ˉ*) ERROR: 请输入参数 -n 和 -k"
	return 1
	;;
	esac
	done
	kubectl get pod -n $namespace|grep $keyword|awk '{print$1}'|xargs -I {} kubectl get pod -n $namespace {} -o jsonpath="{.spec.containers[*].image} "|tr -s '[[:space:]]' '\n'|sort|uniq
}

function delpods(){
    local OPTIND
    while getopts 'n:k:h:' OPT; do
        case $OPT in
            n)
            namespace=$OPTARG
            ;;
            k)
            keyword=$OPTARG
            ;;
            h)
            echo 'examplel: delpods -n ns1 -k name1'
	return 0
	;;
	?)
	echo "(*ˉ︶ˉ*) ERROR: 请输入参数 -n 和 -k"
	return 1
	;;
	esac
	done
	kubectl get pod -n $namespace|grep $keyword|awk '{print$1}'|xargs kubectl -n $namespace delete pod
}

function kedit(){
    local OPTIND
    while getopts 'n:k:t:h:' OPT; do
        case $OPT in
            n)
            namespace=$OPTARG
            ;;
            k)
            keyword=$OPTARG
            ;;
            t)
            type=$OPTARG
            ;;
            h)
            echo  'examplel: kedit -n ns1 -t cm -k name1'
	return 0
	;;
	?)
	echo "(*ˉ︶ˉ*) ERROR: 请输入参数 -n, -t 和 -k"
	return 1
	;;
	esac
	done
	array=($(kubectl get -n $namespace $type |grep $keyword|awk '{print$1}'))

	num=${#array[@]}
	
	if [ "$num" == '1' ]
	then
	name=$array
	else
	echo "Please select the one to operate"
	select var in ${array[*]}
	do
	case $var in
	*)
	name=$var
	break
	;;
	esac
	done
	fi
	
	kubectl -n $namespace edit $type $name
}
function klogs(){
    local OPTIND
    while getopts 'n:k:h:' OPT; do
        case $OPT in
            n)
            namespace=$OPTARG
            ;;
            k)
            keyword=$OPTARG
            ;;
            h)
            echo 'examplel: klogs -n ns1 -k pod1'
	return 0
	;;
	?)
	echo "(*ˉ︶ˉ*) ERROR: 请输入参数 -n , -k"
	return 1
	;;
	esac
	done

	array=($(kubectl get -n $namespace pod |grep $keyword|awk '{print$1}'))

	num=${#array[@]}

    if [ "$num" == '1' ]
    then
        name=$array
    else
        echo "Please select the one to operate"
        select var in ${array[*]}
        do
            case $var in
                *)
                    name=$var
                    break
                ;;
            esac
        done
    fi

    kubectl -n $namespace logs -f $name --tail=1000
}
function kexec(){
    local OPTIND
    while getopts 'n:k:s:h:' OPT; do
        case $OPT in
            n)
            namespace=$OPTARG
            ;;
            k)
            keyword=$OPTARG
            ;;
            s)
            cmd=$OPTARG
            ;;
            h)
            echo 'examplel: kexec -n ns1 -k name1 -s bash'
	return 0
	;;
	?)
	echo "(*ˉ︶ˉ*) ERROR: 请输入参数 -n , -k, -s"
	return 1
	;;
	esac
	done

	array=($(kubectl get -n $namespace pod |grep $keyword|awk '{print$1}'))

	num=${#array[@]}

    if [ "$num" == '1' ]
    then
        name=$array
    else
        echo "Please select the one to operate"
        select var in ${array[*]}
        do
            case $var in
                *)
                    name=$var
                    break
                ;;
            esac
        done
    fi

    kubectl -n $namespace exec -it $name $cmd
}
function kg(){
    local OPTIND
    while getopts 't:k:h:' OPT
    do
        case $OPT in
            t)
            type=$OPTARG
            ;;
            k)
            key=$OPTARG
            ;;
            h)
            echo 
	return 0
	;;
	\?)
	echo "(^_^)v ERROR: 请输入参数 -n" 'examplel: kg -t svc -k keyword'
	return 1
	;;
	esac
	done
	kubectl get $type -A|grep $key
}
`
	// 快捷配置内容写进 ～/.dora
	doraConfig(shortcut)

	dirname, _ := os.UserHomeDir()

	doraPath := fmt.Sprintf("%s/.dora", dirname)
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

func doraConfig(content string) {
	dirname, err := os.UserHomeDir()
	common.CheckAndExit(err)

	doraPath := fmt.Sprintf("%s/.dora", dirname)
	doraFile, err := os.OpenFile(doraPath, os.O_CREATE|os.O_WRONLY, 0777)
	common.CheckAndExit(err)
	defer doraFile.Close()
	writer := bufio.NewWriter(doraFile)
	writer.Write([]byte(content))
	writer.Flush()
}

func addConfigToShell(shellConfigPath, config string) {
	checkSourceCmd := fmt.Sprintf("cat %s |grep '%s'", shellConfigPath, ".dora")
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
