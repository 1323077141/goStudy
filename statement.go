package main

import "fmt"

func main(){
/*
if条件语句:
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
}else {
  /* 在布尔表达式为 false 时执行 */
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



}


