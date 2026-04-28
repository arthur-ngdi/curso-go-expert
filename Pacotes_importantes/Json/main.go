package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 123, Saldo: 100}
	/* res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	} */

	//println(res)
	//println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var conta2 Conta
	err := json.Unmarshal(jsonPuro, &conta2)
	if err != nil {
		panic(err)
	}
}
