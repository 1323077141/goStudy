package main

import "fmt"

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

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


fmt.Println("=============变量的默认值==============")
var i int
var f float64
var b bool
var s string
fmt.Printf("%v %v %v %q\n", i, f, b, s)


/*
二:根据值自行判定变量类型:var v_name = value
三:省略var,注意 := 左侧如果没有声明新的变量，就产生编译错误,:= 为声明语句.格式:  v_name := value
编译错误:var intVal int intVal := 1，产生错误;intVal,intVal1 := 1,2 此时不产生错误。
*/

fmt.Println("==============字符串不同格式下的输出===========")
var str = "hhhhh"
fmt.Printf("%s\n",str)
fmt.Printf("%q\n",str)

fmt.Println("=============声明方式============")
var intVal int
intVal,intVal1 := 1,2
fmt.Printf("%v %v\n",intVal,intVal1)

/*
多变量声明:
类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3
var vname1, vname2, vname3 = v1, v2, v3   和 python 很像,不需要显示声明类型，自动推断
vname1, vname2, vname3 := v1, v2, v3   出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
*/
println("====================多变量声明============")
g, h := 123, "hello"
println(x, y, a, b, c, d, e, f, g, h)

/*
值类型与引用类型：
值类型直接指向存在内存中的值,像int,float,bool,string,当使用=赋值时，实际上是在内存中将i的值进行了拷贝;
引用类型:一个引用类型的变量值存储的是值所在的内存地址，或内存地址中第一个字所在的位置，=赋值时，只有地址被复制
可以通过&i来获取变量的内存地址。值类型变量的值存储在栈中。

:= 赋值操作符,这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明
如果声明了一个局部变量却没有在相同的代码块中使用它，会得到编译错误;
但是全局变量是允许声明但不使用;
同一类型的多个变量可以声明在同一行:var a, b, c int
多变量可以在同一行进行赋值:
var a, b int
var c string
a, b, c = 5, 7, "abc"
并行 或 同时 赋值:
a, b, c := 5, 7, "abc"
如果想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同;
白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃
 _ 实际上是一个只写变量，不能得到它的值。这样做是因为 Go 语言中必须使用所有被声明的变量，但有时并不需要使用从一个函数得到的所有返回值。
并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。 
*/

/*
常量:
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
常量的定义格式：
const identifier [type] = value
可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型:
显式类型定义： const b string = "abc"
隐式类型定义： const b = "abc"
多个相同类型的声明可以简写为:const c_name1, c_name2 = value1, value2
*/
println("=============================常量声明===================")
const LENGTH int = 10
const WIDTH int = 5   
var area int
//const a, b, c = 1, false, "str" //多重赋值
area = LENGTH * WIDTH
fmt.Printf("面积为 : %d", area)
println()
//println(a, b, c)
   
/*
常量还可以用作枚举：
const (
    Unknown = 0
    Female = 1
    Male = 2
)
常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
*/


/*
iota
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
iota 可以被用作枚举值：
const (
    a = iota
    b = iota
    c = iota
)

第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
const (
    a = iota
    b
    c
)



const (
    i=1<<iota
    j=3<<iota
    k
    l
)
iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
简单表述:
i=1：左移 0 位,不变仍为 1;
j=3：左移 1 位,变为二进制 110, 即 6;
k=3：左移 2 位,变为二进制 1100, 即 12;
l=3：左移 3 位,变为二进制 11000,即 24。
*/
println("==========================itoa===================")

/*const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
)
fmt.Println(a,b,c,d,e,f,g,h,i)
//0 1 2 ha ha 100 100 7 8
*/


}




