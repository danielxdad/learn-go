package main

import (
	// "context"
	// "errors"
	"fmt"
	"log"
	"math"
	"runtime"
	"slices"
	"sync"
	"unsafe"

	// "rsc.io/quote"
	"math/cmplx"
	"math/rand"

	// "net"
	"os"
	"os/user"
	"time"
	// "github.com/cockroachdb/errors/assert"
)

var i, j = 1, 2

// x := 10

func main() {
	log.SetPrefix("learn-go: ")
	log.SetFlags(0)

	hostname, _ := os.Hostname()
	user, _ := user.Current()
	fmt.Println(hostname, user.Username, user.HomeDir)

	fmt.Print("=========================================\n\n")

	// fmt.Println("Hello World")
	// fmt.Println(quote.Go())
	// message, err := Hello("barabun")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)

	names := []string{"barabun", "cachan", "foo", "bar"}
	if messages, err := Hellos(names); err != nil {
		log.Fatal(err)
	} else {
		for _, message := range messages {
			fmt.Println(message)
		}
	}

	fmt.Print("=========================================\n\n")

	d, err := time.ParseDuration("1h30m30s")
	now := time.Now()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(now)
	fmt.Println(now.Add(d))

	fmt.Print("=========================================\n\n")
	fmt.Println(add(3, 7))
	fmt.Println(split(17))

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Print("=========================================\n\n")
	basic_types()

	fmt.Print("=========================================\n\n")
	type_conversion()

	fmt.Print("=========================================\n\n")
	for_loop()

	fmt.Print("=========================================\n\n")
	if_stmt()

	fmt.Print("=========================================\n\n")
	fmt.Print("sqrt(2.0): ", sqrt(2.0), "\n")

	fmt.Print("=========================================\n\n")
	switch_stmt()

	fmt.Print("=========================================\n\n")
	defer_stmt()
	defer_stmt_1()

	fmt.Print("=========================================\n\n")
	pointers()

	fmt.Print("=========================================\n\n")
	structs()

	fmt.Print("=========================================\n\n")
	arrays()

	fmt.Print("=========================================\n\n")
	slices_types()

	fmt.Print("=========================================\n\n")
	maps()

	fmt.Print("=========================================\n\n")
	func_values()

	fmt.Print("=========================================\n\n")
	fib := fibonacci()
	for range 10 {
		fmt.Println(fib())
	}

	fmt.Print("=========================================\n\n")
	methods()

	fmt.Print("=========================================\n\n")
	interfaces()

	fmt.Print("=========================================\n\n")
	type_switch()

	fmt.Print("=========================================\n\n")
	var wg sync.WaitGroup
	wg.Go(panic_and_recover)
	wg.Wait()

	fmt.Print("=========================================\n\n")
	time_module()

	fmt.Print("=========================================\n\n")
	goroutimes()

	fmt.Print("=========================================\n\n")
	runtime_mod()

	fmt.Print("=========================================\n\n")
	x := "foo"
	x += "bar"
	fmt.Printf("string concatenation: %v\n", x)
}

func basic_types() {
	var (
		ToBe   bool       = false
		Str    string     = "foo bar"
		MaxInt uint64     = 1<<64 - 1
		Flt    float64    = 3.14
		z      complex128 = cmplx.Sqrt(-5 + 12i)
		ptr    uintptr    = 0xFF_FF_FF_FF_FF_FF_FF_FF
	)

	fmt.Printf("Type: %T, Value: %v, Sizeof: %v\n", ToBe, ToBe, unsafe.Sizeof(ToBe))
	fmt.Printf("Type: %T, Value: %v, Sizeof: %v\n", Str, Str, unsafe.Sizeof(Str))
	fmt.Printf("Type: %T, Value: %v, Sizeof: %v\n", MaxInt, MaxInt, unsafe.Sizeof(MaxInt))
	fmt.Printf("Type: %T, Value: %v, Sizeof: %v\n", Flt, Flt, unsafe.Sizeof(Flt))
	fmt.Printf("Type: %T, Value: %v, Sizeof: %v\n", z, z, unsafe.Sizeof(z))
	fmt.Printf("Type: %T, Value: %v, Sizeof: %v\n", ptr, ptr, unsafe.Sizeof(ptr))
}

