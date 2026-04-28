package main

import (
	"bytes"
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Arthur Antunes Coimbra"}`))
	res, err := c.Post("https://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
