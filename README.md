# GO_API
Creation of an API with Golang

## REST_People

### 1. Install needed library:
	go get github.com/gorilla/mux
### Usage
	Use the REST API by executing the REST_People file. 

	GET: http://localhost:3000/people --> To get the people in it
	GET: http://localhost:3000/people/{id} --> To get the person with id: {id}
	POST: http://localhost:3000/people/{ID} -->To add a new person (Header: Content-Type: app/json, Info in Body Raw)
	DELETE: http://localhost:3000/people/{ID} --> To delete a person in people with id:{id}
