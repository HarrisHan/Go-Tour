package main

import (
	fmt "fmt"
	"math/cmplx"
	"math"
	//"runtime"
	//"time"
	"strings"
)

func add(x, y int) int {
	return  x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, swift bool
var i, j int  = 1, 2

var (
	ToBe bool     = false
	MaxInt uint64 = 1<<64 -1
	z complex128  = cmplx.Sqrt(-5 + 12i)
)

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//func pow(x, y, z int) int {
//	if v := add(x, y); v > z {
//		return v
//	} else {
//		fmt.Printf("%g >= %g \n", v, z)
//	}
//	return z
//}

//type Vetex struct {
//	X int
//	Y int
//}
//
//var (
//	v1 = Vetex{1,2}
//	v2 = Vetex{X: 1}
//	v3 = Vetex{}
//	p = &Vetex{1,2} // 类型为*Vetex
//)

var a[10]int

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i]," "))
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

type Vetex struct {
	Lat, Long float64
}

//var m = map[string]Vetex {
//"Bell Labs": Vetex{
//40.68433, -74.39967,
//},
//"Google": Vetex{
//37.42202, -122.08408,
//},
//}

func compute(fn func(float64, float64) float64) float64 {
	return  fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//fmt.Println(add(111,222))
	//a, b := swap("hello", "world")

	//fmt.Println(a, b)
	//fmt.Println(split(17))

	//var i int
	//fmt.Println(i, c, python, swift)

	//var c, swift, oc  =	true, false, "no!"
	//fmt.Println(i, j, c, swift, oc)

	//var i, j int = 1, 2
	//k := 3
	//c, swift, oc := true, false, "yes!"
	//fmt.Println(i, j, k, c, swift, oc)

	//const f = "%T(%v)\n"
	//fmt.Printf(f, ToBe, ToBe)
	//fmt.Printf(f, MaxInt, MaxInt)
	//fmt.Printf(f, z, z)

	//var x, y int  = 3, 4
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	////var z uint = uint(f)
	//fmt.Println(x, y, f)

	//fmt.Println(needInt(Small))
	//fmt.Println(needFloat(Small))
	//fmt.Println(needFloat(Big))

	//sum := 1
	//for i := 0; i < 10 ; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//for sum < 1000 {
	//	sum += sum
	//}
	//fmt.Println(sum)

	//for {
	//	fmt.Println("死循环")
	//}

	//fmt.Println(sqrt(2), sqrt(-4))

	//fmt.Println(pow(2, 0, 3))

	//fmt.Print("Go runs on")
	//fmt.Printf("%s.", runtime.GOOS)
	//switch os := runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	fmt.Printf("%s.", os)
	//}

	//fmt.Println("when is saturday?")
	//today := time.Now().Weekday()
	//switch time.Saturday {
	//case today + 0:
	//	fmt.Println("Today.")
	//case today + 1:
	//	fmt.Println("Tomorrow")
	//default:
	//	fmt.Println("Toooooo far away")
	//}

	//t := time.Now()
	//switch  {
	//case t.Hour() < 12:
	//	fmt.Println("good morning")
	//case t.Hour() < 17:
	//	fmt.Println("good afternoon")
	//default:
	//	fmt.Println("good evening")
	//}

	//defer fmt.Println("world")
	//fmt.Println("hellow")

	//fmt.Println("counting")

	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//
	//fmt.Println("done")

	//i, j := 42, 2701
	//
	//p := &i
	//fmt.Println(*p)
	//*p = 21
	//fmt.Println(i)
	//
	//p = &j
	//*p = *p / 37
	//fmt.Println(j)

	//v := Vetex{1, 2}
	////v.X = 4
	//p := &v
	//p.X = 1e9
	//fmt.Println(v)

	//fmt.Println(v1, p, v2, v3)

	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0],a[1])
	//fmt.Println(a)

	//s := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println("s ==", s)
	//
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("s[%d] == %d\n", i, s[i])
	//}

	// Create a tic-tac-toe board.
	//game := [][]string{
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//}
	//
	//// The players take turns.
	//game[0][0] = "X"
	//game[2][2] = "O"
	//game[2][0] = "X"
	//game[1][0] = "O"
	//game[0][2] = "X"
	//
	//printBoard(game)

	//s := []int{2, 3, 4, 5, 11, 13}
	//fmt.Println("s ==", s)
	//fmt.Println("s[1:4] ==", s[1: 4])
	//
	////省略下表表示从零开始
	//fmt.Println("s[:3] ==", s[:3])
	////省略上标表示到len（s）结束
	//fmt.Println("s[4:] ==", s[4:])

	//a := make([]int, 5)
	//printSlice("a",a)
	//b := make([]int, 0, 5)
	//printSlice("b",b)
	//c := b[:2]
	//printSlice("c",c)
	//d := c[2:5]
	//printSlice("d",d)

	//var z []int
	//fmt.Println(z, len(z), cap(z))
	//if z == nil {
	//	fmt.Println("nil!")
	//}

	//var a []int
	//printSlice("a", a)
	//
	//// append works on nil slices.
	//a = append(a, 0)
	//printSlice("a", a)
	//
	//a = append(a,1)
	//printSlice("a",a)
	//
	//a = append(a, 2, 3, 4)
	//printSlice("a", a)

	//for i := range pow {
	//	pow[i] = 1 << uint(i)
	//}
	//
	//for _, v := range pow {
	//	fmt.Print("2**%d = %d\n", i, v)
	//}

	//m = make(map[string]Vetex)
	//m["Bell Labs"] = Vetex{
	//	40.68433, -74.39967,
	//}

	//fmt.Println(m)

	//m := make(map[string]int)
	//m["Answer"] = 42
	//fmt.Println("the value:", m["Answer"])
	//m["Answer"] = 48
	//fmt.Println("the value:", m["Answer"])
	//delete(m, "Answer")
	//fmt.Println("The value:", m["Answer"])

//	hypot := func(x, y float64) float64 {
//	return math.Sqrt(x*x + y*y)
//}
//	fmt.Println(hypot(5, 12))
//	fmt.Println(compute(hypot))
//	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
		    neg(-2*i),
		)
	}

}


