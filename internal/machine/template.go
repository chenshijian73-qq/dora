package machine

var Tpl_File = `{{- $root := . -}}
servers:
{{- range .servers.host }}
# 服务器名称
- name: {{ $root.servers.namePrefix }}{{ (split "." .)._3 }}
  # 服务器地址
  address: {{ . }}
  # 服务器端口
  port: {{ $root.servers.port }}
  # 服务器 tag 列表
  tags:
  {{- range $root.servers.tag }}
  - {{ . }}
  {{- end }}
   # 登录用户
  user: {{ $root.servers.user }}
  # 登录密码
  password: {{ $root.servers.password }}
  # 登录私钥
  private_key: ""
  # 如果私钥需要密码则在此填写私钥密码
  private_key_password: ""
  # 连接此服务器需要预先连接的服务器(跳板机)
  proxy:
  # 每隔 n 时间发送心跳包保证 ssh 不会自动断开
  server_alive_interval: 20s
  ########### 以下部分为高级配置，具体请阅读后面的高级应用章节 ###########
  # hook_cmd 用于在登录后 hook 远端输入，以实现自动化
  hook_cmd:
  # 配合 hook_cmd 可以读取远端输出
  hook_stdout: true
  # ssh 认证的键盘挑战 hook 脚本，用户可自行扩展实现键盘挑战登录
  keyboard_auth_cmd:
  # ssh 连接成功后注入 session 环境变量(需要 server sshd 调整配置)
  environment:
    ENABLE_VIM_CONFIG: "true"
    Other_KEY: "Other_String_Value"
  # 开启本地 API 支持，开启后可通过远端调用本地 api
  enable_api: true
{{- end }}
`

var Data_file = `servers:
  namePrefix: dev-
  host:
    - 1.1.1.1
    - 2.2.2.2
  port: 36000
  tag:
    - dev
  user: root
  password: C0ding@2022!
`
