package main
import(
"fmt"
)

/*
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，可以像迭代数组和切片那样迭代它。不过，Map 是无序的，无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

定义 Map
可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
声明变量，默认 map 是 nil:
var map_variable map[key_data_type]value_data_type
使用 make 函数:
map_variable := make(map[key_data_type]value_data_type)
如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

delete() 函数
delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key
*/

func main(){
//创建集合
var countryCapitalMap map[string]string
countryCapitalMap = make(map[string]string)

countryCapitalMap["France"] = "Paris"
countryCapitalMap["Italy"] = "罗马"
countryCapitalMap["Japan"] = "东京"
countryCapitalMap["India"] = "新德里"

for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }


/*查看元素在集合中是否存在 */
    capital, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(capital) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("美国的首都是", capital)
    } else {
        fmt.Println("美国的首都不存在")
    }

println("==========================delete=====================")
fmt.Println("原始地图")

        /* 打印地图 */
        for country := range countryCapitalMap {
                fmt.Println(country, "首都是", countryCapitalMap [ country ])
        }

        /*删除元素*/ delete(countryCapitalMap, "France")
        fmt.Println("法国条目被删除")

        fmt.Println("删除元素后地图")

        /*打印地图*/
        for country := range countryCapitalMap {
                fmt.Println(country, "首都是", countryCapitalMap [ country ])
        }

}
