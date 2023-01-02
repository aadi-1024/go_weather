package main

import (
	"net/http"
	"fmt"
)

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/astronomy.json?key=681cb67743fd4c2181f133308230201&q=Ludhiana")
	if err == nil {
		body := res.Body
		b := make([]byte, 500)
		body.Read(b)
		fmt.Println(string(b))
	}
}