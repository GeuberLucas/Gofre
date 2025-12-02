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
	userId uint `json:"user_id"`
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
	microService, err := rp.selectMicroService("auth")
	if err != nil {
		return 0, err
	}

	baseURL, err := url.Parse(microService)
	if err != nil {
		return 0, err
	}

	remoteUrl := baseURL

	request, err := http.NewRequestWithContext(ctx, "GET", remoteUrl.String()+"/is-authenticated", nil)
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

	var user UserAuthenticated
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return 0, fmt.Errorf("failed to decode user: %w", err)
	}

	if user.userId == 0 {
		return 0, fmt.Errorf("received invalid user ID")
	}

	return int(user.userId), nil
}
func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	proxyRequest, err := rp.createProxyRequest(w, r)
	proxyRequest.Header = r.Header

	if !strings.Contains(proxyRequest.URL.Path, "auth") || strings.Contains(proxyRequest.URL.Path, "/auth/profile") {
		user_id, err := rp.VerifyAuthentication(context.Background(), proxyRequest.Header)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		proxyRequest.Header.Add("user_id", strconv.Itoa(user_id))
	}

	resp, err := http.DefaultClient.Do(proxyRequest)
	if err != nil {
		http.Error(w, "Bad gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
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

	http.ListenAndServe(":8080", nil)

}
