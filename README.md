## A tools by golang
---
项目分层遵循[`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2)。

---
### tools包
> A collection of gadgets 

##### rename工具
根据绝对路径，将文件夹下，包含old name的文件名，替换为new name的文件名（new name可以为空）
```shell
go run main.go tools [absolute path] [old name] [new name]
```