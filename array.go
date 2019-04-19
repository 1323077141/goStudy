package main
import(
"fmt"
)

/*
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
相对于去声明number0, number1, ..., and number99的变量，使用数组形式numbers[0], numbers[1] ..., numbers[99]更加方便且易于扩展。
数组元素可以通过索引（位置）来读取（或者修改），索引从0开始，第一个元素索引为 0，第二个索引为 1，以此类推。
数组声明需要指定元素类型及元素个数，语法格式如下:
var variable_name [SIZE] variable_type
此为一维数组的定义方式;

初始化数组:
以下演示了数组初始化：var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

访问数组:
数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：var salary float32 = balance[9]


多维数组:
声明:var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type,例:var threedim [5][10][4]int
二维数组:
var arrayName [ x ][ y ] variable_type
多维数组可通过大括号来初始值。以下实例为一个 3 行 4 列的二维数组：
a = [3][4]int{  
 {0, 1, 2, 3} ,     第一行索引为 0 
 {4, 5, 6, 7} ,     第二行索引为 1
 {8, 9, 10, 11},    第三行索引为 2 
}
或者
a = [3][4]int{  
 {0, 1, 2, 3} ,    第一行索引为 0 
 {4, 5, 6, 7} ,    第二行索引为 1 
 {8, 9, 10, 11}}   第三行索引为 2 
访问二维数组：
二维数组通过指定坐标来访问。如数组中的行索引与列索引，例如：
val := a[2][3]
或
var value int = a[2][3]


向函数传递数组:
向函数传递数组参数，需要在函数定义时，声明形参为数组，我们可以通过以下两种方式来声明：
一、
形参设定数组大小：
void myFunction(param [10]int)
{
.
.
.
}
二、
形参未设定数组大小：
void myFunction(param []int)
{
.
.
.
}
*/

func getAverage(arr []int,size int) float32{
var i,sum int 
var avg float32
for i = 0;i < size;i++ {
sum += arr[i]
}
avg = float32(sum)/float32(size)
return avg;
}

func main(){

var n [10] int
var i,j int
for i=0;i<10;i++{
n[i]=i+100
}

for j=0;j<10;j++{
fmt.Printf("Element[%d] = %d\n",j,n[j])
}

println("=========================二维数组=========================")
var array = [5][2] int{{0,0},{1,2},{2,4},{3,6},{4,8}}
for i=0;i<5;i++ {
	for j=0;j<2;j++ {
		fmt.Printf("array[%d][%d] = %d\n",i,j,array[i][j])
	}
}

println("============================数组作为形参====================")
/* 数组长度为 5 */
   var  balance = []int {1000, 2, 3, 17, 50}
   var avg float32

   /* 数组作为参数传递给函数 */
   avg = getAverage( balance, 5 ) ;

   /* 输出返回的平均值 */
   fmt.Printf( "平均值为: %f ", avg );
println("")
println("==============================精度问题=======================")
a1 := 1.69
b1 := 1.7
c1 := a1 * b1
fmt.Println(c1)
println("==============================设置固定精度后======================")
a2 := 1690
b2 := 1700
c2 := a2 * b2
fmt.Println(c2)
fmt.Println(float64(c2) / 1000000)


}
