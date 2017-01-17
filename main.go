package main

import (
	"fmt"
	"math"
)

var c, python, java bool

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	fmt.Println("hello")
	fmt.Println(math.Pi)
	fmt.Println(add(2, 3))
	fmt.Println(add1(2, 3))
	a, b := swap("szm", "lyl")
	fmt.Println(a, b)
	//多个变量声明、定义的另一种形式
	x, y := 2, "ToSmile"
	fmt.Printf("学号:%d \t 姓名:%s \n", x, y)

	//Go中有一个特殊的变量_ 任何赋给它的值将被丢弃
	_, Ret := 2, 3
	fmt.Printf("变量的值为：%d \n", Ret)
	var golan int
	fmt.Println(golan, c, python, java)
	// short variable declaration ,not defined outside the function
	s := 10;
	fmt.Println(s)
	fmt.Printf("Type: %T Value: %v\n", s, s)
	// conversion
	var ff float32 = float32(s)
	var fff float32 = 30.11
	fmt.Println(ff)
	fmt.Println(fff)
	// type inference
	v := 42.333
	fmt.Printf("v is of type %T\n", v)
	// Constants
	const Pi = 3.14;
	fmt.Println(Pi)
	// numeric constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// int64
	fmt.Println(int64(Small))

}

func add(x int, y int) int {
	return x + y;
}

func add1(x, y int) int {
	return x + y;
}

func swap(x, y string) (string, string) {
	return y, x
}
func needInt(x int) int {
	return x * 10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

