package main

import "fmt"

func main(){
/*
go文件中，一行代表一个语句结束
单行注释以双斜杠开头
标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字
关键字：
break,default,func,interface,select,case,defer,go,map,struct,chan,else,goto,package,switch,const,fallthrough,if,range,type,continue,for,import,return,var
保留字:
append,bool,byte,cap,close,complex,complex64,complex128,unit16,copy,false,flat32,float64,imag,int,int8,int 16,int32,int64,iota,len,make,new,nil,panic,unit64,print,println,real,recover,string,true,unit,unit8,unitprt

数据类型：
布尔型：布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
数字类型：整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码
字符串类型：字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本
派生类型:包括指针，数组，结构化，Channel，函数,切片,接口(interface),Map类型

数字类型中:
uint8:无符号 8 位整型 (0 到 255);uint16:无符号 16 位整型 (0 到 65535);uint32:无符号 32 位整型 (0 到 4294967295);uint64:无符号 64 位整型 (0 到 18446744073709551615);int8:有符号 8 位整型 (-128 到 127);int16:有符号 16 位整型 (-32768 到 32767);int32:有符号 32 位整型 (-2147483648 到 2147483647);int64:有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
浮点型:float32:32位浮点型数;float64:64位浮点型数;complex64:32 位实数和虚数;complex128:64 位实数和虚数
其他数字类型:byte:类似 uint8;rune:类似 int32;uint:32 或 64 位;int:与 uint 一样大小;uintptr:无符号整型，用于存放一个指针
*/

/*
变量:
变量的声明使用var关键字:var identifier type
一:指定变量类型,如果没有初始化，则变量默认为零值,零值是变量没有做初始化时系统默认设置的值。
var v_name v_type
数值类型:0;布尔类型:false;字符串为"";其他为nil(指针,数组,Map,chan,func,interface)
*/
var i int 
var f float64
var b bool
var s string
fmt.Print(" %v %v %v %q \n",i,f,b,s)

/*
二:根据值自行判定变量类型:var v_name = value
三:省略var,注意 := 左侧如果没有声明新的变量，就产生编译错误,:= 为声明语句.格式:  v_name := value
编译错误:var intVal int intVal := 1，产生错误;intVal,intVal1 := 1,2 此时不产生错误。

多变量声明:
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
*/







}




