package external_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

var (
	RequestCepErr = errors.New("ocorreu um erro na requisição ao viaCep")
	DecoderErr    = errors.New("ocorreu um erro na decodificação da resposta do viaCep")
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func GetAddressByCep(cep string) (*ViaCepResponse, error) {
	host := fmt.Sprintf(viper.GetString("VIA_CEP_HOST"), cep)

	resp, err := http.Get(host)
	if err != nil {
		return nil, RequestCepErr
	}

	var response *ViaCepResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, DecoderErr
	}
	defer resp.Body.Close()

	return response, nil
}
