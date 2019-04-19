package main

import "fmt"

func main(){
println("================================if=========================")
/*
if条件语句:
if 布尔表达式 {
    在布尔表达式为 true 时执行 
}else {
   在布尔表达式为 false 时执行 
}
*/
 var a int = 100;
 
   /* 判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" );
   } else {
       /* 如果条件为 false 则执行以下语句 */
       fmt.Printf("a 不小于 20\n" );
   }
   fmt.Printf("a 的值为 : %d\n", a);

/*
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
case 自带break
若需要执行后面的 case，可以使用 fallthrough
变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3
使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true
switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型：
switch x.(type){
    case type:
       statement(s);      
    case type:
       statement(s); 
     可以定义任意个数的case 
    default:  可选 
       statement(s);
}
*/

println("============================switch================================")

 /* 定义局部变量 */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }

   switch {
      case grade == "A" :
         fmt.Printf("优秀!\n" )     
      case grade == "B", grade == "C" :
         fmt.Printf("良好\n" )      
      case grade == "D" :
         fmt.Printf("及格\n" )      
      case grade == "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" );
   }
   fmt.Printf("你的等级是 %s\n", grade ); 

 var x interface{}
     
   switch i := x.(type) {
      case nil:   
         fmt.Printf(" x 的类型 :%T",i)                
      case int:   
         fmt.Printf("x 是 int 型")                       
      case float64:
         fmt.Printf("x 是 float64 型")           
      case func(int) float64:
         fmt.Printf("x 是 func(int) 型")                      
      case bool, string:
         fmt.Printf("x 是 bool 或 string 型" )       
      default:
         fmt.Printf("未知型")     
   }   


 switch {
    case false:
            fmt.Println("1、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("2、case 条件语句为 true")
            fallthrough
    case false:
            fmt.Println("3、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("4、case 条件语句为 true")
    case false:
            fmt.Println("5、case 条件语句为 false")
            fallthrough
    default:
            fmt.Println("6、默认 case")
    }

/*
select:
select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的.
select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s); 
     你可以定义任意数量的 case 
    default :  可选 
       statement(s);
}

每个 case 都必须是一个通信
所有 channel 表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行，其他被忽略。
如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。 
否则：
如果有 default 子句，则执行该语句。
如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值
*/
println("===================================select==========================")
var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    

}


