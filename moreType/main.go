package main

import "fmt"

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
	slice_s3 =slice_s3[:0]
	printSlice(slice_s3)
	// extend its length
	slice_s3 = slice_s3[:4]
	printSlice(slice_s3)
	slice_s3 = slice_s3[2:]
	printSlice(slice_s3)
	// nil slices
	var slice_s4 []int
	fmt.Println(slice_s4,len(slice_s4),cap(slice_s4))
	if slice_s4 == nil{
		fmt.Println("nil!")
	}
	// Creating a slice with make dynamically-sized arrays
	slice_s5 :=make([]int,5)
	printSlice1("slice_s5",slice_s5)
	slice_s5[1] = 3
	printSlice1("slice_s5",slice_s5)
	slice_s6 :=slice_s5[:2]
	printSlice1("slice_s6",slice_s6)
	printSlice1("slice_s5",slice_s5)
	// slices of slices
	// create a tic-tac-toe board.
	board :=[][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}



}

// struct
type Vertex struct {
	X int
	Y int
}

func printSlice(s []int){
   fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}