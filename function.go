package main

import (
	"fmt"
	"math"
	"time"
	"net/http"
	"image"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

//type Abser interface {
//	Abs() float64
//}
//
//type Reader interface {
//	Read(b []byte) (n int, err error)
//}
//
//type Writer interface {
//	Write(b []byte) (n int, err error)
//}
//
//type ReaderWriter interface {
//	Reader
//	Writer
//}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)",p.Name,p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type Hello struct {}

func (h Hello) ServeHTTP(
w http.ResponseWriter,
r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	//v := &Vertex{3,4}
	//fmt.Println(v.AbsS())

	//f := MyFloat(-math.Sqrt2)
	//fmt.Println(f.Abs())

	//v := &Vertex{3, 4}
	//fmt.Println("before scaling: %+v Abs: %v\n", v, v.Abs())
	//v.Scale(5)
	//fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
	//var a Abser
	//f := MyFloat(-math.Sqrt2)
	//v := Vertex{3, 4}
	//
	//a = f
	//a = &v
	//a = v //cannot use v (type Vertex) as type Abser in assignment: Vertex does not implement Abser (Abs method has pointer receiver)


	//fmt.Println(a.Abs())

	//var w Writer
	//w = os.Stdout
	//
	//fmt.Fprintf(w, "hello, writer\n")

	//a := Person{"Arthur Dent", 42}
	//z := Person{"Zaphod Beeblebrox",900}
	//fmt.Println(a, z)

	//if err := run(); err != nil {
	//	fmt.Print(err)
	//}

	//r := strings.NewReader("Hello, Reader!")
	//b := make([]byte, 8)
	//for {
	//	n, err := r.Read(b)
	//	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	//	fmt.Printf("b[:n] = %q\n",b[:n])
	//}

	//var h Hello
	//err := http.ListenAndServe("localhost:4000",h)
	//if err != nil {
	//	log.Fatal(err)
	//}
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
    fmt.Println(m.At(0,0).RGBA())
}
