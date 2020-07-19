package test

import (
	"time"
	"errors"
)

const maxLengthInComments = 400

type Review struct {
	ID int64 //similar to a long in Java
	Stars int //between 1-5
	Comment string //max 400 chars
	Date time.Time  //created at
}  

//CreateReviewCMD command to create a new review
type CreateReviewCMD struct {  
	Stars int `json:"stars"`
	Comment string `json:"comment"`
}  

//this function is needed to raise errors in case there is an error in the parameters of the CMD, 
//which will be the ones that we receive from outside

//METHOD of CreateReviewCMD Struct
func (cmd *CreateReviewCMD) validate() error { //return error, cmd is a Create... struct type
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("Stars must be less between 1 and 5")
	}

	if len(cmd.Comment) > maxLengthInComments {
		return errors.New("Comments must have less than 400 chars")
	}

	return nil

}
//---------------------------------------------------------------------------------------- Step 1