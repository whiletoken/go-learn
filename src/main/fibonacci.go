package main

import "fmt"

func fibonacci(num int) int {

	if num == 0 {
		return 0
	}
	if num <= 2 {
		return 1
	}

	pre := 1
	curr := 1
	for i := 3; i <= num; i++ {
		temp := curr + pre
		pre = curr
		curr = temp
	}
	return curr
}

func fibonacci1(num int) int {

	if num == 0 {
		return 0
	}
	if num <= 2 {
		return 1
	}

	return fibonacci1(num-1) + fibonacci1(num-2)
}

func fibonacci2() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fibonacci3(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		c <- a
	}

	close(c)
}

func test() {

	fibonacci(5)

	fibonacci(5)

	f := fibonacci2()
	for i := 0; i < 5; i++ {
		fmt.Print(f())
	}

	c := make(chan int, 10)

	fibonacci3(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
