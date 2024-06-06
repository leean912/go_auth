package main

import (
	"fmt"

	"rsc.io/quote"
)

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
	D     = true
)

func main() {
	fmt.Println("hello world")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())

	// var number int = 2
	// var floatNum = 0.2
	var arr = [...]string{"a", "b", "c"}
	// var arr2 = [...]string{2: "a", 4: "b", 6: "c"}
	// mySliceArr := arr[0:2]
	// mySlice := arr2[2:3]
	// myMakeSlice := make([]string, 5, 10)
	// x := 3
	// myAppendSlice := append(arr2[:], "e", "f")
	// appendTwoSlice := append(arr2[:], mySlice...)

	// fmt.Println(number)
	// fmt.Println(floatNum)
	// fmt.Println("hello", x)
	// fmt.Printf("value: %v\n", x)
	// fmt.Printf("type: %T\n", x)
	// fmt.Printf("%#v\n", C)
	// fmt.Printf("%t\n", D)
	// fmt.Println(arr)
	// fmt.Printf("arr2 %v\n", arr2)
	// fmt.Printf("My Slice: %v\n", mySlice)
	// fmt.Printf("My Make Slice cap %d, len %d\n", cap(myMakeSlice), len(myMakeSlice))
	// fmt.Printf("My Append Slice %v\n", myAppendSlice)
	// fmt.Printf("My two appended slice %v\n", appendTwoSlice)

	// fmt.Printf("arr: %v, %v, %v, slice %v\n", arr[0], arr[1], arr[2], mySliceArr)

	var test = make([]string, 4, 5)
	test2 := make([]string, 4, 5)
	fmt.Print("test: ")
	fmt.Println(test)
	fmt.Print("test2: ")
	fmt.Println(test2)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for idx, val := range arr {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range arr {
		fmt.Println(val)
	}

	num := testFunction()
	fmt.Println(num)

	_, word := testFunction2()
	fmt.Println(word)

	var mobile Mobile
	mobile.isAndroid = true
	mobile.platform = "Android"
	mobile.version = "3.3.3"

	//map
	var a = map[string]int{"a": 0, "b": 1}
	fmt.Println(a)
	var b = make(map[string]string)
	var c map[string]string
	// c = b
	fmt.Println(b == nil)
	fmt.Println(c == nil)

	var d = map[string][]int{"a": {4, 5, 6, 7, 8}}
	fmt.Println(d)

	var e = make(map[string][]int)
	e["eee"] = []int{1, 2, 3, 4}
	e["a"] = []int{}

	delete(e, "a")

	fmt.Println(e)

}

func testFunction() (num int) {
	fmt.Println("I just got executed!")
	num = 0
	return num
}

func testFunction2() (num int, world string) {
	fmt.Println("I just got executed!")
	num = 0
	world = "hello world"
	return
}

type Mobile struct {
	platform  string
	version   string
	isAndroid bool
}

func copyMap(original map[string]int) map[string]int {
	// Create a new map
	copied := make(map[string]int)
	// Copy each key-value pair
	for key, value := range original {
		copied[key] = value
	}
	return copied
}
