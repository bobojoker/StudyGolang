package main

import "fmt"

//1. 在终端中打开
//2. 输入 go build ，在目录下会生成一个可执行文件（.exe）。
//3. 运行： 输入helloword.exe就可以运行了
//4. 其他路径下发布，刚刚的我们使用的是将源码路径在终端打开，如果我们随便打开一个终端路径那么，在go build时候就要输入路径。该路径应该从 src路径开始
//这里的src下面的路径是使用'/',而不是'\\'
//5. 自命名发布的可执行文件："go build -o myname.exe"
//6. go run main.go 可以直接运行*.go文件
//7. "go install"可以生成可执行文件*.exe,然后将文件复制到外面的bin目录下面。生成后，可以直接运行可执行文件名，去运行。即数据HelloWorld就可以运行
//8. 编译其他平台
// SET CGO ENABLED=0 //禁用CGO
// SET GOOS=linux //目标平台是linux
// SET GOARCH=amd64 //目标处理器框架是amd64
func main() {
	fmt.Println("Hello World")
}
