package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
)

func main() {
	// goroutines
	//go say("world")
	//say("hello")
	// channels
	//s :=[] int{7,2,8,-9,4, 0}
	//c:=make(chan int)
	//go sum(s[:len(s)/2],c)
	//go sum(s[len(s)/2:],c)
	//x,y := <-c,<-c
	//fmt.Println(x,y,x+y)
	//
	//go fun()
	//<-c1
	//fmt.Println(a)

	// producer and consumer
	s1 :=[] int{7,2,8,-9,4, 0}
	c1:=make(chan int)
	go producer(s1,c1)
	consumer(c1)
	for c:= range c1 {
		fmt.Println(c)
	}
	time.Sleep(3 * time.Second)


	//c := make(chan int, 5)
	//go fibonacci(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}

	fmt.Println(os.Args[0])


}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
var a string
var c1 = make(chan int)
func fun(){
	c1 <- 0
	a = "hello world"
}

func producer(s []int, c chan int){
	//for _,v:=range s {
	//	fmt.Printf("import the data %v\n",v)
	//	c<-v
	//}
	for  {
		c <- rand.Int()
		time.Sleep(1 * time.Second)
	}
}
func consumer(c chan int){
	for {
		fmt.Printf("export the data %v\n",<-c )
	}
}


func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < 3*n; i++ {
		c <- x
		fmt.Printf("import %v\n",x)
		time.Sleep(100 * time.Millisecond) //模拟时间间隔让来一个消耗一个
		x, y = y, x+y
	}

	close(c)
}