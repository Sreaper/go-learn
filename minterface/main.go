package main

import (
	"fmt"
	"math"
	"time"
	"strconv"
	"strings"
	"io"
	"golang.org/x/tour/reader"
	"os"
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

func main() {
	fmt.Println("start")
	// method
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// methods are functions
	fmt.Println(Abs(&v))
	// methods on non-struct types
	// methods declare a method with a receiver whosetype is defined in the
	// same package as the method
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	// Pointer receivers
	// with a value receiver,the Scale method operates on a copy of the
	// original Vertex value means that the original properties never change
	v1 := Vertex{3, 4}
	v1.Scale(10)
	fmt.Println(v1.Abs())
	// Pointers and functions
	v2 := Vertex{3, 4}
	// argument value clone the value
	fmt.Println("function argument  value and point diff")
	Scale(&v2, 10)
	Abs1(v2)
	fmt.Println(Abs1(v2))

	v2_1 := Vertex{3, 4}
	Scale1(v2_1, 10)
	fmt.Println(Abs1(v2_1))

	// methods and pointer indirection
	fmt.Println("methods and pointter indirection")
	v3 := Vertex{3, 4}
	v3.Scale(2)
	ScaleFunc(&v3, 10)
	// ScaleFunc(v3,10) // Compile error!
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println("Methods and pointer indirection")
	fmt.Println(v3, p)

	// the function argument type must be consistency
	fmt.Println("the function argument type must be consistency")
	v4 := Vertex{3, 4}
	fmt.Println(v4.Abs())
	fmt.Println(AbsFunc(v4))
	// fmt.Println(AbsFunc(&v4)) compile error
	p1 := &Vertex{3, 4}
	fmt.Println(p1.Abs())
	fmt.Println(AbsFunc(*p1)) // convert the point to value
	// Choosing a value or pointer receiver
	v5 := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v,Abs:%v\n", v5, v5.Abs())
	v5.Scale(5)
	fmt.Printf("Before scaling: %+v,Abs:%v\n", v5, v5.Abs())
	// interfaces
	// An interface type is defined as a set of method signatures.
	var a Abser
	f1 := MyFloat(-math.Sqrt2)
	v6 := Vertex{3, 4}
	a = f1
	a = &v6 //TODO receiver  is the value is ok ,but receiver is point is compile is error
	a = v6
	fmt.Println(a.Abs())
	// interfaces are implemented implicitly
	var i I = &T{"hello", "Hello1"}
	i.M()

	// interface value and type
	var i1 I
	i1 = &T{"Hello", "Hello1"}
	describe(i1)
	i1.M1()

	i1 = F(math.Pi)
	describe(i1)
	i1.M()
	// interface values with nil underlying value
	fmt.Println("interface value with nil underlying value")
	var i2 I
	var t *T
	i2 = t
	i2.M1()
	describe(i2)
	t3 := &T{"hello", "hello2"}
	i2 = t3;
	i2.M1()

	// Nil interface values
	//var i3 I
	//describe(i3)
	// TODO error日志的理解
	//i3.M1()
	// The empty interface
	fmt.Println("The empty interface")
	var i4 interface{}
	describe1(i4)
	//Type assertions
	fmt.Println("type assertions")
	var i5 interface{} = "hello"
	s := i5.(string)
	fmt.Println(s)
	s, ok := i5.(string)
	fmt.Println(s, ok)
	f3, ok := i5.(float64)
	fmt.Println(f3, ok)
	// f3=i5.(float64)
	// fmt.Println(f3) runtime error
	// Type switches
	do(21)
	do("hello")
	do(true)
	// Stringers one of the most ubiquitous interfaces
	person1 := Person{"szm", 30}
	person2 := Person{"lyl", 30}
	fmt.Println(person1, person2)
	// Exercise Stringers
	hosts := map[string]IPAddr{
		"loopback":{127, 0, 0, 1},
		"googleDNS":{8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	// Errors
	fmt.Println("Erroros info")
	if err := run(); err != nil {
		fmt.Println(err)

	}

	// str convert integer
	i6, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn`t convert number: %v\n", err)
	}
	fmt.Println("Converted integer:", i6)
	// Exercise:Errors
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	// Readers
	r := strings.NewReader("Hello, Reader! 宋")
	b := make([]byte, 15)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	// validate myReader
	reader.Validate(MyReader{})
	// rot13Reader
	s1:= strings.NewReader("Lbh penpxrq gur pbqr!")
	r1:=rot13Reader{s1}
	io.Copy(os.Stdout,&r1)
	fmt.Println("")

	// image
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
	// other image
	m1:=Image{100,100, 128}
	pic.ShowImage(&m1)


}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
// 怎么还给重写,方法不能复用吗
func Abs(v *Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs1(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale1(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type Abser interface {
	Abs() float64
}

type I interface {
	M()
	M1()
}

type T struct {
	S  string
	S1 string
}

//TODO 不让重复的多写是指针point或则是value
func (t *T) M() {
	fmt.Println(t.S)
}

func (t *T) M1() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func (f F) M1() {
	fmt.Println(f * 10)
}
func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}

func describe1(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v * 2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don`t know aout type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var temp string = ""
	for i := 0; i < len(ip); i++ {
		if i != 0 {
			temp += "."
		}
		temp += fmt.Sprint(ip[i])
	}
	return temp
}

type MyError struct {
	When time.Time
	What string
}
// TODO choose to print Error or String ?
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.When, e.What)
}
func (e *MyError) String() string {
	return fmt.Sprintf("MyError")
}

func run() error {
	return &MyError{
		time.Now(),
		"it did`t work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cann`t Sqrt negative number:%g", e)
}
func Sqrt(x MyFloat) (MyFloat, error) {
	if x > 0 {
		return x, nil
	} else {
		return x, ErrNegativeSqrt(x)
	}
}
//func (x MyFloat) String() string{
//	return fmt.Sprintf("test",float64(x))
//}

type MyReader struct {

}

func (reader MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'
	return 1, nil
}

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)
	for i := 0; i < len(b); i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return
}

type Image struct {
	Width,Height int
	color uint8
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0,0,i.Width,i.Height)
}

func (i *Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (i *Image) At(x,y int) color.Color {
	return color.RGBA{i.color +uint8(x),i.color+uint8(y),255,255 }
}