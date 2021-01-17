package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
	// 	"github.com/ajayvm/greetings"
)

type pt struct {
	x int
	y int
}

func mainLater() {
	s := "Ajay Mahajan 2 and see the world "
	defer fmt.Println("hello world")
	fmt.Println(s, addn(4, 5))
	fmt.Println(time.Now())
	fmt.Println(sqrt(64), sqrt(-9))
	fmt.Println(runtime.GOARCH)

	ptr := &s
	fmt.Println(ptr, *ptr)

	v := pt{1, 2}
	fmt.Println(v, v.x)

	ptArr := [][]pt{
		{{1, 2}, {3, 4}}, {{61, 26}, {35, 47}},
	}
	fmt.Println(ptArr)
}

func sqrt(x float64) float64 {
	if x < 0 {
		return sqrt(-x)
	}
	return math.Sqrt(x)
}

func addn(a, b int) int {
	return a + b
}
