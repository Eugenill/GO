package main


import(
	"fmt"
	"errors"
	"math"
)

func main(){
	//Print Hello
	fmt.Println("Hello there!")
	//Sum1
	fmt.Println(sum(3, 4))
	//Sum2
	result := sum2(3, 4)
	fmt.Println(result)
	//Calling a list generator
	list()
	//Calling a map gen
	maping()
	//Errors
	result2, err := sqrt(-15)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
	//Struct
	p := person{name: "Eugeni", age:24}
	fmt.Println(p)
	fmt.Println(p.age)
	//Pointer
	i := 7 
	fmt.Println(&i, i)
	inc(&i) //if we want to change the original number (in the same addres), 
	//we have to create a pointer in another variable that modifies it in a diferent function
	fmt.Println(&i, i) 


}

func sum(x, y int) (result int){
	result = x + y
	return
}
func sum2(x, y int) int { 
	return x + y
}

func list() {
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	b := [5]int{1,2,3,4,5}
	fmt.Println(b)

	c := []int{2,3,4,5}
	c = append(c,7)
	fmt.Println(c)	

}

func maping(){
	vertices := make(map[string]int) //diccionari amb key = string i values = int, the make built in function is used to create the map
	vertices["First"]=1
	vertices["Second"]=2
	vertices["Third"]=3 
	delete(vertices,"Second")
	fmt.Println(vertices)

	for index, value := range vertices {
		fmt.Println(index, vertices[index], value)
	} 
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0,  errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil

}

type person struct { //class
	name string
	age int
}

func inc(x *int) { //we dont create any return variable
	*x++
}
