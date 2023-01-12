package simplified_cmd

var Kubectl_shortcut = `
alias k='kubectl'
alias kgp='kubectl get pod'
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
            printYello 'example:  noready -n ns1'
	return 0
	;;
	?)
	printRed "(^_^)v ERROR: 请输入参数 -n"
    printYello 'example:  noready -n ns1'
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
            printYello 'example:  podimage -n ns1 -k name1'
	return 0
	;;
	?)
	printRed "(*ˉ︶ˉ*) ERROR: 请输入参数 -n 和 -k"
    printYello 'example:  podimage -n ns1 -k name1'
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
            printYello 'example:  delpods -n ns1 -k name1'
	return 0
	;;
	?)
	printRed "(*ˉ︶ˉ*) ERROR: 请输入参数 -n 和 -k"
    printYello 'example:  delpods -n ns1 -k name1'
	return 1
	;;
	esac
	done
	kubectl get pod -n $namespace|grep $keyword|awk '{print$1}'|xargs kubectl -n $namespace delete pod
}

select_object=""
function selectWhichOneToContine()
{
    local namespace=$1
    local type=$2
    local keyword=$3
    array=($(kubectl get -n $namespace $type |grep $keyword|awk '{print$1}'))

    if [ "$array" == "" ];then
        printRed "can't find anything with keyword you provide"
        return 1
    fi

    local num=${#array[@]}

    if [ "$num" == '1' ]
	then
	select_object=$array
	else
	printYello "Please select the one to operate"
	select var in ${array[*]}
	do
	case $var in
	*)
	select_object=$var
	break
	;;
	esac
	done
	fi    
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
            printYello 'example: kedit -n ns1 -t cm -k name1'
	return 0
	;;
	?)
	printRed "(*ˉ︶ˉ*) ERROR: 请输入参数 -n, -t 和 -k"
    printYello 'example: kedit -n ns1 -t cm -k name1'
	return 1
	;;
	esac
	done
    
	selectWhichOneToContine $namespace $type $keyword

    name=$select_object

	kubectl -n $namespace edit $type $name
}

function kdc(){
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
            printYello 'example: kdc -n ns1 -t cm -k name1'
	return 0
	;;
	?)
	printRed "(*ˉ︶ˉ*) ERROR: 请输入参数 -n, -t 和 -k"
    printYello 'example: kdc -n ns1 -t cm -k name1'
	return 1
	;;
	esac
	done
	selectWhichOneToContine $namespace $type $keyword

    name=$select_object
	kubectl -n $namespace describe $type $name
}

function kdl(){
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
            printYello 'example: kdl -n ns1 -t cm -k name1'
	return 0
	;;
	?)
	printRed "(*ˉ︶ˉ*) ERROR: 请输入参数 -n, -t 和 -k"
    printYello 'example: kdl -n ns1 -t cm -k name1'
	return 1
	;;
	esac
	done

    selectWhichOneToContine $namespace $type $keyword

    name=$select_object

	kubectl -n $namespace delete $type $name
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
            printYello 'example:  klogs -n ns1 -k pod1'
	return 0
	;;
	?)
	printRed "(*ˉ︶ˉ*) ERROR: 请输入参数 -n , -k"
    printYello 'example:  klogs -n ns1 -k pod1'
	return 1
	;;
	esac
	done

	selectWhichOneToContine $namespace "pod" $keyword

    name=$select_object

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
            printYello 'example:  kexec -n ns1 -k name1 -s bash'
	return 0
	;;
	?)
	printRed "(*ˉ︶ˉ*) ERROR: 请输入参数 -n , -k, -s"
    printYello 'example:  kexec -n ns1 -k name1 -s bash'
	return 1
	;;
	esac
	done

	selectWhichOneToContine $namespace "pod" $keyword

    name=$select_object

    container_arrays=($(kubectl get -n $namespace pod $name -o jsonpath='{range .spec.containers[*]}{.name}{"\n"}{end}'))
    container_num=${#array[@]}

    if [ "$container_num" == '1' ]
    then
        name=$container_arrays
    else
        printYello "Please select the one to operate"
        select var in ${container_arrays[*]}
        do
            case $var in
                *)
                    container=$var
                    break
                ;;
            esac
        done
    fi
    

    kubectl -n $namespace exec -it $name -c $container $cmd
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
	printRed "(^_^)v ERROR: 请输入参数 -n" 
    printYello 'example: kg -t svc -k keyword'
	return 1
	;;
	esac
	done
	kubectl get $type -A|grep $key
}

function printYello(){
    echo -e "\033[33m$1\033[0m"
}

function printRed(){
    echo -e "\033[31m$1\033[0m"
}

function k_help(){
    echo -e "\033[33mk           \033[0m" "= kubectl"
    echo -e "\033[33mkgp         \033[0m" "= kubectl get pod "
    echo -e "\033[33mkgnode      \033[0m" "= kubectl get node"
    echo -e "\033[33mktnode      \033[0m" "= kubectl top node"
    echo -e "\033[33manoready    \033[0m" "= kubectl get pod -A|grep 0/|grep -v Complete"
    echo ""
    echo -e "\033[33mnoready     \033[0m" "获取异常 pod, 使用方式: noready -n ns1"
    echo -e "\033[33mdelpods     \033[0m" "删除包含关键字的 pod, 使用方式: delpods -n ns1 -k keyword"
    echo -e "\033[33mpodimage    \033[0m" "获取包含关键字的 pod 镜像, 使用方式:  podimage -n ns1 -k keyword"
    echo -e "\033[33mkdl         \033[0m" "通过关键词筛选, 选择对象删除, 使用方式: kdl -n ns1 -t cm -k name1"
    echo -e "\033[33mklogs       \033[0m" "通过关键词筛选后, 选择对象打印日志, 使用方式: klogs -n ns1 -k keyword"
    echo -e "\033[33mkg          \033[0m" "通过关键词筛选后, 过滤出对象, 使用方式: example: kg -t svc -k keyword"
    echo -e "\033[33mkedit       \033[0m" "通过关键词筛选后, 选择编辑对象, 使用方式: kedit -n ns1 -t cm -k keyword"
    echo -e "\033[33mkexec       \033[0m" "通过关键词筛选后, 选择对象进行容器, 使用方式: kexec -n ns1 -k keyword -s bash"
    echo -e "\033[33mkdc         \033[0m" "通过关键词筛选后，选择要描述的对象, 使用方式: kdc -n ns1 -t cm -k name1"
}
`
