package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	//Print Hello
	fmt.Println("Hello there!")
	//------------------------------------------------------------------------------------------------------

	//Sum1
	fmt.Println("\nSum1------------------------------------------------------------------")
	fmt.Println(sum(3, 4))
	//------------------------------------------------------------------------------------------------------

	//Sum2
	fmt.Println("\nSum2------------------------------------------------------------------")
	result := sum2(3, 4)
	fmt.Println(result)

	//------------------------------------------------------------------------------------------------------
	//Calling a list generator
	fmt.Println("\nList------------------------------------------------------------------")
	list()
	//------------------------------------------------------------------------------------------------------

	//Calling a map gen
	fmt.Println("\nMap------------------------------------------------------------------")
	fmt.Printf("%d", maping())
	//------------------------------------------------------------------------------------------------------

	//Errors
	fmt.Println("\nErrors------------------------------------------------------------------")
	result2, err := sqrt(-15)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
	//------------------------------------------------------------------------------------------------------

	//Struct
	fmt.Println("\nStruct------------------------------------------------------------------")
	p := person{name: "Eugeni", age: 24}
	fmt.Println(p)
	fmt.Println(p.age)

	//------------------------------------------------------------------------------------------------------
	//Pointer
	fmt.Println("\nPointer------------------------------------------------------------------")
	i := 7
	fmt.Println(&i, i)
	inc(&i) //if we want to change the original number (in the same addres),
	//we have to create a pointer in another variable that modifies it in a diferent function
	fmt.Println(&i, i)

}

//------------------------------------------------------------------------------------------------------
//------------------------------------------------------------------------------------------------------
//FUNCTIONS

//sum1
func sum(x, y int) (result int) {
	result = x + y
	return
}

//------------------------------------------------------------------------------------------------------
//sum2
func sum2(x, y int) int {
	return x + y
}

//------------------------------------------------------------------------------------------------------
//list
func list() {
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := []int{2, 3, 4, 5}
	c = append(c, 7)
	fmt.Println(c)

}

//------------------------------------------------------------------------------------------------------
//map
func maping() int {
	vertices := make(map[string]int) //diccionari amb key = string i values = int, the make built in function is used to create the map
	vertices["First"] = 1
	vertices["Second"] = 2
	vertices["Third"] = 3
	delete(vertices, "Second")
	fmt.Println(vertices)

	for index, value := range vertices {
		fmt.Println(index, vertices[index], value)
	}
	return vertices[""]
}

//------------------------------------------------------------------------------------------------------

//Errors
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil

}

//------------------------------------------------------------------------------------------------------
//struct
type person struct { //class
	name string
	age  int
}

//------------------------------------------------------------------------------------------------------
//pointer
func inc(x *int) { //we dont create any return variable
	*x++
}

// how pointers work: https://www.golang-book.com/books/intro/8

//------------------------------------------------------------------------------------------------------
