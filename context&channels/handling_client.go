package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	//we make a request to the server
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	//res, err := http.Get("http://localhost:8080") without passing the context&channels created above
	if err != nil {
		log.Fatal(err)
	}
	req.WithContext(ctx) //we add the context&channels created in the beggining
	res, er := http.DefaultClient.Do(req)
	if er != nil {
		log.Fatal(er)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	//printegem la resposta al command. Pero fent un copy del body al Stdout
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err)
	}
	//log.Print(res.Body)->2020/07/20 17:13:44 &{0xc0000cc380 {0 0} false 0xc000098030 <nil> 0x636650}
}
