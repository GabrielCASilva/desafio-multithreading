package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Resultado struct {
	Api       string
	Resultado []byte
	Erro      error
}

func main() {
	cep := "01001000"
	canal := make(chan Resultado, 2)

	go func() {
		canal <- chamadaApi(cep, func(ctx context.Context, cep string) (*http.Request, error, string) {
			req, err := http.NewRequestWithContext(ctx, "GET", "http://viacep.com.br/ws/"+cep+"/json/", nil)
			return req, err, "viaCep"
		})
	}()

	go func() {
		canal <- chamadaApi(cep, func(ctx context.Context, cep string) (*http.Request, error, string) {
			req, err := http.NewRequestWithContext(ctx, "GET", "https://brasilapi.com.br/api/cep/v1/"+cep, nil)
			return req, err, "brasilApi"
		})
	}()

	for i := 0; i < 2; i++ {
		result := <-canal
		if result.Erro == nil {
			fmt.Printf("Requisição feita com: %s \nResultado: %s", result.Api, string(result.Resultado))
			return
		}
	}

	fmt.Println("Ambas as chamadas de API falharam.")
}

func chamadaApi(cep string, cepRequest func(ctx context.Context, cep string) (*http.Request, error, string)) (resultado Resultado) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err, api := cepRequest(ctx, cep)
	if err != nil {
		return Resultado{Api: api, Resultado: nil, Erro: err}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Resultado{Api: api, Resultado: nil, Erro: err}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Resultado{Api: api, Resultado: nil, Erro: err}
	}

	return Resultado{Api: api, Resultado: body, Erro: nil}
}
