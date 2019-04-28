package main

import(
"fmt"
)

/*
切片是对数组的抽象,数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
声明一个未指定大小的数组来定义切片：var identifier []type (切片不需要说明长度)
或使用make()函数来创建切片:var slice1 []type = make([]type, len)
也可以简写为slice1 := make([]type, len)
也可以指定容量，其中capacity为可选参数:make([]T, length, capacity), len(length) 是数组的长度并且也是切片的初始长度。

切片初始化:
s :=[] int {1,2,3 } :直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
s := arr[:] 初始化切片s,是数组arr的引用
s := arr[startIndex:endIndex] 	将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
s := arr[startIndex:] 缺省endIndex时将表示一直到arr的最后一个元素
s := arr[:endIndex] 缺省startIndex时将表示从arr的第一个元素开始
s1 := s[startIndex:endIndex] 通过切片s初始化切片s1
s :=make([]int,len,cap) 通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片

当指定cap与len时，初始化有默认值

len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度。
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

空(nil)切片
一个切片在未初始化之前默认为 nil，长度为 0

切片截取
可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]

append() 和 copy() 函数
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来

*/

func printSlice(x []int){
fmt.Printf("len=%d,cap=%d,slice=%v\n",len(x),cap(x),x)
}

func main(){
println("=======================len与cap===================")
var numbers = make([]int,3,5)
printSlice(numbers)
println("======================nil切片=====================")
if(numbers == nil){
fmt.Println("切片为空")
}
var numbers1 []int
printSlice(numbers1)
if(numbers1 == nil){
fmt.Println("切片numbers为空")
}

println("=====================切片截取=====================")
number := []int{0,1,2,3,4,5,6,7,8}
printSlice(number)
fmt.Println("number == ",number)
//从1开始，不包括4
fmt.Println("number[1:4] == ",number[1:4])
//默认下限为0
fmt.Println("number[:3] == ",number[:3])
//默认上限为len(s)
fmt.Println("number[4:] == ",number[4:])
number1 := make([]int,0,5)
printSlice(number1)
number2 := number1[:2]
printSlice(number2)
number3 := number1[2:5]
printSlice(number3)


println("=========================append与copy=================")
var numbers7 []int
printSlice(numbers7)
/* 允许追加空切片 */
numbers7 = append(numbers7, 0)
printSlice(numbers7)
/* 向切片添加一个元素 */
numbers7 = append(numbers7, 1)
printSlice(numbers7)
/* 同时添加多个元素 */
numbers7 = append(numbers7, 2,3,4)
printSlice(numbers7)
/* 创建切片 numbers1 是之前切片的两倍容量*/
numbers8 := make([]int, len(numbers7), (cap(numbers7))*2)
/* 拷贝 numbers 的内容到 numbers1 */
copy(numbers8,numbers7)
printSlice(numbers7)   

}
