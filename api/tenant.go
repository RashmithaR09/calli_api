package api

import (
	"api/types"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchcreateTenantHandler(w http.ResponseWriter, r *http.Request) (*types.TenantUpdateResponse, error) {
	var req types.TenantCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid tenant data", http.StatusBadRequest)
		return nil, err

	}
	url := fmt.Sprintf("%s/master/tenants", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"))
	// fmt.Println("Post URL:", url)

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to marshal request body", http.StatusInternalServerError)
		return nil, err
	}

	ctx := context.Background()
	req3rdParty, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return nil, err

	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
		return nil, err

	}
	req3rdParty.Header.Set("Authorization", token)
	req3rdParty.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req3rdParty)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return nil, err

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return nil, err

	}

	if resp.StatusCode != http.StatusCreated {
		http.Error(w, fmt.Sprintf("Failed to create tenant: %s, response body: %s", resp.Status, string(body)), resp.StatusCode)
		return nil, err

	}
	var response types.TenantUpdateResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil

}

// GET
func FetchgetAllTenantsHandler(w http.ResponseWriter, r *http.Request) (*types.TenantUpdateRequest, error) {
	url := fmt.Sprintf("%s/master/tenants", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"))

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return nil, err
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
		return nil, err
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Failed to retrieve tenants: %s, response body: %s", resp.Status, string(body)), resp.StatusCode)
		return nil, err
	}

	var response types.TenantUpdateRequest
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		return nil, err
	}
	return &response, nil
}

//GET

func FetchgetTenantHandler(tenantID int, w http.ResponseWriter, r *http.Request) (*types.TenantUpdateRequest, error) {
	url := fmt.Sprintf("%s/master/tenants/:%d", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantID)

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return nil, err
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
		return nil, err
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Failed to retrieve tenants: %s, response body: %s", resp.Status, string(body)), resp.StatusCode)
		return nil, err
	}

	var response types.TenantUpdateRequest
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		return nil, err
	}
	return &response, nil
}

func FetchupdateTenantHandler(tenantID int, w http.ResponseWriter, r *http.Request) (*types.TenantUpdateResponse, error) {
	var req types.TenantUpdateRequest

	id := r.URL.Query().Get("tenantID")
	if id == "" {
		http.Error(w, "Invalid tenant ID", http.StatusBadRequest)
		return nil, fmt.Errorf("invalid tenant ID")
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid tenant data", http.StatusBadRequest)
		return nil, err
	}

	url := fmt.Sprintf("%s/master/tenants/:%d", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantID)

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to marshal request body", http.StatusInternalServerError)
		return nil, err
	}
	ctx := context.Background()
	req3rdParty, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return nil, err
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
		return nil, fmt.Errorf("authorization token is missing")
	}

	req3rdParty.Header.Set("Authorization", token)
	req3rdParty.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req3rdParty)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Failed to update tenant: %s, response body: %s", resp.Status, string(body)), resp.StatusCode)
		return nil, fmt.Errorf("failed to update tenant: %s", resp.Status)
	}

	var response types.TenantUpdateResponse
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		return nil, err
	}

	return &response, nil
}

//DELETE

func FetchdeleteTenantHandler(tenantID int, w http.ResponseWriter, r *http.Request) error {

	url := fmt.Sprintf("%s/master/tenants/:%d", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantID)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, url, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return err
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
		return err
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return err
		}
		http.Error(w, fmt.Sprintf("Failed to delete tenant: %s, response body: %s", resp.Status, string(body)), resp.StatusCode)
		return nil
	}

	return nil
}
