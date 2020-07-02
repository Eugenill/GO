package main

import(
	"math"
	"fmt"
	"io"
	"image"
	"strings"
)

func main() {
//------------------------------------------------------------------------------------------------------

	//Methods from types
	fmt.Println("\nMethod from types------------------------------------------------------------------")
	f := MyFloat(-math.Sqrt2) //to define a type value we write it into ()
	fmt.Println(f.Abs())
//------------------------------------------------------------------------------------------------------

	//Pointer receivers 
	fmt.Println("\nPointer receivers------------------------------------------------------------------")
	v := Vertex{3, 4}
	v.Scale(10) //Go interprets the statement v.Scale(10) as (&v).Scale(10) 
				//since the Scale method has a pointer receiver.
	fmt.Println(v.Abs(), v)
//------------------------------------------------------------------------------------------------------

	//Pointers and functions
	fmt.Println("\nPointer vs Funtions------------------------------------------------------------------")
	m := &Vertex{1, 2}
	Scale(m, 10)
	fmt.Println(m.Abs(), m)
//------------------------------------------------------------------------------------------------------
	//Interfaces
	fmt.Println("\nInterfaces------------------------------------------------------------------")
	var a Abser
	//f := MyFloat(-math.Sqrt2)
	//v := Vertex{3, 4}
	describe(a)
	a = f // a MyFloat type implements Abser, 
		//we can only assign those type's with a return type's
		//the same as the interface return type's
	fmt.Println(a.Abs())
	describe(a) // the type of a now is 'main.Myfloat' as we are asigning the interface Myfloat in package main
	a = &v // a *Vertex type implements Abser
	describe(a) // the type of a now is '*main.Vertex' as we are asigning the interface *Vertex in package main
	fmt.Println(a.Abs())
	
	fmt.Printf("%T\n",a.Abs()) //float64 type, as we specified in Abs() float64
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser. We add &
	//a = v
	a = &v //'a' must be a pointer to a Vertex because the Abs method is defined only 
			//on *Vertex (pointer type)
	fmt.Println(v)
	fmt.Println(a.Abs())
	fmt.Println(v.Abs()) //the original type Vertex = v dont need to be a pointer
						//to use (*Vertex) Abs(), because it takes (&v) automatically instead of v
//------------------------------------------------------------------------------------------------------

	//The empty interface
	fmt.Println("\nEmpty interfaces------------------------------------------------------------------")
	var g interface{}
	describe2(g)

	g = 42
	describe2(g)

	g = "hello"
	describe2(g)
	//TYPE ASSERTION
	s, ok := g.(string)
	fmt.Println(s, ok)

//------------------------------------------------------------------------------------------------------

	//stringers
	fmt.Println("\nStringers------------------------------------------------------------------")
	r := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(r, z) //the fmt package look for the Stringer interface to print values
	fmt.Printf("%T", r)

//------------------------------------------------------------------------------------------------------

	//READERS: The io package specifies the io.Reader interface, which represents the read end of a stream of data.
	/*

	type Reader interface {
    	Read(p []byte) (n int, err error)
	}

	Read method: func (T) Read(b []byte) (n int, err error)

	Read populates the given byte slice with data and returns the number of bytes populated and an error value.
	It returns an io.EOF error when the stream ends.
	*/
	fmt.Println("\nReader------------------------------------------------------------------")
	t := strings.NewReader("Hello, Reader!") //t is a *string.Reader (pointer) with value= "Hello, Reader!"
	// func NewReader(s string) *Reader, returns a *reader
	fmt.Printf("%T\n",t)

	b := make([]byte, 15) //depending on the value of bytes we will need 1 or more loopings
	for {
		n, err := t.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		} //stops when it has read all the string t introduced
	}

//------------------------------------------------------------------------------------------------------

	//images

	/*
	type Image interface {
	    ColorModel() color.Model
	    Bounds() Rectangle
	    At(x, y int) color.Color
	}
	*/
	fmt.Println("\nImages------------------------------------------------------------------")

	im := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Printf("%T\n",im) // *image.RGBA type
	fmt.Println(im.Bounds()) //bounds is a method of image, and returns a image.Rectangle
	fmt.Printf("%T\n",im.Bounds()) // image.Rectagle, as the declaration is inside package image
	fmt.Println(im.At(0, 0).RGBA())

}

//FUNCTIONS

//------------------------------------------------------------------------------------------------------
//Methods
type MyFloat float64

func (f MyFloat) Abs() float64 { //methods: You can only declare a method with
								// a receiver whose type is defined in the same package as
								// the method. You cannot declare a method with a receiver
								// whose type is defined in another package (which includes 
								//the built-in types such as int).
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//------------------------------------------------------------------------------------------------------
//Pointer receivers
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 { //in case we have a return in the function
								//we just need to input a pointer in case
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//methods can take a value or a pointer
func (v *Vertex) Scale(f float64) { //we must add a pointer because we are modifying the 
									//the variable of type vertex itself
	v.X = v.X * f
	v.Y = v.Y * f
}
//------------------------------------------------------------------------------------------------------

//Pointers and functions
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


//------------------------------------------------------------------------------------------------------
//Interfaces
type Abser interface { //A value of interface type can hold any value that implements those methods.
	Abs() float64 	   //An interface type is defined as a set of method signatures.
}

func describe(a Abser) {
	fmt.Printf("(%v, %T)\n", a, a) //values, type
}


//The empty interface
func describe2(g interface{}) {
	fmt.Printf("(%v, %T)\n", g, g)
}


//------------------------------------------------------------------------------------------------------

//STRINGERS, defined by the fmt package

//Is used to define another way to print a type
/*

type Stringer interface {
    String() string
}

*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { //if we add this method to a type (Person) now the var is also a Stringer type (interface)
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}