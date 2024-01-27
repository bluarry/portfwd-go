# portfwd
端口转发工具cli
# 描述
用golang写的tcp/udp端口转发cli工具
# 用法
```
NAME:
   portfwd - port forward cli tool

USAGE:
   portfwd [global options] command [command options] 

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --type value, -t value  tcp|udp
   --daemon, -d            (default: false)
   --log value, -l value   (default: "./run.log")
   --help, -h              show help
```
# 使用示例
## 映射本地tcp 1022端口转发到 192.168.1.77:22

    portfwd -t tcp 0.0.0.0:1022 192.168.1.77:22     # 所有的ip均可访问
    portfwd -t tcp 127.0.0.1:1022 192.168.1.77:22   # 仅本地可访问
    portfwd -t tcp [::]:1022 192.168.1.77:22        # 1022端口ipv6均可访问

## 映射本地 udp 53端口到  8.8.8.8:53

    portfwd -t udp 0.0.0.0:53 8.8.8.8:53
    portfwd -t udp [::]:53 8.8.8.8:53

## IPv4-IPv6 相互转换

    portfwd -t udp [::]:1701 localhost:1701         # 增加本地 L2TP 服务的ipv6支持
    portfwd -t tcp 0.0.0.0:80 [2409:8c1e:75b0:1121::15]:80    # 增加本地80端口转发到ipv6网址

## 后台运行端口转发并设置日志路径
    portfwd -t tcp -d -l ./log/run.log 0.0.0.0:80 [2409:8c1e:75b0:1121::15]:80
