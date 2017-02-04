package main

import (
	"fmt"
	"strings"
	//"golang.org/x/tour/wc"
	"math"
)

func main() {
	// Pointers
	fmt.Println("start")
	i, j := 42, 2701
	p := &i;
	fmt.Println(*p) // 拿到指针的地址有什么用
	*p = 21;
	fmt.Println(i)

	p = &j;
	*p = *p / 37
	q := p
	fmt.Println(j)
	fmt.Println(q)
	// struct
	fmt.Println(Vertex{1, 2})
	v := Vertex{2, 3}
	v.X = 4
	fmt.Println(v.X)
	// Pointers to structs
	vv := Vertex{2, 3}
	vvv := &vv;
	vvv.X = 5
	fmt.Println(vv.X)
	fmt.Println(*vvv)
	// struct Literals
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X:1}
		v3 = Vertex{}
		pp = &Vertex{1, 2}
	)

	fmt.Println(v1, pp, v2, v3)

	// Arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	var sub []int = primes[1:4]
	fmt.Println(sub)

	// Slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)
	b1[0] = "szm"
	fmt.Println(a1, b1)
	fmt.Println(names)
	// TODO Slice literals slice和array数据之间的区别
	slice_s := [] int{2, 3, 5}
	fmt.Println(slice_s)
	// Slice struct
	slice_s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(slice_s1)
	// Slice default
	slice_s2 := []int{2, 3, 5, 7, 11, 13}
	slice_s2 = slice_s2[:]
	fmt.Println(slice_s2)
	slice_s2 = slice_s2[1:4]
	fmt.Println(slice_s2)
	slice_s2 = slice_s2[:2]
	fmt.Println(slice_s2)
	slice_s2 = slice_s2[1:]
	fmt.Println(slice_s2)
	// Slice length and capacity
	slice_s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(slice_s3)
	slice_s3 = slice_s3[:0]
	printSlice(slice_s3)
	// extend its length
	slice_s3 = slice_s3[:4]
	printSlice(slice_s3)
	slice_s3 = slice_s3[2:]
	printSlice(slice_s3)
	// nil slices
	var slice_s4 []int
	fmt.Println(slice_s4, len(slice_s4), cap(slice_s4))
	if slice_s4 == nil {
		fmt.Println("nil!")
	}
	// Creating a slice with make dynamically-sized arrays
	slice_s5 := make([]int, 5)
	printSlice1("slice_s5", slice_s5)
	slice_s5[1] = 3
	printSlice1("slice_s5", slice_s5)
	slice_s6 := slice_s5[:2]
	printSlice1("slice_s6", slice_s6)
	printSlice1("slice_s5", slice_s5)
	// slices of slices
	// create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// the player take turns
	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	// Appending to a slice
	var slice_s7 []int
	printSlice(slice_s7)
	//var arr_1 [3]int
	var array = [5]int{1, 2, 3, 4, 5}
	printArr(array)
	slice_s7 = append(slice_s7, 0)
	printSlice(slice_s7)
	slice_s7 = append(slice_s7, 2, 3)
	printSlice(slice_s7)
	// range
	var pow = []int{1, 2, 4, 8, 16, 32}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow1 := make([]int, 10)
	for i := range pow1 {
		// TODO 符号位和非符号位的区别
		pow1[i] = 1 << uint(i)
	}
	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}
	//TODO slice exercise
	//Maps
	var m map[string]Vertex1
	m = make(map[string]Vertex1)
	m["szm"] = Vertex1{40, 30}
	m["lyl"] = Vertex1{40, 30}
	fmt.Println(m)
	var m1 = map[string]Vertex1{
		"szm1":Vertex1{
			40, 30,
		},
	}
	fmt.Println(m1)
	var m2 = map[string]Vertex1{
		"szm":{40, 40},
	}
	fmt.Println(m2)
	//Mutating Maps
	m3 := make(map[string]int)
	m3["szm"] = 30
	fmt.Println("The value:", m3["szm"])
	m3["szm"] = 31
	fmt.Println("The value:", m3["szm"])
	delete(m3, "szm")
	fmt.Println("The value:", m3["szm"])
	m3["szm"] = 30
	m1_1, ok := m3["szm"]
	fmt.Println("The value:", m1_1, "Present?", ok)
	// Maps exercise
	//wc.Test(WordCount1)
	// function values
	hypot := func(x,y float64) float64 {
		return math.Sqrt(x*x +y*y)
	}
	fmt.Println(hypot(5,12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	//Function closures
	pos,neg :=adder(),adder()
	for i = 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	middle :=adder1()
	for i = 0; i < 10; i++ {
		fmt.Println(
			middle(i),
		)
	}
	// fibonacci
	//f:=fibonacci1()
	//for i = 0; i < 10; i++ {
	//	fmt.Println(f())
	//}
	testPro()

}

// struct
type Vertex struct {
	X int
	Y int
}

type Vertex1 struct {
	Lat, LOng float64
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
func printArr(array [5]int) {
	fmt.Println(array)
}
func WordCount(s string) map[string]int {
	return map[string]int{"x":1}
}
func WordCount1(s string) map[string]int {
	var map_1 = make(map[string]int)
	strs := strings.Fields(s)
	length := len(strs);
	for i := 0; i < length; i++ {
		//map_1[strs[i]] = map_1[strs[i]] + 1
		(map_1[strs[i]])++
	}
	return map_1
}
func compute(fn func(float64,float64) float64) float64{
	return fn(3,4)
}
func adder() func(int) int{
	sum:=0
	return func(x int)int{
		sum+=x
		return sum
	}
}
func adder1() func(int) int{
	sum:=0
	return func(x int) int{
		sum++
		fmt.Println("sum value",sum)
		return 5
	}
}

func fiboncci() func() int{
	first:=0
	second :=1
	result :=0
	index :=0
	return func() int{
		if index == 0 {
                  	index++
			result = first
			return result
		}
		if index == 1 {
			index++
			result = second
			return result
		}
		if index >= 2 {
			temp := second
			result = first + second
			first = temp
			second = result
		}
		return result
	}
}
func fibonacci1() func()int{
	x:=0
	y:=1
	//index :=0
	return func() int{
		//if index == 0 {
		//	index++
		//	return x
		//}
		x,y = y,x+y
                fmt.Println("x value" ,x ,"y value",y)
		return x
	}
}

func testPro() int{
	x:=2
	y:=1
	// 先评估,再进行赋值运算
	x,y = y,x+y
	fmt.Println("x value" ,x ,"y value",y)
	return 0
}