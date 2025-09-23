package controllers

import (
	"HareInteract.WebApp/models/IAM"
	"HareInteract.WebApp/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type InstanceCreationResponse struct {
	InstanceID    string `json:"id"`
	InstanceToken string `json:"token"`
}

func WhatsAppHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informação de sessão não encontrada.", http.StatusUnauthorized)
		return
	}

	credencial := IAM.ObterCredencialPorTitulo(searchPath, "WhatsApp")
	fmt.Printf("Credencial obtida: %+v\n", credencial)

	data := map[string]interface{}{
		"hasInstance":  false,
		"isConnected":  false,
		"qrCodeBase64": "",
		"searchPath":   searchPath, // Enviar para o template para a lógica de exibição.
	}

	if credencial.TokenApi != "" && credencial.InstanceApi != "" {
		data["hasInstance"] = true

		_, err := services.GetZAPIApiQrCode(credencial.TokenApi, credencial.InstanceApi)
		if err == nil {
			data["isConnected"] = false
		} else if strings.Contains(err.Error(), "404") {
			data["isConnected"] = true
		} else {
			data["error"] = fmt.Sprintf("Erro ao verificar o status: %v", err)
			data["isConnected"] = false
		}
	}
	fmt.Printf("Data: %+v\n", data)

	templates.ExecuteTemplate(w, "whatsapp.html", data)
}

func CriarInstanciaHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informação de sessão não encontrada.", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/whatsapp", http.StatusSeeOther)
		return
	}

	instanceName := r.FormValue("instanceName")
	if instanceName == "" {
		http.Redirect(w, r, "/whatsapp?error=nome-obrigatorio", http.StatusSeeOther)
		return
	}

	apiToken := "ea49bd3e-652d-4cfd-be9d-4aea210d73b0"

	instanceDataBytes, err := services.CreateZAPIApiInstance(apiToken, instanceName)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/whatsapp?error=falha-criacao-instancia&details=%v", err), http.StatusSeeOther)
		return
	}

	var instanceResponse InstanceCreationResponse
	err = json.Unmarshal(instanceDataBytes, &instanceResponse)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/whatsapp?error=falha-decodificacao-resposta&details=%v", err), http.StatusSeeOther)
		return
	}

	instanceID := instanceResponse.InstanceID
	instanceToken := instanceResponse.InstanceToken

	// TODO: Configurar o webhook
	// services.ConfigureWebhook(instanceID, instanceToken)

	IAM.CriarCredencial(searchPath, "WhatsApp", "", instanceToken, instanceID, "")

	http.Redirect(w, r, "/whatsapp", http.StatusSeeOther)
}

func QrCodeHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informação de sessão não encontrada.", http.StatusUnauthorized)
		return
	}

	credencial := IAM.ObterCredencialPorTitulo(searchPath, "WhatsApp")
	if credencial.TokenApi == "" || credencial.InstanceApi == "" {
		http.Error(w, "Nenhuma instância encontrada.", http.StatusNotFound)
		return
	}

	// AQUI: A API retorna uma string Base64, então 'qrCodeImage' será uma string.
	qrCodeBase64, err := services.GetZAPIApiQrCode(credencial.TokenApi, credencial.InstanceApi)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter o QR Code: %v", err), http.StatusInternalServerError)
		return
	}

	// Prepara a resposta JSON
	response := map[string]string{"qrcode": string(qrCodeBase64)}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Erro ao serializar a resposta JSON", http.StatusInternalServerError)
		return
	}

	// Envia a resposta como JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
