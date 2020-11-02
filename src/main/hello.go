package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum - 1
	y = sum - x
	return
}

var (
	ToBe          = false
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func typeString() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func changeType() {
	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}

func loop() {
	sum := 1
	for sum < 10 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func switchString() {

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func point() {

	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

type Vertex struct {
	X int
	Y int
}

func subString() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s = primes[1:4]
	fmt.Println(s)
}

func arrayToString() {

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[1:8]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

type Vertex1 struct {
	Lat, Long float64
}

var m = map[string]Vertex1{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func sayHello() {
	const World = "世界"
	fmt.Println(World, math.Pi)
}
