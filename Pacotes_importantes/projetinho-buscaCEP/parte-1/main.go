package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/"+ url +"/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			continue
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
			continue
		}
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar JSON: %v\n", err)
			continue
		}
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
			continue
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nLogradouro: %s\nComplemento: %s\nUnidade: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nEstado: %s\nRegião: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n",
			data.Cep, data.Logradouro, data.Complemento, data.Unidade, data.Bairro, data.Localidade, data.Uf, data.Estado, data.Regiao, data.Ibge, data.Gia, data.Ddd, data.Siafi))

	}
}
