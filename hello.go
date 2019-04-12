package main
//定义了包名,每个GO程序都包含一个名为main的包
import "fmt"
//需要使用fmt包，该包为格式化IO的函数
func main(){
/*
多行注释
main 函数是每一个可执行程序所必须包含的，
一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）

当标识符以大写字母开头，则可被外部包的所使用:导出
小写字母开头，则对包外不可见，但在包内可见

运行GO代码命令:go run hello.go

*/
fmt.Println("Hello,World")

var s string 
fmt.Println("%q",s)


}


