package main
import (
"fmt"
"math"
)

/*
函数定义格式:
func function_name( [parameter list] ) [return_types] {
   函数体
}
函数定义解析：
func：函数由 func 开始声明
function_name：函数名称，函数名和参数列表一起构成了函数签名。
parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
函数体：函数定义的代码集合。
当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。
调用函数，向函数传递参数，并返回值

函数参数:
函数如果使用参数，该变量可称为函数的形参。
形参就像定义在函数体内的局部变量。
调用函数，可以通过两种方式来传递参数：
值传递	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
引用传递	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

函数用法:
函数作为值：
声明函数变量 
getSquareRoot := func(x float64) float64 {
   return math.Sqrt(x)
 }

使用函数 
fmt.Println(getSquareRoot(9))
闭包：匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明：
方法：Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：
func (variable_name variable_data_type) function_name() [return_type]{
   函数体
}
*/

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 声明局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}

/*函数返回多个值*/
func swap(x, y string) (string, string) {
   return y, x
}

func swap1(x,y int) int{
var temp int
temp = x
x = y
y = temp
return temp
}

func swap2(x *int,y *int){
var temp int
temp = *x
*x = *y
*y = temp
}

/*
闭包函数：
匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明
以下实例中，创建函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量
*/

func getSe1() func() int{
i := 0
return func() int{
i+=1
return i
}
}

/*
Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：
func (variable_name variable_data_type) function_name() [return_type]{
    函数体
}
*/

/* 定义结构体 */
type Circle struct {
  radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}

func main(){
 /* 定义局部变量 */
   var a int = 100
   var b int = 200
   var ret int
   
   /* 调用函数并返回最大值 */
   ret = max(a, b)

fmt.Printf("最大值是:%d\n", ret)

c,d := swap("Mahesh","Kumar")
fmt.Println(c,d)

println("============================值传递==============================")
fmt.Printf("交换前 a 的值为 : %d\n", a )
   fmt.Printf("交换前 b 的值为 : %d\n", b )

   /* 通过调用函数来交换值 */
   swap1(a, b)

   fmt.Printf("交换后 a 的值 : %d\n", a )
   fmt.Printf("交换后 b 的值 : %d\n", b )

println("================================引用传递===========================")
fmt.Printf("交换前，a 的值 : %d\n", a )
   fmt.Printf("交换前，b 的值 : %d\n", b )
   /* 调用 swap() 函数
   * &a 指向 a 指针，a 变量的地址
   * &b 指向 b 指针，b 变量的地址
   */
   swap2(&a, &b)
   fmt.Printf("交换后，a 的值 : %d\n", a )
   fmt.Printf("交换后，b 的值 : %d\n", b )
println("================================函数作为值============================")
getSquareRoot := func(x float64) float64{
return math.Sqrt(x)
}
fmt.Println(getSquareRoot(9))
println("================================函数闭包=============================")
nextNumber := getSe1()
fmt.Println(nextNumber())
fmt.Println(nextNumber())
fmt.Println(nextNumber())
nextNumber1 := getSe1()
fmt.Println(nextNumber1())
fmt.Println(nextNumber1())
println("===================================函数方法=========================")
var c1 Circle
  c1.radius = 10.00
  fmt.Println("圆的面积 = ", c1.getArea())


}