func type_conversion() {
	var i int = 42
	var f float32 = float32(i)
	u := uint(f)

	fmt.Printf("%T: %v\n%T: %v\n%T: %v\n", i, i, f, f, u, u)
}

func for_loop() {
	sum := 1

	for i := range 10 {
		fmt.Println(i, i*2)
	}

	// Like a "while"
	for sum < 100 {
		sum += 1
	}
	fmt.Println(sum)

	// Infinite loop
	for {
		if rand.Intn(10) == 5 {
			fmt.Println("5, breaking infinite loop")
			break
		}
	}
}

func if_stmt() {
	if x := rand.Intn(10); x == 5 {
		fmt.Println("is 5")
	} else {
		fmt.Println("isn't 5")
	}

	// Doesnt work, "x" is only defined inside the "if" block scope
	// return x
}

func add(x, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqrt(x float64) float64 {
	z := 1.0

	for range 10 {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func switch_stmt() {
	f := func() string {
		return "linux"
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case f():
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Other")
	}
}

func defer_stmt() {
	// A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("world")

	fmt.Println("hello")

	// print: "hello world"
}

func defer_stmt_1() {
	fmt.Println("counting")

	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func pointers() {
	i := 10
	p := &i

	fmt.Println(p, *p)
	*p = 15
	fmt.Println(p, *p, i)
}

func structs() {
	type Vertex struct {
		x int
		y int
	}
	v := Vertex{2, 3}
	fmt.Println(v, v.x, v.y)
}

func arrays() {
	// var arr [2]string
	arr := [2]string{"hello", "world"}

	// arr[0] = "hello"
	// arr[1] = "world"

	// arr[2] = "foo"

	fmt.Println(arr)
}

func slices_types() {
	// array literal
	primes := []int{2, 3, 5, 7, 11, 13}
	// slice
	var s []int = primes[1:4]
	// slice literal
	var b []bool = []bool{true, false, true}
	// slice nil
	var snil []int

	// The make function allocates a zeroed array and returns a slice that refers to that array
	// len = 0, cap = 5
	var sarr []int = make([]int, 0, 5)

	slice_of_slices := [][]string{
		{"foo", "bar"},
		{"qaz", "zaq"},
	}

	s[1] = 10

	fmt.Println(s, len(s), cap(s))
	fmt.Println(primes, len(primes), cap(primes))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(snil, snil == nil, len(snil), cap(snil))
	fmt.Println(sarr, len(sarr), cap(sarr))
	fmt.Println(slice_of_slices)

	b = append(b, true, false)

	for i, v := range primes {
		fmt.Println(i, v)
	}

	// Delete items from index 1 up to index 3 (non included)
	primes = slices.Delete(primes, 1, 3)

	for i, v := range primes {
		fmt.Println(i, v)
	}
}

func maps() {
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{40.68433, -74.39967}

	// Map literal
	var m1 = map[string]Vertex{
		"Google": {37.42202, -122.08408},
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(m1["Google"])

	// Insert or update an element in map m:
	m1["foo"] = Vertex{3.14, 2.72}

	// Retrieve an element:
	elem := m1["foo"]
	fmt.Println(elem)

	// Delete an element:
	delete(m1, "foo")

	// Test that a key is present with a two-value assignment:
	// If key is not in the map, then elem is the zero value for the map's element type.
	elem, ok := m1["foo"]
	fmt.Println(elem, ok)
}

func func_values() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Printf("%T, %v\n", hypot, hypot(2, 3))
}

func fibonacci() func() int {
	// Fibonacci closure
	a, b, r := 0, 1, 1
	return func() int {
		r = a + b
		a, b = b, r
		return r
	}
}
