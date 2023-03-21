package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ViaCEP struct {
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

// go run main.go and curl localhost:8080
func main() {
	http.HandleFunc("/", BuscaCEP)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, err := BuscaCepHttp(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// -------------------------------------------------
	//Transforme o STRUCT de CEP em JSON e salve em W
	json.NewEncoder(w).Encode(cep)

	/*
		PODERIA SER FEITO DESSA FORMA

		result, err := json.Marshal(cep)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(result)
	*/
}

func BuscaCepHttp(cep string) (*ViaCEP, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var c ViaCEP
	//Transforme o JSON de BODY em STRUCT e salve na variavel C
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil

}
