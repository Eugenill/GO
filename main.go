package main

import ( 

	"encoding/json" //to encode to json
	"log" //to see errors on the server
	"net/http" //to write the server
	"github.com/gorilla/mux" //the one installed before
)

type Person struct {

	ID string `json:"id,omitempty"` //id minusculas , ID mayusculas, segun como lo escribamos sera como lo recibamos
	FirstName string `json:"FirstName,omitempty"`
	LastName string `json:"Lastname,omitempty"`
	Address *Address `json:"Address,omitempty"` //we use the structure create below by using *address
}

type Address struct{
	City string `json:"City,omitempty"`
	State string `json:"State,omitempty"`
}

//People database
var people []Person  //people contains Person's

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){ //w is response, req is the requesting info
	json.NewEncoder(w).Encode(people) //Encode from struct to json
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){ //w is response, req is the requesting info
	params := mux.Vars(req) //we set the request info in params
	for _, item := range people{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item) //we assign the item encoded to w, which has to be a new encoder
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{}) //if we dont find anything we respond with an empty Person json
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){ //w is response, req is the requesting info
	//we need to send a POST but we cannot do it with the browser, we can do it with POSTMAN
	params := mux.Vars(req)
	for _, item := range people{
		if params["id"] == item.ID {
			DeletePersonEndpoint(w,req) //if there is a user with the same ID we substitute it: delete it and insert it again
		}
	}
	var person Person //creamos una varible Person
	_ = json.NewDecoder(req.Body).Decode(&person) //we will find the content in Body, we also have to add a header of COntent-Type: application/json
	person.ID = params["id"]
	people = append(people, person)
	//example


	//json.NewEncoder(w).Encode(person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request){ //DELETE BY ID
	params := mux.Vars(req)
	for index, item := range people{
		if item.ID == params["id"]{
			people = append(people[:index], people[index + 1:]...) 
			json.NewEncoder(w).Encode(people)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}


func main() {
	//enroutador
	router := mux.NewRouter()

	//test people
	people = append(people, Person{ID:"1", FirstName:"Ryan", LastName:"Wazowsky", Address: &Address{City:"San Francisco", State:"California"}})
	people = append(people, Person{ID:"2", FirstName:"Joe", LastName:"Zowsky", Address: &Address{City:"San Francisco", State:"California"}})
	
	//endpoints
	router.HandleFunc("/people",GetPeopleEndpoint).Methods("GET") //when call /people we execute the function GetPeopleEndpoint
	
	//we can call the same endpoint with different methods
	//USAGE: localhost:300/people/1 or 2,3,4,5....
	router.HandleFunc("/people/{id}",GetPersonEndpoint).Methods("GET") //call person
	router.HandleFunc("/people/{id}",CreatePersonEndpoint).Methods("POST") //create person by posting
	router.HandleFunc("/people/{id}",DeletePersonEndpoint).Methods("DELETE") //delete person

	// to create the server locally in a port: http.ListenAndServe(":3000", router)
	//if we want to see if there is an error we introduce it in a log.Fatal

	log.Fatal(http.ListenAndServe(":3000", router))

}