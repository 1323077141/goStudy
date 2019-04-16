package main

import "fmt"

func main(){
/*
算术运算符:
+,-,*,/,%,++,--
*/

println("=======================算术运算符=====================")
var a int = 21
var b int = 10
var c int

c = a + b
fmt.Printf("第一行 - c 的值为 %d\n", c )
c = a - b
fmt.Printf("第二行 - c 的值为 %d\n", c )
c = a * b
fmt.Printf("第三行 - c 的值为 %d\n", c )
c = a / b
fmt.Printf("第四行 - c 的值为 %d\n", c )
c = a % b
fmt.Printf("第五行 - c 的值为 %d\n", c )
a++
fmt.Printf("第六行 - a 的值为 %d\n", a )
a=21   // 为了方便测试，a 这里重新赋值为 21
a--
fmt.Printf("第七行 - a 的值为 %d\n", a )

/*
关系运算符:
==	检查两个值是否相等，如果相等返回 True 否则返回 False。		(A == B) 为 False
!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。	(A != B) 为 True
>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。	(A > B) 为 False
<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。	(A < B) 为 True
>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	(A >= B) 为 False
<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。	(A <= B) 为 True 
*/
println("===========================关系运算符========================")
   a = 21
   b = 10
   if( a == b ) {
      fmt.Printf("第一行 - a 等于 b\n" )
   } else {
      fmt.Printf("第一行 - a 不等于 b\n" )
   }
   if ( a < b ) {
      fmt.Printf("第二行 - a 小于 b\n" )
   } else {
      fmt.Printf("第二行 - a 不小于 b\n" )
   } 
   
   if ( a > b ) {
      fmt.Printf("第三行 - a 大于 b\n" )
   } else {
      fmt.Printf("第三行 - a 不大于 b\n" )
   }
   /* Lets change value of a and b */
   a = 5
   b = 20
   if ( a <= b ) {
      fmt.Printf("第四行 - a 小于等于 b\n" )
   }
   if ( b >= a ) {
      fmt.Printf("第五行 - b 大于等于 a\n" )
   }
   
/*
逻辑运算符:
&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 	(A && B) 为 False
||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。	(A || B) 为 True
!	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。		(A && B) 为 True 
*/
println("=======================================逻辑运算符========================")
var a1 = true
var b1 = false
if ( a1 && b1 ) {
      fmt.Printf("第一行 - 条件为 true\n" )
   }
   if ( a1 || b1 ) {
      fmt.Printf("第二行 - 条件为 true\n" )
   }
   /* 修改 a 和 b 的值 */
   a1 = false
   b1 = true
   if ( a1 && b1 ) {
      fmt.Printf("第三行 - 条件为 true\n" )
   } else {
      fmt.Printf("第三行 - 条件为 false\n" )
   }
   if ( !(a1 && b1) ) {
      fmt.Printf("第四行 - 条件为 true\n" )
   }
/*
位运算符:&,|,^(异或,不同则为1)
&	按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 	(A & B) 结果为 12, 二进制为 0000 1100
|	按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或	(A | B) 结果为 61, 二进制为 0011 1101
^	按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。	(A ^ B) 结果为 49, 二进制为 0011 0001
<<	左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 	A << 2 结果为 240 ，二进制为 1111 0000
>>	右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 	A >> 2 结果为 15 ，二进制为 0000 1111
*/
var a2 uint = 60
var b2 uint = 13
var c2 uint = 0
println("===========================位运算符===============")
 c2 = a2 & b2       /* 12 = 0000 1100 */ 
   fmt.Printf("第一行 - c 的值为 %d\n", c2 )

   c2 = a2 | b2       /* 61 = 0011 1101 */
   fmt.Printf("第二行 - c 的值为 %d\n", c2 )

   c2 = a2 ^ b2       /* 49 = 0011 0001 */
   fmt.Printf("第三行 - c 的值为 %d\n", c2 )

   c2 = a2 << 2     /* 240 = 1111 0000 */
   fmt.Printf("第四行 - c 的值为 %d\n", c2 )

   c2 = a2 >> 2     /* 15 = 0000 1111 */
   fmt.Printf("第五行 - c 的值为 %d\n", c2 )
/*
赋值运算符
=	简单的赋值运算符，将一个表达式的值赋给一个左值	C = A + B 将 A + B 表达式结果赋值给 C
+=	相加后再赋值	C += A 等于 C = C + A
-=	相减后再赋值	C -= A 等于 C = C - A
*=	相乘后再赋值	C *= A 等于 C = C * A
/=	相除后再赋值	C /= A 等于 C = C / A
%=	求余后再赋值	C %= A 等于 C = C % A
<<=	左移后赋值 	C <<= 2 等于 C = C << 2
>>=	右移后赋值 	C >>= 2 等于 C = C >> 2
&=	按位与后赋值	C &= 2 等于 C = C & 2
^=	按位异或后赋值	C ^= 2 等于 C = C ^ 2
|=	按位或后赋值	C |= 2 等于 C = C | 2
*/
a = 21
println("====================赋值运算符====================")
c =  a
   fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )

   c +=  a
   fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )

   c -=  a
   fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )

   c *=  a
   fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )

   c /=  a
   fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )

   c  = 200; 

   c <<=  2
   fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )

   c >>=  2
   fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )

   c &=  2
   fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )

   c ^=  2
   fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )

   c |=  2
   fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )
/*
其他运算符:
&	返回变量存储地址	&a; 将给出变量的实际地址。
*	指针变量。		*a; 是一个指针变量
*/
var ptr *int
println("===========================指针与地址运算符============")
fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
   fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
   fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

   /*  & 和 * 运算符实例 */
   ptr = &a    /* 'ptr' 包含了 'a' 变量的地址 */
   fmt.Printf("a 的值为  %d\n", a);
   fmt.Printf("*ptr 为 %d\n", *ptr);
/*
运算符优先级:
7 	^ !
6 	* / % << >> & &^
5 	+ - | ^
4 	== != < <= >= >
3 	<-
2 	&&
1 	||
*/
println("==========================运算符优先级========================")
 var d int = 5
   var e int;
e = (a + b) * c / d;      // ( 30 * 15 ) / 5
   fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e );

   e = ((a + b) * c) / d;    // (30 * 15 ) / 5
   fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e );

   e = (a + b) * (c / d);   // (30) * (15/5)
   fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e );

   e = a + (b * c) / d;     //  20 + (150/5)
   fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e ); 










}
