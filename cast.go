package main
import "fmt"

/*
类型转换:
type_name(expression)

*/


func main(){
var sum int = 17
var count int = 5
var mean float32

mean = float32(sum)/float32(count)
fmt.Printf("mean的值为%f\n",mean)


}

