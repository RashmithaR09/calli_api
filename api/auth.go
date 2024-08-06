package api

import (
	"api/types"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func FetchRegisterOAuthRoutes(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/auth/tenant/login", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"))

	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, "Failed to connect to third-party server", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Third-party server returned status: %d", resp.StatusCode), http.StatusInternalServerError)
		return
	}

	data := struct {
		RedirectUri string
	}{
		RedirectUri: r.URL.Query().Get("redirect-uri"),
	}

	tmpl, err := template.ParseFiles("tenant-login.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

// POST
func FetchtenantLoginHandler(ctx context.Context, loginRequest types.LoginRequest, w http.ResponseWriter, r *http.Request) (*types.LoginRequest, error) {
	url := fmt.Sprintf("%s/auth/tenant/login", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"))

	reqBody, err := json.Marshal(loginRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("authorization token is missing")
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create client: %s, response body: %s", resp.Status, string(body))
	}

	var response types.LoginRequest
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
