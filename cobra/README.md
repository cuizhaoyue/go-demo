# cobra练习

1. 安装

安装 `cobra` 包
```bash
go get -u github.com/spf13/cobra/cobra@latest
```

安装 `cobra-cli` 命令行工具
```shell
go install github.com/spf13/cobra-cli@latest
```

2. 创建demo

```bash
mkdir -p cobra
cd cobra
cobra-cli init --viper demo // TODO: 使用cobra-cli命令初始化后main.go文件中的import引入的路径有误
```

3. 创建命令

```bash
cd demo
cobra-cli add config
cobra-cli add serve
```

4. 编译运行

```bash
go build -v .
./demoapp config
./demoapp serve
```


