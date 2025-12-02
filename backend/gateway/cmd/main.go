package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type UserAuthenticated struct {
	UserId int `json:"user_id"`
}
type ReverseProxy struct {
	routes map[string][]string
}

func NewProxyRoutes() *ReverseProxy {
	return &ReverseProxy{
		routes: map[string][]string{
			"auth":        {"http://auth:80"},
			"transaction": {"http://transactions:80"},
			"budgets":     {},
			"forecast":    {},
			"goals":       {},
			"investments": {},
			"property":    {},
			"reports":     {},
			"simulator":   {},
		},
	}
}

func (rp *ReverseProxy) selectMicroService(path string) (string, error) {

	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")

	if len(parts) < 2 {
		return "", errors.New("Invalid path structure")
	}
	serviceName := parts[1]
	targetUrl, exists := rp.routes[serviceName]
	if !exists || len(targetUrl) == 0 {
		return "", errors.New("Service not found")
	}

	return targetUrl[0], nil

}

func prepareCompleteUrlPath(path string) string {

	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	newpath := strings.TrimPrefix(path, "/"+parts[0]+"/"+parts[1])

	return newpath

}
func (rp *ReverseProxy) createProxyRequest(w http.ResponseWriter, r *http.Request) (*http.Request, error) {
	microService, err := rp.selectMicroService(r.URL.Path)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return nil, nil
	}
	newPathRequest := prepareCompleteUrlPath(r.URL.Path)
	remoteUrl, err := url.Parse(microService)
	if err != nil {
		http.Error(w, "Bad gateway", http.StatusBadGateway)
		return nil, nil
	}

	return http.NewRequest(r.Method, remoteUrl.String()+newPathRequest, r.Body)
}

func (rp *ReverseProxy) VerifyAuthentication(ctx context.Context, headers http.Header) (int, error) {
	microService, err := rp.selectMicroService("/api/auth")
	if err != nil {
		return 0, err
	}

	baseURL, err := url.Parse(microService)
	if err != nil {
		return 0, err
	}

	remoteUrl := baseURL
	remoteUrl.Path = "/isAuthenticated"
	println(remoteUrl)
	request, err := http.NewRequestWithContext(ctx, "GET", remoteUrl.String(), nil)
	if err != nil {
		return 0, err
	}

	request.Header = headers.Clone()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return 0, fmt.Errorf("auth service returned status: %d", resp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	log.Println("JSON Recebido:", string(bodyBytes))
	var user UserAuthenticated
	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		return 0, fmt.Errorf("failed to decode user: %w", err)
	}

	if user.UserId == 0 {
		return 0, fmt.Errorf("authentication:received invalid user ID")
	}

	return int(user.UserId), nil
}
func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var userId string
	pathRoute := r.URL.Path
	isAuthRoute := strings.HasPrefix(pathRoute, "/api/auth")
	isProfile := strings.HasPrefix(pathRoute, "/api/auth/profile")

	if !isAuthRoute || isProfile {
		user_id, err := rp.VerifyAuthentication(context.Background(), r.Header)
		if err != nil {
			log.Println(err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		userId = strconv.Itoa(user_id)
	}

	proxyRequest, err := rp.createProxyRequest(w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad gateway", http.StatusBadGateway)
		return
	}
	proxyRequest.Header = r.Header.Clone()
	proxyRequest.Header.Add("user_id", userId)

	resp, err := http.DefaultClient.Do(proxyRequest)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad gateway", http.StatusBadGateway)
		return
	}

	for k, v := range resp.Header {
		w.Header()[k] = v
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(bodyBytes)

	log.Printf("\n request: %s response:%s", r.URL.Path, string(bodyBytes))
}

func main() {
	rp := NewProxyRoutes()
	http.Handle("/api/", rp)

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Api stopped with error: %v\n", err)
	}
}
