package main

import(
"fmt"
)

/*
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址;
指针声明:var var_name *var-type
var ip *int        指向整型
var fp *float32    指向浮点型 

空指针:
当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
一个指针变量通常缩写为 ptr

空指针判断:
if(ptr != nil)     ptr 不是空指针 
if(ptr == nil)     ptr 是空指针 

指针数组:如例

指向指针的指针:
如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址:pointer1存放pointer2的地址,pointer2存放变量var1的地址,var1存放值
声明如下:var ptr **int;

指针作为函数参数:

*/

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址的值 */
   *x = *y      /* 将 y 赋值给 x */
   *y = temp    /* 将 temp 赋值给 y */
}

func main(){
var a int = 10 /* 声明实际变量 */
fmt.Printf("变量的地址:%d\n",&a)
var ip *int        /* 声明指针变量 */
ip = &a  /* 指针变量的存储地址 */
/* 指针变量的存储地址 */
fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
/* 使用指针访问值 */
fmt.Printf("*ip 变量的值: %d\n", *ip )

println("================================空指针===================")
var  ptr *int
fmt.Printf("ptr 的值为 : %x\n", ptr  )

println("=============================指针数组====================")
println("一般数组:")
attr := []int{10,100,200}
var i int
attrLen := len(attr)
for i=0;i <attrLen;i++ {
fmt.Printf("attr[%d]=%d\n",i,attr[i])
}
println("指针数组:")
ptrAttr := make([]*int,attrLen)
//var ptrAttr [attrLen]*int;
for i=0;i<attrLen;i++ {
ptrAttr[i]=&attr[i]
}
for i=0;i<attrLen;i++ {
fmt.Printf("ptrAttr[%d]=%d\n",i,*ptrAttr[i])
}

println("===========================指向指针的指针====================")
var a4 int
var ptr4 *int
var pptr4 **int
a4 = 3000
ptr4 = &a4
pptr4 = &ptr4

fmt.Printf("变量a=%d\n",a4)
fmt.Printf("指针变量*ptr=%d\n",*ptr4)
fmt.Printf("指向指针的指针变量**pptr = %d\n",**pptr4)

println("==========================指针作为函数参数=====================")
var a5 int = 100
var b5 int = 200
fmt.Printf("交换前 a 的值 : %d\n", a5 )
fmt.Printf("交换前 b 的值 : %d\n", b5 )
swap(&a5, &b5);
fmt.Printf("交换后 a 的值 : %d\n", a5 )
fmt.Printf("交换后 b 的值 : %d\n", b5 )


}
