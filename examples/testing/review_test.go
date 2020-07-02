package models

/*------TESTING-----------------
Whenever we have a function with a param = (x *testing.T), go test automates the execution of this function

*/
//to test all the possible tests in a folder, run in terminal: go test -v ./... (-verbose, more data)
//test files have a unique structure, is the next one: func Test_......(t *testing.X)

import (
	"testing" //https://golang.org/pkg/testing/
)

//creamos la review test para testear la función de validate
func NewReview(stars int, comment string) CreateReviewCMD {  //va a retornar un puntero del tipo struct CMD
	return CreateReviewCMD{
		Stars: stars,
		Comment: comment,
	} //create a new 
}

//Test function to validate a review depending on the boundings of the parameters we created
func Test_passReview(t *testing.T) { //T is a type passed to Test functions to manage test state and support formatted test logs.
	r := NewReview(4, "The iPhone X looks good")

	err := r.validate() //this method returns an error and its value depends on the struct

	if err != nil {
		t.Error("The validation did not pass") //what we tell to the cmd
		t.Fail()//Error is always followed by Fail()
	}
}

//to test the reviews we just have to run: go test review_test.go
func Test_notPassReview(t *testing.T) {
	r := NewReview(8, "The iPhone X looks good")

	err := r.validate()

	if err != nil {
		t.Error("The validation did not pass")
		t.Fail()
	}
}
//---------------------------------------------------------------------------------------- Step 1

