package main

import (
	// "context"
	// "errors"
	"fmt"
	"log"
	"math"
	"runtime"
	"sync"

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

	fmt.Println("=========================================\n")

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

	fmt.Println("=========================================\n")

	d, err := time.ParseDuration("1h30m30s")
	now := time.Now()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(now)
	fmt.Println(now.Add(d))

	fmt.Printf("asd")

	fmt.Println("=========================================\n")
	fmt.Println(add(3, 7))
	fmt.Println(split(17))

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Println("=========================================\n")
	basic_types()

	fmt.Println("=========================================\n")
	type_conversion()

	fmt.Println("=========================================\n")
	for_loop()

	fmt.Println("=========================================\n")
	if_stmt()

	fmt.Println("=========================================\n")
	fmt.Println("sqrt:", sqrt(8.0))

	fmt.Println("=========================================\n")
	switch_stmt()

	fmt.Println("=========================================\n")
	defer_stmt()
	defer_stmt_1()

	fmt.Println("=========================================\n")
	pointers()

	fmt.Println("=========================================\n")
	structs()

	fmt.Println("=========================================\n")
	arrays()

	fmt.Println("=========================================\n")
	slices()

	fmt.Println("=========================================\n")
	maps()

	// fmt.Println("=========================================\n")
	// resolver := new(net.Resolver)
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// addrs, err := resolver.LookupHost(ctx, "secure.etecsa.cu")
	// addrs, err := net.LookupHost("secure.etecsa.cu")
	// if err != nil {
	// 	dnsError, isDNSError := err.(*net.DNSError)
	// 	if isDNSError {
	// 		fmt.Printf("[ERROR]: %T, %v\n", err, err)
	// 		fmt.Println(dnsError.IsTimeout, dnsError.IsTemporary, dnsError.Server)
	// 	} else {
	// 		fmt.Printf("[ERROR]: %T, %v\n", err, err)
	// 	}
	// } else {
	// 	fmt.Println(addrs)
	// }

	fmt.Println("=========================================\n")
	func_values()

	fmt.Println("=========================================\n")
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

	fmt.Println("=========================================\n")
	methods()

	fmt.Println("=========================================\n")
	interfaces()

	fmt.Println("=========================================\n")
	type_switch()

	fmt.Println("=========================================\n")
	var wg sync.WaitGroup
	wg.Go(panic_and_recover)
	wg.Wait()

	fmt.Println("=========================================\n")
	time_module()
}

func basic_types() {
	var (
		ToBe   bool       = false
		Str    string     = "foo bar"
		MaxInt uint64     = 1<<64 - 1
		Flt    float64    = 3.14
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", Str, Str)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Flt, Flt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func type_conversion() {
	var i int = 42
	var f float32 = float32(i)
	u := uint(f)

	fmt.Println(i, f, u)
}

func for_loop() {
	sum := 1

	for i := 0; i < 10; i++ {
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

	for i := 0; i < 10; i++ {
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
	for i := 0; i < 10; i++ {
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

func slices() {
	// array literal
	primes := [6]int{2, 3, 5, 7, 11, 13}
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
		[]string{"foo", "bar"},
		[]string{"qaz", "zaq"},
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
