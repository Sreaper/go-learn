package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main()  {
	sum :=0;
	fmt.Println("")
	for i:=0;i<10;i++{
		sum +=i;
	}
	fmt.Println(sum)
	// the init and post statement are optional
	for ; sum<100; {
		sum+=sum
	}
	fmt.Println(sum)
	// for is go`s "while"
	for sum < 200{
		sum+=sum
	}
	fmt.Println(sum)
	// forever
	//for{
	//    fmt.Println("love forever")
	//}
	// if
	fmt.Println(sqrt(2),sqrt(-4))
	fmt.Println(fmt.Sprint("szm","lyl"))
	// if convenient
	fmt.Println(
		pow(3,2,20),
		pow(3,3,20),
	)
	// if else
	fmt.Println(
		pow1(3,2,20),
		pow1(3,3,20),
	)
	//TODO  Exercise: Loops and Functions
        // switch
	fmt.Print("go run on")
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s" ,os)
	}
	// switch evaluation order
	fmt.Println("When`s Saturday?")
	today :=time.Now().Weekday()
	switch time.Saturday{
	case today+0:
		fmt.Println("tody")
	case today+ 1:
		fmt.Println("Tomorrow")
	case today+ 2:
		fmt.Println("in two days")
	default:
		fmt.Println("Too far away")
	}
	// switch with no condition
	t :=time.Now();
	switch  {
	case t.Hour()<12:
		fmt.Println("Good morning")
	case t.Hour()< 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
	// Defer
	defer fmt.Println("world")
	fmt.Println("hello")
	// Stacking defers
	fmt.Println("counting")
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")



}

func sqrt(x float64)(string) {
	if x <0{
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x,n ,lim float64)(float64){
	if x:=math.Pow(x,n);x< lim {
		return x
	}
	return lim
}
func pow1(x,n,lim float64)(float64){
	if x:=math.Pow(x, n) ;x<lim{
		return x
	} else {
		fmt.Printf("%g >=%g\n",x,lim)
	}
	return lim

}


