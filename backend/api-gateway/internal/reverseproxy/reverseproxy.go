package reverseproxy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

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
			"investments": {"http://investments:80"},
			"property":    {},
			"reports":     {},
			"simulator":   {},
		},
	}
}
func (rp *ReverseProxy) ProxyHandler(c *gin.Context) {
	targetURL, err := rp.resolveTarget(c.Request.URL.Path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid microservice host"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	originalHost := c.Request.Host
	originalPath := c.Request.URL.Path

	c.Request.Host = targetURL.Host
	c.Request.URL.Path = prepareCompleteUrlPath(originalPath)

	if userID, exists := c.Get("user_id"); exists {
		c.Request.Header.Set("user_id", fmt.Sprint(userID))
	}

	proxy.ServeHTTP(c.Writer, c.Request)

	c.Request.Host = originalHost
	c.Request.URL.Path = originalPath
}

func (rp *ReverseProxy) resolveTarget(path string) (*url.URL, error) {
	host, err := rp.selectMicroService(path)
	if err != nil {
		return nil, err
	}

	parsed, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}
func (rp *ReverseProxy) forwardRequest(target *url.URL, userID string, w http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Salva estado original do request
	originalHost := r.Host
	originalPath := r.URL.Path

	// Reescreve request
	r.Host = target.Host
	r.URL.Path = prepareCompleteUrlPath(originalPath)
	if userID != "" {
		r.Header.Set("user_id", userID)
	}

	proxy.ServeHTTP(w, r)

	// Restaura estado original
	r.Host = originalHost
	r.URL.Path = originalPath
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

type UserAuthenticated struct {
	UserId int `json:"user_id"`
}

func (rp *ReverseProxy) VerifyAuthentication(ctx context.Context, headers http.Header) (int, error) {

	microServiceHost, err := rp.selectMicroService("/api/auth/check")
	if err != nil {
		return 0, err
	}

	remoteUrl := fmt.Sprintf("%s/isAuthenticated", microServiceHost)

	request, err := http.NewRequestWithContext(ctx, "GET", remoteUrl, nil)
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

	var user UserAuthenticated
	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		return 0, fmt.Errorf("failed to decode user: %w", err)
	}

	if user.UserId == 0 {
		return 0, fmt.Errorf("authentication: received invalid user ID")
	}

	return int(user.UserId), nil
}
