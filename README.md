goctl api new core
go run core.go -f etc/core-api.yaml
可以通过.api文件一键快速生成一个api服务，
goctl api go -api core.api -dir . -style go_zero
用户注册
1 看是否存在 , 不存在生成新的
2 返回token