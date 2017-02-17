package fibonacci

import "fmt"

func fibonacci() func()int{
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

func Fibonacci1(n int64) int64 {
	if n < 2 {
		return n
	}
	return Fibonacci1(n - 1) + Fibonacci1(n - 2)
}