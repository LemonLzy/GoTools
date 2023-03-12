## A collection of gadgets  by golang
##### new项目结构生成工具
---
项目分层遵循[`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2)。

---
##### chatgpt robot工具
将个人微信变为chatgpt对话机器人（请先更新chatgpt/config/config.json中的配置）
```shell
go run main.go chat
```

---
##### rename工具
根据绝对路径，将文件夹下，包含old name的文件名，替换为new name的文件名（new name可以为空）
```shell
go run main.go rename [absolute path] [old name] [new name]
```
示例如下：
```shell
$ go run main.go rename /users/lzy/Downloads/test flora pig
```

---
##### json pretty工具
格式化json
```shell
go run main.go pretty '[unformatted json string]'
```
示例如下：
```shell
$ go run main.go pretty '{"name": "Strawberry", "color": "red"}'
```