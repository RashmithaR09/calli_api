package api

import (
	"api/types"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//user.go
//POST

func FetchcreateUserHandler(ctx context.Context, tenantName string, userRequest types.UserRequest, r *http.Request) (*types.UserResponse, error) {
	url := fmt.Sprintf("%s/tenants/%s/users", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantName)

	reqBody, err := json.Marshal(userRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Retrieve the authorization token from the request header
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("authorization token is missing")
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create user: %s, response body: %s", resp.Status, string(body))
	}

	var userResponse types.UserResponse
	if err := json.Unmarshal(body, &userResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return &userResponse, nil
}

// GET
func FetchgetAllUsersHandler(ctx context.Context, tenantName string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/tenants/:%s/users", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantName)
	fmt.Println(url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer your_token_here")
	req.Header.Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s, response body: %s", resp.Status, string(body))
	}

	var UserResponse map[string]interface{}
	if err := json.Unmarshal(body, &UserResponse); err != nil {
		return nil, err
	}

	return UserResponse, nil

}

// GET
func FetchgetUserHandler(ctx context.Context, tenantName string, userID uint) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/tenants/:%s/users/:%d", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantName, userID)
	fmt.Println(url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer your_token_here")
	req.Header.Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s, response body: %s", resp.Status, string(body))
	}

	var UserResponse map[string]interface{}
	if err := json.Unmarshal(body, &UserResponse); err != nil {
		return nil, err
	}

	return UserResponse, nil

}

// PUT
func FetchupdateUserHandler(ctx context.Context, tenantName string, userId int, user types.ClientRequest) (*types.UserResponse, error) {
	url := fmt.Sprintf("%s/tenants/:%s/users/:%d", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantName, userId)

	userRequestBody, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(userRequestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("AUTH_TOKEN")) // Set your authorization token here
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update user, status code: %d", resp.StatusCode)
	}

	var userResponse types.UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return nil, err
	}
	return &userResponse, nil
}

// DELETE
func FetchdeleteUserHandler(ctx context.Context, tenantName string, clientID uint) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/tenants/:%s/users/:%d", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantName, clientID)
	fmt.Println("DELETE Request URL:", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer your_token_here")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s, response body: %s", resp.Status, string(body))
	}

	var UserResponse map[string]interface{}
	if err := json.Unmarshal(body, &UserResponse); err != nil {
		return nil, err
	}

	return UserResponse, nil
}

//POST

func FetchaddRolesToUserHandler(ctx context.Context, userId int, tenantName string, userRequest types.UserRequest, r *http.Request) (*types.UserResponse, error) {
	url := fmt.Sprintf("%s/tenants/:%s/users/:%d/roles", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantName, userId)

	reqBody, err := json.Marshal(userRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Retrieve the authorization token from the request header
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("authorization token is missing")
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to add roles to user: %s, response body: %s", resp.Status, string(body))
	}

	var userResponse types.UserResponse
	if err := json.Unmarshal(body, &userResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return &userResponse, nil
}

// GET
func FetchgetUserRolesHandler(ctx context.Context, tenantName string, userID uint) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/tenants/:%s/users/:%d/roles", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantName, userID)
	fmt.Println(url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer your_token_here")
	req.Header.Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s, response body: %s", resp.Status, string(body))
	}

	var UserResponse map[string]interface{}
	if err := json.Unmarshal(body, &UserResponse); err != nil {
		return nil, err
	}

	return UserResponse, nil

}

// DELETE
func FetchremoveRolesFromUserHandler(ctx context.Context, tenantName string, userID uint) error {
	url := fmt.Sprintf("%s/tenants/:%s/users/:%d/roles", os.Getenv("AIOTRIX-GAURD-IDP-API-URL"), tenantName, userID)
	fmt.Println("DELETE Request URL:", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer your_token_here")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data: %s, response body: %s", resp.Status, string(body))
	}

	var UserResponse map[string]interface{}
	if err := json.Unmarshal(body, &UserResponse); err != nil {
		return err
	}

	return err
}
