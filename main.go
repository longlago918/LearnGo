package main

import(
"fmt"
"goProject/learnFile"
)
const PI = 3.13145
const Num = 1 << 3
// 枚举
const (
	type1 = 1 << iota + 5
	type2 
	type3 
)
func main(){

	fmt.Println("learn begin")

	println(type1)
	println(type2)
	println(type3)

	
	println(Num)

	t1 := Learn.Teacher{Name:"paul",Age:40,Object:"英语"}
	t1.Teach()
	
	Learn.Play()

	varTest()

	rangeTest()

	// arrTest()
	sliceTest()

	mapTest()
	fmt.Println("learn end")
}

// 变量申明测试
func varTest(){

	// var age int
	// var name string
	// var arr [10] int
	// var mArr [] int
	
	// var bird struct{

	// 	age int
	// }

	// var m map[string] int
	var (
		a int
		s string
	)

	a = 10
	s = "ahaha"
	fmt.Println(a)
	fmt.Println(s)
		
	var block func(s string) int

	block = func(s string)int {
		fmt.Println(s + "block")
		return 1
	}
	block("i'm is ")
	// 左边变量必须是没有申明过的
	height := 1.8
	fmt.Println(height)

	// 这种方法左边变量不能都是申明过的，应该至少有一个没有被申明
	height,y1 := 20,"china"
	fmt.Println(height)
	fmt.Println(y1)




	_, _, str := fTest("str1","str2","str3")
	fmt.Println(str)
	
	fmt.Println(PI)
}

func fTest(s1,s2,s3 string)(rs1,rs2,rs3 string){
	
	s3 += " apended"
	return s1,s2,s3
}

func rangeTest(){
	str := "hello,世界"
	// 每个字符类型是rune
	for i,ch := range str {
		fmt.Println(i , ch)
	}
	// 每个字符是byte
	print("\n")
	for index := 0; index < len(str); index++ {
		fmt.Println(index,str[index])
	}
}

/// 数组是值类型，作为参数传递都将产生一次复制动作，无法修改传入数组的内容
/// oc的不可变数组是长度不能增加，并且元素也不能修改，但是Go可以修改元素
func arrTest(){

	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)
	// 改变不了arr，因为Go函数参数传递数的是值类型
	modify(arr)
	fmt.Println(arr)

	arr[0] = 100
	fmt.Println(arr)
}

func modify(array [5]int){

	array[0] = 10
	fmt.Println(array)
}

func sliceTest(){

	fmt.Println("slice test begin")

	arr := [10]int{1,2,3,4,5,6,7,8,9,10}

	// 从数组创建是slice 
	// var slice []int
	var slice []int = arr[2:5]  //[:] 包含所有原数组  [:5] 从0开始到多少个元素创建slice  [5:] 从第5个元素到所有
	fmt.Println(slice)

	// 方括号里面写容量就是数组，不写就是slice，只有slice才能append
	slice2 := []string{"a","b","c"}
	fmt.Println(slice2,len(slice2),cap(slice2))
	slice2 = append(slice2,"d","e","f","g")
	// 往slice里添加一个元素后，如果容量不够大会向外扩充一倍
	fmt.Println(slice2,len(slice2),cap(slice2))
	slice2 = append(slice2,"x")
	fmt.Println(slice2,len(slice2),cap(slice2))
	fmt.Println("slice test end")
}

type Person struct{

	ID string
	Name string
	Phone string
}

func mapTest(){

	a := true ? 1 : 2;
	fmt.Println("三目运算符测试",a)
	// make(map[string] PersonInfo) map创建
	// make([]int, 5) 切片创建

	fmt.Println("map test begin")

	// personMap := map[string]Person
	var personMap map[string]Person
	personMap = make(map[string]Person)
	personMap["Num 1"] = Person{ID:"1",Name:"name1"}
	fmt.Println("map test",personMap)

	fmt.Println("map test end")
}