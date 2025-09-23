package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const receivedCallbackURL = "https://service-api.hareinteract.com.br/webhook-zapi-foa"

func CreateZAPIApiInstance(token, instanceName string) ([]byte, error) {
	url := "https://api-prd.joindeveloper.com.br/instances/integrator/on-demand"

	payload := fmt.Sprintf(`{"name": "%s", "receivedCallbackUrl": "%s"}`, instanceName, receivedCallbackURL)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição: %w", err)
	}

	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status de resposta inesperado: %s, corpo da resposta: %s", res.Status, string(body))
	}

	return body, nil
}

func GetZAPIApiQrCode(token, instanceID string) ([]byte, error) {
	url := fmt.Sprintf("https://api-prd.joindeveloper.com.br/instances/%s/token/%s/qr-code/image", instanceID, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status de resposta inesperado: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %w", err)
	}

	return body, nil
}
