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

	// Dados padrão para o template
	data := map[string]interface{}{
		"hasInstance":  false,
		"isConnected":  false,
		"qrCodeBase64": "",
	}

	// 1️⃣ Se não tiver instância → exibe formulário de criação
	if credencial.TokenApi == "" || credencial.InstanceApi == "" {
		fmt.Println("Nenhuma instância encontrada. Solicitando criação...")
		templates.ExecuteTemplate(w, "whatsapp.html", data)
		return
	}

	data["hasInstance"] = true

	// 2️⃣ Verifica status da instância via API
	status, err := services.GetZAPIApiStatus(credencial.TokenApi, credencial.InstanceApi)
	if err != nil {
		data["error"] = fmt.Sprintf("Erro ao verificar o status da API: %v", err)
		templates.ExecuteTemplate(w, "whatsapp.html", data)
		return
	}

	fmt.Printf("Status da instância: %+v\n", status)

	// 3️⃣ Se estiver conectada → exibe mensagem de sucesso
	if status.Connected && status.SmartphoneConnected {
		data["isConnected"] = true
		templates.ExecuteTemplate(w, "whatsapp.html", data)
		return
	}

	// 4️⃣ Se estiver desconectada → tenta buscar QR Code
	qrCodeBase64, err := services.GetZAPIApiQrCode(credencial.TokenApi, credencial.InstanceApi)
	if err != nil {
		// Se a mensagem indicar resposta vazia → instância ainda iniciando
		if strings.Contains(err.Error(), "resposta vazia") {
			data["error"] = "A instância está inicializando. Aguarde alguns segundos e recarregue a página para gerar o QR Code."
		} else {
			data["error"] = fmt.Sprintf("Erro ao obter QR Code: %v", err)
		}
		templates.ExecuteTemplate(w, "whatsapp.html", data)
		return
	}

	// 5️⃣ QR obtido → renderiza página com o código
	data["qrCodeBase64"] = qrCodeBase64
	data["isConnected"] = false
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

	IAM.CriarCredencial(searchPath, "WhatsApp", "", instanceResponse.InstanceToken, instanceResponse.InstanceID, "")

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

	qrCodeBase64, err := services.GetZAPIApiQrCode(credencial.TokenApi, credencial.InstanceApi)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter o QR Code: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"qrCode": qrCodeBase64}
	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
